.PHONY: build build_docker deploy run test

.EXPORT_ALL_VARIABLES:
AWS_PROFILE = default
GOPROXY = direct

run:
	go run cmd/main.go

run_docker:
	docker run -it -e PARAM1=test1 go-fargate

build:
	GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/main ./cmd

build_docker:
	docker build -t go-fargate .

test:
	go test ./pkg/...

deploy: build_docker
	cd cdk;\
	cdk deploy --profile ${AWS_PROFILE}

destroy:
	cd cdk;\
	cdk destroy --profile ${AWS_PROFILE}