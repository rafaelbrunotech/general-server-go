import { Environment } from './environment.enum';
import { Region } from './region.enum';

export interface EnvironmentOptions {
  env: Environment;
  region?: Region;
  user?: string;
}
