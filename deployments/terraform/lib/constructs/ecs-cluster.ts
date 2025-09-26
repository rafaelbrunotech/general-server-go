import { Construct } from 'constructs';

import { AutoscalingGroup } from '../../.gen/providers/aws/autoscaling-group';
import { EcsCluster } from '../../.gen/providers/aws/ecs-cluster';
import { LaunchConfiguration } from '../../.gen/providers/aws/launch-configuration';

export interface EcsClusterConstructProps {
  clusterType: 'ec2' | 'fargate';
  env: string;
  instanceType?: string; // only for EC2
  subnets?: string[];
  tags?: Record<string, string>;
}

export class EcsClusterConstruct extends Construct {
  public readonly cluster: EcsCluster;

  constructor(scope: Construct, id: string, props: EcsClusterConstructProps) {
    super(scope, id);

    const clusterName = `${props.env}-${props.clusterType}-cluster`;

    this.cluster = new EcsCluster(this, clusterName, {
      name: clusterName,
      setting: [
        {
          name: 'containerInsights',
          value: 'enabled',
        },
      ],
      tags: {
        Environment: props.env,
        Type: props.clusterType,
        ...props.tags,
      },
    });

    // EC2 only: create single-instance per AZ capacity provider
    if (props.clusterType === 'ec2' && props.subnets && props.instanceType) {
      const lc = new LaunchConfiguration(this, `${clusterName}-lc`, {
        iamInstanceProfile: '', // attach proper ECS instance profile
        imageId: 'ami-ecs-optimized', // replace with correct ECS-optimized AMI
        instanceType: props.instanceType,
        name: `${clusterName}-lc`,
        securityGroups: [], // optional
      });

      new AutoscalingGroup(this, `${clusterName}-asg`, {
        desiredCapacity: 1,
        launchConfiguration: lc.name,
        maxSize: 1,
        minSize: 1,
        name: `${clusterName}-asg`,
        tag: [
          { key: 'Environment', propagateAtLaunch: true, value: props.env },
          { key: 'Type', propagateAtLaunch: true, value: props.clusterType },
        ],
        vpcZoneIdentifier: props.subnets,
      });
    }
  }
}
