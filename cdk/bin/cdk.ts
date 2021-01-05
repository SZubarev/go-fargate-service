#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from '@aws-cdk/core';

import { FargateServiceStack } from '../lib/FargateServiceStack';


const app = new cdk.App();
new FargateServiceStack(app, 'FargateServiceStack');

