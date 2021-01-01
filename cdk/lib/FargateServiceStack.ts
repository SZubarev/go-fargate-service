import * as cdk from '@aws-cdk/core';
import * as path from 'path';
import * as ec2 from '@aws-cdk/aws-ec2';
import * as ecs from '@aws-cdk/aws-ecs';
import * as iam from '@aws-cdk/aws-iam'
import { LogGroup } from '@aws-cdk/aws-logs'
import { DockerImageAsset } from '@aws-cdk/aws-ecr-assets';
import { ContainerImage, FargatePlatformVersion } from '@aws-cdk/aws-ecs';


export class FargateServiceStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const asset = new DockerImageAsset(this, 'container-lambda', {
      directory: path.join(__dirname, "..", ".."),
    });

    const vpc = new ec2.Vpc(this, "EcsVpc", {
      maxAzs: 3 // Default is all AZs in region
    });

    const cluster = new ecs.Cluster(this, "TestCluster", {
      vpc: vpc,
      clusterName: "go-service-cluster",
      containerInsights: false
    });

    const logGroup = new LogGroup(this, "FargateLogGroup", {
      logGroupName: "/ecs/go-service"
    })

    const taskDef = new ecs.FargateTaskDefinition(this, "MyTask", {
      cpu: 512,
      memoryLimitMiB: 1024,
    })

    const container = new ecs.ContainerDefinition(this, "MyContainer", {
      image: ContainerImage.fromDockerImageAsset(asset),
      taskDefinition: taskDef,
      environment: {
        PARAM1: "test1"
      },
      logging: new ecs.AwsLogDriver({
        logGroup: logGroup,
        streamPrefix: `go-service`,
      })
    }
    )

    const myService = new ecs.FargateService(this, "MyService", {
      taskDefinition: taskDef,
      cluster: cluster,
      platformVersion: FargatePlatformVersion.VERSION1_4,
      serviceName: "go-fargate-service",
      desiredCount: 1
    })

    taskDef.addToTaskRolePolicy(new iam.PolicyStatement({
      actions:["s3:ListAllMyBuckets"],
      resources:["arn:aws:s3:::*"],
      effect: iam.Effect.ALLOW
    }))

  }
}
