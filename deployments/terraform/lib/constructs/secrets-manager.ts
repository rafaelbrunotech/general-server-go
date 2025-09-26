import { TerraformOutput } from 'cdktf';
import { Construct } from 'constructs';

import { SecretsmanagerSecret } from '../../.gen/providers/aws/secretsmanager-secret';
import { SecretsmanagerSecretRotation } from '../../.gen/providers/aws/secretsmanager-secret-rotation';
import { SecretsmanagerSecretVersion } from '../../.gen/providers/aws/secretsmanager-secret-version';

interface SecretConstructProps {
  env: string;
  name: string; // short name for secret, e.g., 'db-credentials'
  rotationDays?: number; // rotation frequency
  rotationLambdaArn?: string; // optional Lambda ARN for automatic rotation
  secretString?: Record<string, string>; // optional initial values
  tags?: Record<string, string>;
}

export class SecretsManagerConstruct extends Construct {
  public readonly secretArn: string;
  public readonly secretId: string;

  constructor(scope: Construct, id: string, props: SecretConstructProps) {
    super(scope, id);

    const secret = new SecretsmanagerSecret(this, `${props.env}-${props.name}`, {
      description: `Secret for ${props.name} in ${props.env}`,
      name: `${props.env}-${props.name}`,
      recoveryWindowInDays: 7,
      tags: { Environment: props.env, ...(props.tags || {}) },
    });

    if (props.secretString) {
      new SecretsmanagerSecretVersion(this, `${props.env}-${props.name}-version`, {
        secretId: secret.id,
        secretString: JSON.stringify(props.secretString),
      });
    }

    if (props.rotationLambdaArn && props.rotationDays) {
      new SecretsmanagerSecretRotation(this, `${props.env}-${props.name}-rotation`, {
        rotationLambdaArn: props.rotationLambdaArn,
        rotationRules: { automaticallyAfterDays: props.rotationDays },
        secretId: secret.id,
      });
    }

    this.secretArn = secret.arn;
    this.secretId = secret.id;

    new TerraformOutput(scope, `${props.env}_${props.name}_secret_arn`, {
      value: this.secretArn,
    });
  }
}
