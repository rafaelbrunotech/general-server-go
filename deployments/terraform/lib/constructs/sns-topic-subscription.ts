import { TerraformOutput } from 'cdktf';
import { Construct } from 'constructs';

import { SnsTopicSubscription } from '../../.gen/providers/aws/sns-topic-subscription';

interface SnsTopicSubscriptionProps {
  env: string;
  queueArn: string;
  topicArn: string;
}

export class SnsTopicSubscriptionConstruct extends Construct {
  public readonly subscriptionArn: string;

  constructor(scope: Construct, id: string, props: SnsTopicSubscriptionProps) {
    super(scope, id);

    const subscription = new SnsTopicSubscription(
      this,
      `${props.env}-sns-subscription-${props.topicArn}`,
      {
        endpoint: props.queueArn,
        protocol: 'sqs',
        rawMessageDelivery: true,
        topicArn: props.topicArn,
      },
    );

    this.subscriptionArn = subscription.arn;

    new TerraformOutput(this, `${props.env}_sns_subscription_arn`, {
      value: this.subscriptionArn,
    });
  }
}
