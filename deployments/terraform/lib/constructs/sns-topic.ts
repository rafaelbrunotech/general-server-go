import { TerraformOutput } from 'cdktf';
import { Construct } from 'constructs';

import { KmsKey } from '../../.gen/providers/aws/kms-key';
import { SnsTopic } from '../../.gen/providers/aws/sns-topic';

interface SnsTopicProps {
  env: string;
  name: string;
}

export class SnsTopicConstruct extends Construct {
  public readonly topicArn: string;

  constructor(scope: Construct, id: string, props: SnsTopicProps) {
    super(scope, id);

    const kmsKey = new KmsKey(this, `${props.env}-${props.name}-kms`, {
      description: `KMS key for ${props.env}-${props.name} SNS encryption`,
    });

    const topic = new SnsTopic(this, `${props.env}-sns-${props.name}`, {
      kmsMasterKeyId: kmsKey.id,
      name: `${props.env}-${props.name}`,
    });

    this.topicArn = topic.arn;
    new TerraformOutput(this, `${props.env}_${props.name}_sns_arn`, {
      value: this.topicArn,
    });
  }
}
