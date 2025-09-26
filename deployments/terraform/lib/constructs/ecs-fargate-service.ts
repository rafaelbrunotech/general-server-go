import { TerraformOutput } from 'cdktf';
import { Construct } from 'constructs';

import { CloudwatchLogGroup } from '../../.gen/providers/aws/cloudwatch-log-group';
import { EcsService } from '../../.gen/providers/aws/ecs-service';
import { EcsTaskDefinition } from '../../.gen/providers/aws/ecs-task-definition';
import { IamRole } from '../../.gen/providers/aws/iam-role';

export interface FargateServiceProps {
  assignPublicIp?: boolean;
  clusterArn: string;
  containerImage: string;
  cpu?: number;
  desiredCount?: number;
  env: string;
  memory?: number;
  name: string;
  port?: number;
  secrets?: Record<string, string>;
  securityGroups: string[];
  subnets: string[];
}

export class FargateServiceConstruct extends Construct {
  public readonly serviceArn: string;

  constructor(scope: Construct, id: string, props: FargateServiceProps) {
    super(scope, id);

    // CloudWatch Logs
    const logGroup = new CloudwatchLogGroup(
      this,
      `${props.env}-${props.name}-log-group`,
      {
        name: `/ecs/${props.env}/${props.name}`,
        retentionInDays: 30,
      },
    );

    const taskRole = new IamRole(this, `${props.env}-task-role-${props.name}`, {
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
      name: `${props.env}-${props.name}-task-role`,
    });

    const taskDef = new EcsTaskDefinition(
      this,
      `${props.env}-task-def-${props.name}`,
      {
        containerDefinitions: JSON.stringify([
          {
            environment: [],
            essential: true,
            image: props.containerImage,
            logConfiguration: {
              logDriver: 'awslogs',
              options: {
                'awslogs-group': logGroup.name,
                'awslogs-region': 'us-east-1', // replace with your region
                'awslogs-stream-prefix': props.name,
              },
            },
            name: props.name,
            portMappings: props.port ? [{ containerPort: props.port }] : [],
            secrets: props.secrets
              ? Object.entries(props.secrets).map(([k, v]) => ({
                  name: k,
                  valueFrom: v,
                }))
              : [],
          },
        ]),
        cpu: props.cpu?.toString() || '512',
        executionRoleArn: taskRole.arn,
        family: `${props.env}-${props.name}`,
        memory: props.memory?.toString() || '1024',
        networkMode: 'awsvpc',
        requiresCompatibilities: ['FARGATE'],
      },
    );

    const service = new EcsService(this, `${props.env}-ecs-service-${props.name}`, {
      cluster: props.clusterArn,
      desiredCount: props.desiredCount || 1,
      launchType: 'FARGATE',
      name: `${props.env}-${props.name}`,
      networkConfiguration: {
        assignPublicIp: props.assignPublicIp || false,
        securityGroups: props.securityGroups,
        subnets: props.subnets,
      },
      tags: { Environment: props.env, Service: props.name },
      taskDefinition: taskDef.arn,
    });

    this.serviceArn = service.id;

    new TerraformOutput(this, `${props.env}_${props.name}_service_arn`, {
      value: this.serviceArn,
    });
  }
}
