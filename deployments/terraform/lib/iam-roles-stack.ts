import { TerraformStack } from 'cdktf';
import { Construct } from 'constructs';

import { IamRole } from '../.gen/providers/aws/iam-role';
import { IamRolePolicy } from '../.gen/providers/aws/iam-role-policy';
import { EnvironmentOptions } from './shared/models/environment.interface';

interface IamRolesStackProps extends EnvironmentOptions {}

export class IamRolesStack extends TerraformStack {
  public readonly ecsTaskRoleArn: string;
  public readonly lambdaRoleArn: string;

  constructor(scope: Construct, id: string, props: IamRolesStackProps) {
    super(scope, id);

    // ECS Task Role
    const ecsTaskRole = new IamRole(this, `${props.env}-ecs-task-role`, {
      assumeRolePolicy: JSON.stringify({
        Statement: [
          {
            Action: 'sts:AssumeRole',
            Effect: 'Allow',
            Principal: { Service: 'ecs-tasks.amazonaws.com' },
          },
        ],
        Version: '2012-10-17',
      }),
      name: `${props.env}-ecs-task-role`,
      tags: { Environment: props.env },
    });

    new IamRolePolicy(this, `${props.env}-ecs-task-policy`, {
      name: `${props.env}-ecs-task-policy`,
      policy: JSON.stringify({
        Statement: [
          {
            Action: ['secretsmanager:GetSecretValue', 'sqs:*', 'sns:*', 'logs:*'],
            Effect: 'Allow',
            Resource: '*',
          },
        ],
        Version: '2012-10-17',
      }),
      role: ecsTaskRole.name,
    });

    this.ecsTaskRoleArn = ecsTaskRole.arn;

    // Lambda Role
    const lambdaRole = new IamRole(this, `${props.env}-lambda-role`, {
      assumeRolePolicy: JSON.stringify({
        Statement: [
          {
            Action: 'sts:AssumeRole',
            Effect: 'Allow',
            Principal: { Service: 'lambda.amazonaws.com' },
          },
        ],
        Version: '2012-10-17',
      }),
      name: `${props.env}-lambda-role`,
      tags: { Environment: props.env },
    });

    new IamRolePolicy(this, `${props.env}-lambda-policy`, {
      name: `${props.env}-lambda-policy`,
      policy: JSON.stringify({
        Statement: [
          {
            Action: ['secretsmanager:GetSecretValue', 'logs:*', 'sqs:*', 'sns:*'],
            Effect: 'Allow',
            Resource: '*',
          },
        ],
        Version: '2012-10-17',
      }),
      role: lambdaRole.name,
    });

    this.lambdaRoleArn = lambdaRole.arn;
  }
}
