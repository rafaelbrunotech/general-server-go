import { App } from 'cdktf';

import {
  DatabaseStack,
  EcsStack,
  MessageBrokerStack,
  NetworkStack,
  ObservabilityStack,
  SecretsStack,
} from '../lib';
import { Environment, remoteState } from '../lib/shared';

const env: Environment = process.env.ENV as Environment || Environment.DEV;

const app = new App();

remoteState(app, env);

new NetworkStack(app, `network-stack-${env}`, { env });
new EcsStack(app, `ecs-stack-${env}`, { env });
new DatabaseStack(app, `database-stack-${env}`, { env });
new SecretsStack(app, `secrets-stack-${env}`, { env });
new MessageBrokerStack(app, `message-broker-stack-${env}`, { env });
new ObservabilityStack(app, `observability-stack-${env}`, { env });

app.synth();
