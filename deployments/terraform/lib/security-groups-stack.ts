import { TerraformOutput, TerraformStack } from 'cdktf';
import { Construct } from 'constructs';

import { SecurityGroup } from '../.gen/providers/aws/security-group';
import { EnvironmentOptions } from './shared/models/environment.interface';

interface SecurityGroupsStackProps extends EnvironmentOptions {
  vpcId: string;
}

export class SecurityGroupsStack extends TerraformStack {
  public readonly backendSecurityGroup: SecurityGroup;
  public readonly leadsSecurityGroup: SecurityGroup;
  public readonly optimizerSecurityGroup: SecurityGroup;
  public readonly rdsSecurityGroup: SecurityGroup;
  public readonly replyGuySecurityGroup: SecurityGroup;

  constructor(scope: Construct, id: string, props: SecurityGroupsStackProps) {
    super(scope, id);

    // 1️⃣ RDS SG
    this.rdsSecurityGroup = new SecurityGroup(
      this,
      `${props.env}-rds-sg`,
      {
        description: 'Allow access to RDS from services',
        egress: [
          { cidrBlocks: ['0.0.0.0/0'], fromPort: 0, protocol: '-1', toPort: 0 },
        ],
        ingress: [], // attached later via references
        name: `${props.env}-rds-sg`,
        tags: { Environment: props.env },
        vpcId: props.vpcId,
      },
    );

    // 2️⃣ Backend SG
    this.backendSecurityGroup = new SecurityGroup(
      this,
      `${props.env}-backend-sg`,
      {
        description: 'Main backend ECS service SG',
        egress: [
          { cidrBlocks: ['0.0.0.0/0'], fromPort: 0, protocol: '-1', toPort: 0 },
        ],
        ingress: [
          { cidrBlocks: ['0.0.0.0/0'], fromPort: 443, protocol: 'tcp', toPort: 443 }, // public HTTPS
          {
            fromPort: 5432,
            protocol: 'tcp',
            securityGroups: [this.rdsSecurityGroup.id],
            toPort: 5432,
          }, // access RDS
        ],
        name: `${props.env}-backend-sg`,
        tags: { Environment: props.env },
        vpcId: props.vpcId,
      },
    );

    // 3️⃣ Reply Guy SG
    this.replyGuySecurityGroup = new SecurityGroup(
      this,
      `${props.env}-reply-guy-sg`,
      {
        description: 'Reply Guy ECS service SG',
        egress: [
          { cidrBlocks: ['0.0.0.0/0'], fromPort: 0, protocol: '-1', toPort: 0 },
        ],
        ingress: [
          {
            fromPort: 443,
            protocol: 'tcp',
            securityGroups: [this.backendSecurityGroup.id],
            toPort: 443,
          },
          {
            fromPort: 5432,
            protocol: 'tcp',
            securityGroups: [this.rdsSecurityGroup.id],
            toPort: 5432,
          },
        ],
        name: `${props.env}-reply-guy-sg`,
        tags: { Environment: props.env },
        vpcId: props.vpcId,
      },
    );

    // 4️⃣ Leads SG
    this.leadsSecurityGroup = new SecurityGroup(
      this,
      `${props.env}-leads-sg`,
      {
        description: 'Leads ECS service SG',
        egress: [
          { cidrBlocks: ['0.0.0.0/0'], fromPort: 0, protocol: '-1', toPort: 0 },
        ],
        ingress: [
          {
            fromPort: 443,
            protocol: 'tcp',
            securityGroups: [this.backendSecurityGroup.id],
            toPort: 443,
          },
          {
            fromPort: 5432,
            protocol: 'tcp',
            securityGroups: [this.rdsSecurityGroup.id],
            toPort: 5432,
          },
        ],
        name: `${props.env}-leads-sg`,
        tags: { Environment: props.env },
        vpcId: props.vpcId,
      },
    );

    // 5️⃣ Optimizer SG
    this.optimizerSecurityGroup = new SecurityGroup(
      this,
      `${props.env}-optimizer-sg`,
      {
        description: 'Lambda optimizer SG',
        egress: [
          { cidrBlocks: ['0.0.0.0/0'], fromPort: 0, protocol: '-1', toPort: 0 },
        ],
        ingress: [
          {
            fromPort: 443,
            protocol: 'tcp',
            securityGroups: [this.backendSecurityGroup.id],
            toPort: 443,
          },
        ],
        name: `${props.env}-optimizer-sg`,
        tags: { Environment: props.env },
        vpcId: props.vpcId,
      },
    );

    // 6️⃣ Outputs for remote state consumption
    new TerraformOutput(this, `${props.env}_rds_sg`, {
      value: this.rdsSecurityGroup.id,
    });
    new TerraformOutput(this, `${props.env}_backend_sg`, {
      value: this.backendSecurityGroup.id,
    });
    new TerraformOutput(this, `${props.env}_reply_guy_sg`, {
      value: this.replyGuySecurityGroup.id,
    });
    new TerraformOutput(this, `${props.env}_leads_sg`, {
      value: this.leadsSecurityGroup.id,
    });
    new TerraformOutput(this, `${props.env}_optimizer_sg`, {
      value: this.optimizerSecurityGroup.id,
    });
  }
}
