import { Environment } from './environment.enum';
import { Region } from './region.enum';

export interface EnvironmentOptions {
  environment: Environment;
  region?: Region;
  user?: string;
}
