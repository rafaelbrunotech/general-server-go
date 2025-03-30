import { App } from 'cdktf';

import { AppStack } from '../lib';
import { Environment, getStackName } from '../lib/shared';

const app = new App();

const createEnvironment = (environment: Environment) => {
  new AppStack(app, getStackName('app', environment), { environment });
};

createEnvironment(Environment.DEV);
createEnvironment(Environment.PROD);

app.synth();
