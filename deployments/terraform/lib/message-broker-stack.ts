import { TerraformOutput, TerraformStack } from 'cdktf';
import { Construct } from 'constructs';

import { SnsTopicConstruct } from './constructs/sns-topic';
import { SnsTopicSubscriptionConstruct } from './constructs/sns-topic-subscription';
import { SqsQueueConstruct } from './constructs/sqs-queue';
import { EnvironmentOptions } from './shared/models/environment.interface';

export interface MessageBrokerStackProps extends EnvironmentOptions {
  services: string[]; // list of service names
}

export class MessageBrokerStack extends TerraformStack {
  public readonly queues: Record<string, string> = {};
  public readonly subscriptions: Record<string, string> = {};
  public readonly topics: Record<string, string> = {};

  constructor(scope: Construct, id: string, props: MessageBrokerStackProps) {
    super(scope, id);

    for (const service of props.services) {
      // DLQ for the service
      const dlq = new SqsQueueConstruct(this, `${props.env}-${service}-dlq`, {
        env: props.env,
        name: `${service}-dlq`,
      });

      // Main queue
      const queue = new SqsQueueConstruct(this, `${props.env}-${service}-queue`, {
        deadLetterArn: dlq.queueArn,
        env: props.env,
        name: `${service}-queue`,
      });

      // Topic
      const topic = new SnsTopicConstruct(this, `${props.env}-${service}-topic`, {
        env: props.env,
        name: `${service}-topic`,
      });

      // Subscribe queue to topic
      const subscription = new SnsTopicSubscriptionConstruct(
        this,
        `${props.env}-${service}-subscription`,
        {
          env: props.env,
          queueArn: queue.queueArn,
          topicArn: topic.topicArn,
        },
      );

      // Outputs
      this.topics[String(service)] = topic.topicArn;
      this.queues[String(service)] = queue.queueArn;
      this.subscriptions[String(service)] = subscription.subscriptionArn;

      new TerraformOutput(this, `${props.env}_${service}_queue_url`, {
        value: queue.queueUrl,
      });
      new TerraformOutput(this, `${props.env}_${service}_queue_arn`, {
        value: queue.queueArn,
      });
      new TerraformOutput(this, `${props.env}_${service}_topic_arn`, {
        value: topic.topicArn,
      });
      new TerraformOutput(this, `${props.env}_${service}_subscription_arn`, {
        value: subscription.subscriptionArn,
      });
    }
  }
}
