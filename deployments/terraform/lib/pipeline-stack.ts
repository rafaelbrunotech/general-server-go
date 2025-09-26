import { TerraformOutput, TerraformStack } from 'cdktf';
import { Construct } from 'constructs';

import { CodebuildProject } from '../.gen/providers/aws/codebuild-project';
import {
  Codepipeline,
  CodepipelineArtifactStore,
} from '../.gen/providers/aws/codepipeline';
import { S3Bucket } from '../.gen/providers/aws/s3-bucket';
import { EnvironmentOptions } from './shared/models/environment.interface';

export interface PipelineStackProps extends EnvironmentOptions {}

export class PipelineStack extends TerraformStack {
  constructor(scope: Construct, id: string, props: PipelineStackProps) {
    super(scope, id);

    // Artifact Bucket for pipeline
    const artifactBucket = new S3Bucket(this, `${props.env}-pipeline-artifacts`, {
      bucket: `${props.env}-pipeline-artifacts`,
      tags: { Environment: props.env },
    });

    // CodeBuild Project (for build steps)
    const buildProject = new CodebuildProject(this, `${props.env}-build-project`, {
        artifacts: {
          type: 'CODEPIPELINE', // required for integration with CodePipeline
        },
        environment: {
          computeType: 'BUILD_GENERAL1_SMALL',
          image: 'aws/codebuild/standard:7.0',
          type: 'LINUX_CONTAINER',
        },
        name: `${props.env}-build-project`,
        serviceRole: 'REPLACE_WITH_CODEBUILD_ROLE_ARN',
        source: {
          buildspec: 'buildspec.yml', // or inline commands
          gitCloneDepth: 1,
          location: 'https://github.com/YOUR_GITHUB_USERNAME/YOUR_REPO_NAME.git',
          type: 'GITHUB',
        },
    });

    // CodePipeline
    const pipeline = new Codepipeline(this, `${props.env}-pipeline`, {
      artifactStore: [
        {
          location: artifactBucket.bucket,
          type: 'S3',
        },
      ] as CodepipelineArtifactStore[],
      name: `${props.env}-pipeline`,
      roleArn: 'REPLACE_WITH_CODEPIPELINE_ROLE_ARN', // Use IAM role from platform
      stage: [
        {
          action: [
            {
              category: 'Source',
              configuration: {
                Branch: 'main',
                OAuthToken: 'YOUR_GITHUB_OAUTH_TOKEN',
                Owner: 'YOUR_GITHUB_USERNAME',
                Repo: 'YOUR_REPO_NAME',
              },
              name: 'Source',
              outputArtifacts: ['SourceOutput'],
              owner: 'ThirdParty',
              provider: 'GitHub',
              version: '1',
            },
          ],
          name: 'Source',
        },
        {
          action: [
            {
              category: 'Build',
              configuration: {
                ProjectName: buildProject.name,
              },
              inputArtifacts: ['SourceOutput'],
              name: 'Build',
              outputArtifacts: ['BuildOutput'],
              owner: 'AWS',
              provider: 'CodeBuild',
              version: '1',
            },
          ],
          name: 'Build',
        },
      ],
      tags: { Environment: props.env },
    });

    new TerraformOutput(this, `${props.env}_pipeline_name`, {
      value: pipeline.name,
    });
    new TerraformOutput(this, `${props.env}_artifact_bucket`, {
      value: artifactBucket.bucket,
    });
  }
}
