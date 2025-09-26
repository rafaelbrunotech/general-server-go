import { TerraformOutput, TerraformStack } from 'cdktf';
import { Construct } from 'constructs';

import { AlbTargetGroup } from '../.gen/providers/aws/alb-target-group';
import { AlbTargetGroupAttachment } from '../.gen/providers/aws/alb-target-group-attachment';
import { AutoscalingGroup } from '../.gen/providers/aws/autoscaling-group';
import { LaunchTemplate } from '../.gen/providers/aws/launch-template';
import { EnvironmentOptions } from './shared/models/environment.interface';

export interface AutoscalingStackProps extends EnvironmentOptions {
  albArn: string;
  privateSubnets: string[];
  publicSubnets: string[];
  vpcId: string;
}

export class AutoscalingStack extends TerraformStack {
  public readonly asgArn: string;
  public readonly targetGroupArn: string;

  constructor(scope: Construct, id: string, props: AutoscalingStackProps) {
    super(scope, id);

    // 1️⃣ Launch Template for EC2 instances
    const launchTemplate = new LaunchTemplate(this, `${props.env}-launch-template`, {
      imageId: 'ami-0abcdef1234567890', // replace with latest Amazon ECS optimized AMI
      instanceType: 't3.medium',
      keyName: 'my-key', // optional
      name: `${props.env}-launch-template`,
      vpcSecurityGroupIds: [], // add SGs if needed
    });

    // 2️⃣ Target Group for ALB
    const targetGroup = new AlbTargetGroup(this, `${props.env}-tg`, {
      healthCheck: {
        healthyThreshold: 2,
        interval: 30,
        path: '/',
        protocol: 'HTTP',
        timeout: 5,
        unhealthyThreshold: 2,
      },
      name: `${props.env}-tg`,
      port: 80,
      protocol: 'HTTP',
      targetType: 'instance', // EC2 instances
      vpcId: props.vpcId,
    });
    this.targetGroupArn = targetGroup.arn;

    // 3️⃣ AutoScaling Group
    const asg = new AutoscalingGroup(this, `${props.env}-asg`, {
      desiredCapacity: 2,
      healthCheckType: 'EC2',
      launchTemplate: {
        id: launchTemplate.id,
        version: '$Latest',
      },
      maxSize: 4,
      minSize: 1,
      name: `${props.env}-asg`,
      tag: [
        { key: 'Environment', propagateAtLaunch: true, value: props.env },
        { key: 'Name', propagateAtLaunch: true, value: `${props.env}-asg` },
      ],
      vpcZoneIdentifier: props.publicSubnets,
    });
    this.asgArn = asg.arn;

    // 4️⃣ Attach ASG to Target Group
    new AlbTargetGroupAttachment(this, `${props.env}-asg-tg-attach`, {
      targetGroupArn: targetGroup.arn,
      targetId: asg.name,
    });

    // 5️⃣ Outputs
    new TerraformOutput(this, `${props.env}_asg_arn`, { value: this.asgArn });
    new TerraformOutput(this, `${props.env}_tg_arn`, { value: this.targetGroupArn });
  }
}
