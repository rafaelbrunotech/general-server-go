import { TerraformOutput, TerraformStack } from 'cdktf';
import { Construct } from 'constructs';

import { InternetGateway } from '../.gen/providers/aws/internet-gateway';
import { Route } from '../.gen/providers/aws/route';
import { RouteTable } from '../.gen/providers/aws/route-table';
import { RouteTableAssociation } from '../.gen/providers/aws/route-table-association';
import { Subnet } from '../.gen/providers/aws/subnet';
import { Vpc } from '../.gen/providers/aws/vpc';
import { EnvironmentOptions } from './shared/models/environment.interface';

export interface NetworkStackProps extends EnvironmentOptions {}

export class NetworkStack extends TerraformStack {
  public readonly internetGatewayId: string;
  public readonly privateSubnets: string[];
  public readonly publicRouteTableId: string;
  public readonly publicSubnets: string[];
  public readonly vpcId: string;

  constructor(scope: Construct, id: string, props: NetworkStackProps) {
    super(scope, id);

    // VPC
    const vpc = new Vpc(this, `${props.env}-vpc`, {
      cidrBlock: '10.0.0.0/16',
      enableDnsHostnames: true,
      enableDnsSupport: true,
      tags: { Environment: props.env, Name: `${props.env}-vpc` },
    });

    // Internet Gateway
    const igw = new InternetGateway(this, `${props.env}-igw`, {
      tags: { Environment: props.env, Name: `${props.env}-igw` },
      vpcId: vpc.id,
    });

    // Public Subnets
    const publicSubnet1 = new Subnet(this, `${props.env}-public-subnet-a`, {
      availabilityZone: 'us-east-1a',
      cidrBlock: '10.0.1.0/24',
      mapPublicIpOnLaunch: true,
      tags: { Environment: props.env, Name: `${props.env}-public-a` },
      vpcId: vpc.id,
    });

    const publicSubnet2 = new Subnet(this, `${props.env}-public-subnet-b`, {
      availabilityZone: 'us-east-1b',
      cidrBlock: '10.0.3.0/24',
      mapPublicIpOnLaunch: true,
      tags: { Environment: props.env, Name: `${props.env}-public-b` },
      vpcId: vpc.id,
    });

    // Private Subnets
    const privateSubnet1 = new Subnet(this, `${props.env}-private-subnet-a`, {
      availabilityZone: 'us-east-1a',
      cidrBlock: '10.0.2.0/24',
      mapPublicIpOnLaunch: false,
      tags: { Environment: props.env, Name: `${props.env}-private-a` },
      vpcId: vpc.id,
    });

    const privateSubnet2 = new Subnet(this, `${props.env}-private-subnet-b`, {
      availabilityZone: 'us-east-1b',
      cidrBlock: '10.0.4.0/24',
      mapPublicIpOnLaunch: false,
      tags: { Environment: props.env, Name: `${props.env}-private-b` },
      vpcId: vpc.id,
    });

    // Public Route Table
    const publicRt = new RouteTable(this, `${props.env}-public-rt`, {
      tags: { Environment: props.env, Name: `${props.env}-public-rt` },
      vpcId: vpc.id,
    });

    new Route(this, `${props.env}-public-route`, {
      destinationCidrBlock: '0.0.0.0/0',
      gatewayId: igw.id,
      routeTableId: publicRt.id,
    });

    new RouteTableAssociation(this, `${props.env}-public-rt-assoc-a`, {
      routeTableId: publicRt.id,
      subnetId: publicSubnet1.id,
    });

    new RouteTableAssociation(this, `${props.env}-public-rt-assoc-b`, {
      routeTableId: publicRt.id,
      subnetId: publicSubnet2.id,
    });

    // Outputs
    this.vpcId = vpc.id;
    this.internetGatewayId = igw.id;
    this.publicRouteTableId = publicRt.id;
    this.publicSubnets = [publicSubnet1.id, publicSubnet2.id];
    this.privateSubnets = [privateSubnet1.id, privateSubnet2.id];

    new TerraformOutput(this, `${props.env}_vpc_id`, { value: this.vpcId });
    new TerraformOutput(this, `${props.env}_public_subnets`, {
      value: this.publicSubnets,
    });
    new TerraformOutput(this, `${props.env}_private_subnets`, {
      value: this.privateSubnets,
    });
    new TerraformOutput(this, `${props.env}_igw_id`, {
      value: this.internetGatewayId,
    });
    new TerraformOutput(this, `${props.env}_public_rt_id`, {
      value: this.publicRouteTableId,
    });
  }
}
