import { TerraformOutput, TerraformStack } from 'cdktf';
import { Construct } from 'constructs';

import { AcmCertificate } from '../.gen/providers/aws/acm-certificate';
import { AcmCertificateValidation } from '../.gen/providers/aws/acm-certificate-validation';
import { Route53Record } from '../.gen/providers/aws/route53-record';
import { Route53Zone } from '../.gen/providers/aws/route53-zone';
import { EnvironmentOptions } from './shared/models/environment.interface';

export interface DnsStackProps extends EnvironmentOptions {
  domainName: string;
  services: string[]; // e.g., ["frontend", "backend", "optimizer", ...]
}

export class DnsStack extends TerraformStack {
  public readonly certificateArn: string;
  public readonly hostedZoneId: string;
  public readonly serviceFqdns: Record<string, string> = {};

  constructor(scope: Construct, id: string, props: DnsStackProps) {
    super(scope, id);

    // 1️⃣ Hosted Zone
    const zone = new Route53Zone(this, `${props.env}-zone`, {
      name: props.domainName,
      tags: { Environment: props.env },
    });
    this.hostedZoneId = zone.zoneId;

    // 2️⃣ ACM Wildcard Certificate
    const cert = new AcmCertificate(this, `${props.env}-wildcard-cert`, {
      domainName: props.domainName,
      subjectAlternativeNames: props.services.map((s) => `${s}.${props.domainName}`),
      tags: { Environment: props.env },
      validationMethod: 'DNS',
    });
    this.certificateArn = cert.arn;

    // 3️⃣ DNS Validation Records for ACM
    props.services.forEach((service) => {
      const record = new Route53Record(
        this,
        `${props.env}-${service}-cert-validation`,
        {
          name: `_acme-challenge.${service}.${props.domainName}`,
          records: ['TO_BE_FILLED_BY_CERT_VALIDATION'], // Replace dynamically in real scenario
          ttl: 300,
          type: 'CNAME',
          zoneId: zone.zoneId,
        },
      );

      new AcmCertificateValidation(
        this,
        `${props.env}-${service}-cert-validation-step`,
        {
          certificateArn: cert.arn,
          validationRecordFqdns: [record.fqdn],
        },
      );
    });

    // 4️⃣ Subdomain records for services
    props.services.forEach((service) => {
      const fqdn = `${service}.${props.domainName}`;
      new Route53Record(this, `${props.env}-${service}-record`, {
        alias: {
          evaluateTargetHealth: true,
          name: `REPLACE_WITH_ALB_DNS_${service}`, // consumed from ALB outputs
          zoneId: 'REPLACE_WITH_ALB_ZONE_ID', // consumed from ALB outputs
        },
        name: fqdn,
        ttl: 300,
        type: 'A', // or 'CNAME' depending on target
        zoneId: zone.zoneId,
      });
      this.serviceFqdns[String(service)] = fqdn;
    });

    // 5️⃣ Outputs
    new TerraformOutput(this, `${props.env}_hosted_zone_id`, {
      value: this.hostedZoneId,
    });
    new TerraformOutput(this, `${props.env}_certificate_arn`, {
      value: this.certificateArn,
    });
    props.services.forEach((s) => {
      new TerraformOutput(this, `${props.env}_${s}_fqdn`, {
        value: this.serviceFqdns[String(s)],
      });
    });
  }
}
