import { TerraformOutput, TerraformStack } from 'cdktf';
import { Construct } from 'constructs';

import { AcmCertificate } from '../.gen/providers/aws/acm-certificate';
import { EnvironmentOptions } from './shared/models/environment.interface';

export interface CertificateStackProps extends EnvironmentOptions {
  domainName: string; // e.g., rafaelbruno.com
  hostedZoneId: string; // Route53 Hosted Zone ID
  subdomains: string[]; // e.g., ["frontend", "backend", "api"]
}

export class CertificateStack extends TerraformStack {
  public readonly certificateArns: Record<string, string> = {};

  constructor(scope: Construct, id: string, props: CertificateStackProps) {
    super(scope, id);

    // Root domain certificate
    const rootCert = new AcmCertificate(this, `${props.env}-root-cert`, {
      domainName: props.domainName,
      tags: { Environment: props.env, Name: `${props.env}-root-cert` },
      validationMethod: 'DNS',
    });

    new TerraformOutput(this, `${props.env}_root_cert_arn`, { value: rootCert.arn });
    this.certificateArns['root'] = rootCert.arn;

    // Subdomain certificates
    props.subdomains.forEach((sub) => {
      const fqdn = `${sub}.${props.domainName}`;
      const cert = new AcmCertificate(this, `${props.env}-${sub}-cert`, {
        domainName: fqdn,
        tags: { Environment: props.env, Name: `${props.env}-${sub}-cert` },
        validationMethod: 'DNS',
      });

      new TerraformOutput(this, `${props.env}_${sub}_cert_arn`, { value: cert.arn });
      this.certificateArns[String(sub)] = cert.arn;
    });
  }
}
