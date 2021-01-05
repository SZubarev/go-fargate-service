# Golang AWS Fargate Service Boilerplate
Boilerplate code of AWS Fargate service with Go deployed with AWS CDK.

Batteries included:
* Service logic stub in Go
* Using Go modules
* Using other AWS services in container (S3)
* Grant permissions to Fargate Task to call AWS services
* Multi-stage build
* Testing in build container
* Running locally as an app
* Running locally in Docker container
* Deployment as a AWS Fargate service
* Deployment to AWS using AWS CDK

Installation:
* Clone the repo
* Have Docker installed
* Install AWS CDK: https://docs.aws.amazon.com/cdk/latest/guide/getting_started.html
* Run `cd cdk; npm install`
* Configure AWS profile using `aws configure`
* Check the profile is setup using `aws s3 ls --profile default` (if you want to use another AWS profile - edit `AWS_PROFILE` variable in `Makefile`)
* Deploy the stack using `make deploy`
* Enjoy!
* Undeploy the stack using `make destroy`

Commands (see Makefile for details):
* `make run` - run locally as an app
* `make run_docker` - run locally in Docker container
* `make build` - build Linux executable
* `make build_docker` - build Docker container
* `make test` - run tests locally
* `make deploy` - deploys the stack to AWS (using default AWS profile)
* `make destroy` - destroys the stack