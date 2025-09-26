import { TerraformOutput, TerraformStack } from 'cdktf';
import { Construct } from 'constructs';

import { DbInstance } from '../.gen/providers/aws/db-instance';
import { DbParameterGroup } from '../.gen/providers/aws/db-parameter-group';
import { DbSubnetGroup } from '../.gen/providers/aws/db-subnet-group';
import { SecretsManagerConstruct } from './constructs';
import { Environment, EnvironmentOptions } from './shared/models';

export interface DatabaseStackProps extends EnvironmentOptions {
  privateSubnetIds: string[];
  rdsSecurityGroupId: string; // from platform SG stack
  rotationLambdaArns: Record<string, string>;
}

export class DatabaseStack extends TerraformStack {
  public readonly readReplicaEndpoints: string[] = [];

  constructor(scope: Construct, id: string, props: DatabaseStackProps) {
    super(scope, id);

    // 1️⃣ Subnet Group
    const dbSubnetGroup = new DbSubnetGroup(this, `${props.env}-db-subnet-group`, {
      name: `${props.env}-db-subnet-group`,
      subnetIds: props.privateSubnetIds,
      tags: {
        Environment: props.env,
        Name: `${props.env}-db-subnet-group`,
      },
    });

    const rdsParameterGroup = new DbParameterGroup(this, `${props.env}-db-parameter-group`, {
      description: `Parameter group for ${props.env} RDS`,
      family: 'postgres13', // adjust for your engine version
      name: `${props.env}-db-parameter-group`,
      parameter: [
        { name: 'rds.force_ssl', value: '1' }, // enforce TLS
      ],
      tags: { Environment: props.env, Name: `${props.env}-db-parameter-group` },
    });

    // 2️⃣ Secrets Manager
    const dbSecret = new SecretsManagerConstruct(this, 'dbSecret', {
      env: props.env,
      name: 'db-credentials',
      rotationDays: 30,
      rotationLambdaArn: props.rotationLambdaArns[props.env],
      secretString: {
        password: 'SuperSecret123!', // optionally use CI-provided or generated
        username: 'admin',
      },
    });

    // 3️⃣ RDS Primary Instance
    const rdsInstance = new DbInstance(this, `${props.env}-rds`, {
      allocatedStorage: 20,
      backupRetentionPeriod: 30,
      dbName: `${props.env}-appdb`,
      dbSubnetGroupName: dbSubnetGroup.name,
      engine: 'postgres',
      instanceClass: props.env === Environment.PROD ? 'db.t3.medium' : 'db.t3.small',
      multiAz: props.env === Environment.PROD,
      parameterGroupName: rdsParameterGroup.name,
      password: 'SuperSecret123!', // TODO: optionally use CI-provided or generated
      publiclyAccessible: false,
      skipFinalSnapshot: props.env !== Environment.PROD,
      storageEncrypted: true, // TLS/encryption
      tags: {
        Environment: props.env,
        Name: `${props.env}-rds`,
      },
      username: 'admin', // TODO: optionally use CI-provided or generated
      vpcSecurityGroupIds: [props.rdsSecurityGroupId],
    });

    // 4️⃣ Read replicas for prod
    if (props.env === Environment.PROD) {
      for (let i = 1; i <= 3; i++) {
        const replica = new DbInstance(this, `${props.env}-rds-replica-${i}`, {
          dbSubnetGroupName: dbSubnetGroup.name,
          engine: 'postgres',
          instanceClass: 'db.t3.medium',
          parameterGroupName: rdsParameterGroup.name,
          publiclyAccessible: false,
          replicateSourceDb: rdsInstance.id,
          skipFinalSnapshot: true,
          tags: {
            Environment: props.env,
            Name: `${props.env}-rds-replica-${i}`,
          },
          vpcSecurityGroupIds: [props.rdsSecurityGroupId],
        });

        this.readReplicaEndpoints.push(replica.endpoint);

        new TerraformOutput(this, `${props.env}_rds_read_endpoint_${i}`, {
          value: replica.endpoint,
        });
      }
    }

    new TerraformOutput(this, `${props.env}_rds_endpoint`, {
      value: rdsInstance.endpoint,
    });
    new TerraformOutput(this, `${props.env}_db_secret_arn`, {
      value: dbSecret.secretArn,
    });
  }
}
