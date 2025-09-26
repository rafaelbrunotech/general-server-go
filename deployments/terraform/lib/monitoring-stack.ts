import { TerraformOutput, TerraformStack } from 'cdktf';
import { Construct } from 'constructs';

import { CloudwatchLogGroup } from '../.gen/providers/aws/cloudwatch-log-group';
import { CloudwatchMetricAlarm } from '../.gen/providers/aws/cloudwatch-metric-alarm';
import { SnsTopic } from '../.gen/providers/aws/sns-topic';
import { EnvironmentOptions } from './shared/models/environment.interface';

export interface MonitoringStackProps extends EnvironmentOptions {
  services: string[]; // ["frontend", "backend", "optimizer", "reply-guy", "leads", "message"]
}

export class MonitoringStack extends TerraformStack {
  public readonly logGroupArns: Record<string, string> = {};
  public readonly snsTopicArn: string;

  constructor(scope: Construct, id: string, props: MonitoringStackProps) {
    super(scope, id);

    // SNS Topic for alerts
    const alertTopic = new SnsTopic(this, `${props.env}-alerts-topic`, {
      name: `${props.env}-alerts`,
    });
    this.snsTopicArn = alertTopic.arn;

    new TerraformOutput(this, `${props.env}_alerts_sns_arn`, {
      value: this.snsTopicArn,
    });

    // CloudWatch log groups for each service
    props.services.forEach((service) => {
      const logGroup = new CloudwatchLogGroup(
        this,
        `${props.env}-${service}-log-group`,
        {
          name: `/${props.env}/${service}`,
          retentionInDays: 30,
          tags: { Environment: props.env, Service: service },
        },
      );

      this.logGroupArns[String(service)] = logGroup.arn;

      new TerraformOutput(this, `${props.env}_${service}_log_group_arn`, {
        value: logGroup.arn,
      });

      // Example: CPU alarm for ECS tasks (can expand to memory, RDS, etc.)
      new CloudwatchMetricAlarm(this, `${props.env}-${service}-cpu-alarm`, {
        alarmActions: [alertTopic.arn],
        alarmDescription: `High CPU utilization for ${service}`,
        alarmName: `${props.env}-${service}-high-cpu`,
        comparisonOperator: 'GreaterThanThreshold',
        evaluationPeriods: 2,
        metricName: 'CPUUtilization',
        namespace: 'AWS/ECS',
        period: 60,
        statistic: 'Average',
        threshold: 80,
      });
    });
  }
}
