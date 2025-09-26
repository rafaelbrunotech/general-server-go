import { TerraformStack } from 'cdktf';
import { Construct } from 'constructs';

import { CloudwatchLogGroup } from '../.gen/providers/aws/cloudwatch-log-group';
import { S3Bucket } from '../.gen/providers/aws/s3-bucket';
import { Environment, EnvironmentOptions } from './shared/models';

export interface ObservabilityStackProps extends EnvironmentOptions {
  albLogBucketName?: string; // ALB logs destination
  frontendBucketName?: string; // CloudFront logs destination
  services?: string[]; // list of service names
}

export class ObservabilityStack extends TerraformStack {
  constructor(scope: Construct, id: string, props: ObservabilityStackProps) {
    super(scope, id);

    // 1️⃣ Default retention based on environment
    const retentionDays = props.env === Environment.PROD ? 90 : 30;

    // 2️⃣ Platform-wide log group
    new CloudwatchLogGroup(this, `${props.env}-platform-log-group`, {
      name: `/${props.env}/platform`,
      retentionInDays: retentionDays,
      tags: { Environment: props.env },
    });

    // 3️⃣ Service-specific log groups
    if (props.services) {
      props.services.forEach((service) => {
        new CloudwatchLogGroup(this, `${props.env}-${service}-log-group`, {
          name: `/${props.env}/${service}`,
          retentionInDays: retentionDays,
          tags: { Environment: props.env, Service: service },
        });
      });
    }

    // 4️⃣ S3 buckets for frontend / ALB logs
    if (props.frontendBucketName) {
      new S3Bucket(this, `${props.env}-frontend-log-bucket`, {
        acl: 'private',
        bucket: props.frontendBucketName,
        tags: { Environment: props.env, Purpose: 'CloudFrontLogs' },
        versioning: { enabled: true },
      });
    }

    if (props.albLogBucketName) {
      new S3Bucket(this, `${props.env}-alb-log-bucket`, {
        acl: 'private',
        bucket: props.albLogBucketName,
        tags: { Environment: props.env, Purpose: 'ALBLogs' },
        versioning: { enabled: true },
      });
    }
  }
}
