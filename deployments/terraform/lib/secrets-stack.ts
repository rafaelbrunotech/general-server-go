import { TerraformOutput, TerraformStack } from 'cdktf';
import { Construct } from 'constructs';

import { SecretsManagerConstruct } from './constructs/secrets-manager';
import { EnvironmentOptions } from './shared/models/environment.interface';

export interface SecretsStackProps extends EnvironmentOptions {
  rotationLambdaArns: Record<string, string>; // mapping env -> rotation Lambda
}

export class SecretsStack extends TerraformStack {
  constructor(scope: Construct, id: string, props: SecretsStackProps) {
    super(scope, id);

    // -------------------------
    // 1️⃣ Database Secret
    // -------------------------
    const dbSecret = new SecretsManagerConstruct(this, `${props.env}-db-secret`, {
      env: props.env,
      name: 'db-credentials',
      rotationDays: 30,
      rotationLambdaArn: props.rotationLambdaArns[props.env],
      secretString: {
        password: 'SuperSecret123!', // can be generated dynamically
        username: 'admin',
      },
    });

    new TerraformOutput(this, `${props.env}_db_secret_arn`, {
      value: dbSecret.secretArn,
    });

    // -------------------------
    // 2️⃣ Optional: Other Service Secrets
    // Example: API key for 3rd party service
    // -------------------------
    const apiKeySecret = new SecretsManagerConstruct(this, `${props.env}-api-key-secret`, {
      env: props.env,
      name: 'service-api-key',
      rotationDays: 30,
      rotationLambdaArn: props.rotationLambdaArns[props.env],
      secretString: {
        apiKey: 'SuperSecretApiKey', // replace with CI-generated or random value
      },
    });

    new TerraformOutput(this, `${props.env}_api_key_secret_arn`, {
      value: apiKeySecret.secretArn,
    });

    // -------------------------
    // Additional secrets for other microservices can be added here following the same pattern
    // -------------------------
  }
}
