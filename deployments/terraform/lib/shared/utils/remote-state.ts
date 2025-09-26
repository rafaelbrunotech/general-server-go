import { S3Backend } from 'cdktf';
import { Construct } from 'constructs';

export const remoteState = (scope: Construct, env: string) =>
  new S3Backend(scope, {
    bucket: 'platform-terraform-state',
    dynamodbTable: 'platform-tf-locks',
    encrypt: true,
    key: `${env}/terraform.tfstate`,
    region: 'us-east-1',
  });
