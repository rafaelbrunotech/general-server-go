import { TerraformStack } from 'cdktf';
import { Construct } from 'constructs';

import { EnvironmentOptions } from './shared/models/environment.interface';

export interface AppStackProps extends EnvironmentOptions {}

export class AppStack extends TerraformStack {
  constructor(scope: Construct, id: string, props: AppStackProps) {
    console.log(props);
    super(scope, id);
  }
}
