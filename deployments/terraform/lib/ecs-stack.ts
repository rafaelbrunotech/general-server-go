import { TerraformOutput, TerraformStack } from 'cdktf';
import { Construct } from 'constructs';

import { EcsClusterConstruct } from './constructs/ecs-cluster';
// import { Ec2ServiceConstruct } from './constructs/ecs-ec2-service';
// import { FargateServiceConstruct } from './constructs/ecs-fargate-service';
import { EnvironmentOptions } from './shared/models/environment.interface';

export interface EcsStackProps extends EnvironmentOptions {
  privateSubnets: string[];
}

export class EcsStack extends TerraformStack {
  public readonly ec2ClusterArn: string;
  public readonly fargateClusterArn: string;

  constructor(scope: Construct, id: string, props: EcsStackProps) {
    super(scope, id);

    // Create Clusters
    const fargateCluster = new EcsClusterConstruct(
      this,
      `${props.env}-fargate-cluster`,
      {
        clusterType: 'fargate',
        env: props.env,
      },
    );

    const ec2Cluster = new EcsClusterConstruct(this, `${props.env}-ec2-cluster`, {
      clusterType: 'ec2',
      env: props.env,
      instanceType: 't3.medium',
      subnets: props.privateSubnets,
    });

    this.fargateClusterArn = fargateCluster.cluster.arn;
    this.ec2ClusterArn = ec2Cluster.cluster.arn;

    new TerraformOutput(this, `${props.env}_fargate_cluster_arn`, {
      value: fargateCluster.cluster.arn,
    });
    new TerraformOutput(this, `${props.env}_ec2_cluster_arn`, {
      value: ec2Cluster.cluster.arn,
    });
  }
}
