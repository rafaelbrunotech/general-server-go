import { Environment } from '../models/environment.enum';
import { Region } from '../models/region.enum';

export const getStackName = (
  componentName: string,
  environment: Environment,
  region?: Region,
) => {
  if (region) {
    return `${environment}-${componentName}-${region}`;
  } else {
    return `${environment}-${componentName}`;
  }
};
