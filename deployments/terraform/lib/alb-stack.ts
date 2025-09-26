import { TerraformOutput, TerraformStack } from 'cdktf';
import { Construct } from 'constructs';

import { Alb } from '../.gen/providers/aws/alb';
import { AlbListener } from '../.gen/providers/aws/alb-listener';
import { AlbTargetGroup } from '../.gen/providers/aws/alb-target-group';
import { EnvironmentOptions } from './shared/models/environment.interface';

export interface AlbStackProps extends EnvironmentOptions {
  albSecurityGroupId: string;
  publicSubnets: string[];
  vpcId: string;
}

export class AlbStack extends TerraformStack {
  public readonly albDnsName: string;
  public readonly targetGroupArn: string;

  constructor(scope: Construct, id: string, props: AlbStackProps) {
    super(scope, id);

    // 2️⃣ ALB
    const alb = new Alb(this, `${props.env}-alb`, {
      enableDeletionProtection: false,
      internal: false,
      loadBalancerType: 'application',
      name: `${props.env}-alb`,
      securityGroups: [props.albSecurityGroupId],
      subnets: props.publicSubnets,
      tags: { Environment: props.env, Name: `${props.env}-alb` },
    });

    // 3️⃣ Target Group (HTTP 80)
    const targetGroup = new AlbTargetGroup(this, `${props.env}-tg`, {
      healthCheck: {
        healthyThreshold: 2,
        interval: 30,
        matcher: '200-399',
        path: '/',
        timeout: 5,
        unhealthyThreshold: 2,
      },
      name: `${props.env}-tg`,
      port: 80,
      protocol: 'HTTP',
      tags: { Environment: props.env, Name: `${props.env}-tg` },
      targetType: 'ip',
      vpcId: props.vpcId,
    });

    // 4️⃣ Listener (HTTP)
    new AlbListener(this, `${props.env}-listener`, {
      defaultAction: [
        {
          targetGroupArn: targetGroup.arn,
          type: 'forward',
        },
      ],
      loadBalancerArn: alb.arn,
      port: 80,
      protocol: 'HTTP',
    });

    // Outputs
    this.albDnsName = alb.dnsName;
    this.targetGroupArn = targetGroup.arn;

    new TerraformOutput(this, `${props.env}_alb_dns_name`, { value: this.albDnsName });
    new TerraformOutput(this, `${props.env}_alb_target_group_arn`, { value: this.targetGroupArn });
  }
}
