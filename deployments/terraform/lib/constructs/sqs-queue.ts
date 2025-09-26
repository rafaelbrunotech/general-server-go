import { TerraformOutput } from 'cdktf';
import { Construct } from 'constructs';

import { KmsKey } from '../../.gen/providers/aws/kms-key';
import { SqsQueue } from '../../.gen/providers/aws/sqs-queue';

interface SqsQueueProps {
  deadLetterArn?: string;
  env: string;
  name: string;
}

export class SqsQueueConstruct extends Construct {
  public readonly queueArn: string;
  public readonly queueUrl: string;

  constructor(scope: Construct, id: string, props: SqsQueueProps) {
    super(scope, id);

    const kmsKey = new KmsKey(this, `${props.env}-${props.name}-kms`, {
      description: `KMS key for ${props.env}-${props.name} SQS encryption`,
    });

    const queue = new SqsQueue(this, `${props.env}-sqs-${props.name}`, {
      kmsMasterKeyId: kmsKey.id,
      name: `${props.env}-${props.name}`,
      redrivePolicy: props.deadLetterArn
        ? JSON.stringify({
            deadLetterTargetArn: props.deadLetterArn,
            maxReceiveCount: 5,
          })
        : undefined,
    });

    this.queueUrl = queue.url;
    this.queueArn = queue.arn;

    new TerraformOutput(this, `${props.env}_${props.name}_sqs_url`, {
      value: this.queueUrl,
    });
    new TerraformOutput(this, `${props.env}_${props.name}_sqs_arn`, {
      value: this.queueArn,
    });
  }
}
