import { TerraformOutput, TerraformStack } from 'cdktf';
import { Construct } from 'constructs';

import { Eip } from '../.gen/providers/aws/eip';
import { NatGateway } from '../.gen/providers/aws/nat-gateway';
import { Route } from '../.gen/providers/aws/route';
import { RouteTable } from '../.gen/providers/aws/route-table';
import { RouteTableAssociation } from '../.gen/providers/aws/route-table-association';
import { EnvironmentOptions } from './shared/models/environment.interface';

export interface NatGatewayStackProps extends EnvironmentOptions {
  privateSubnets: string[];
  publicSubnetId: string; // pick one public subnet to host the NAT
  vpcId: string;
}

export class NatGatewayStack extends TerraformStack {
  constructor(scope: Construct, id: string, props: NatGatewayStackProps) {
    super(scope, id);

    // Elastic IP for NAT
    const eip = new Eip(this, `${props.env}-nat-eip`, {
      tags: { Environment: props.env, Name: `${props.env}-nat-eip` },
      vpc: true,
    });

    // NAT Gateway
    const natGateway = new NatGateway(this, `${props.env}-nat-gateway`, {
      allocationId: eip.id,
      subnetId: props.publicSubnetId,
      tags: { Environment: props.env, Name: `${props.env}-nat` },
    });

    // Private Route Table
    const privateRt = new RouteTable(this, `${props.env}-private-rt`, {
      tags: { Environment: props.env, Name: `${props.env}-private-rt` },
      vpcId: props.vpcId,
    });

    new Route(this, `${props.env}-private-route`, {
      destinationCidrBlock: '0.0.0.0/0',
      natGatewayId: natGateway.id,
      routeTableId: privateRt.id,
    });

    props.privateSubnets.forEach((subnetId, i) => {
      new RouteTableAssociation(this, `${props.env}-private-rt-assoc-${i}`, {
        routeTableId: privateRt.id,
        subnetId,
      });
    });

    // Outputs
    new TerraformOutput(this, `${props.env}_nat_gateway_id`, {
      value: natGateway.id,
    });
    new TerraformOutput(this, `${props.env}_private_rt_id`, { value: privateRt.id });
  }
}
