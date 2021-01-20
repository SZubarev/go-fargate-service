package main

import (
	"context"
	"os"
	"time"

	_ "fargate-boilerplate/pkg/log" // Configure logrus

	log "github.com/sirupsen/logrus"

	"fargate-boilerplate/pkg/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var cfg aws.Config
var ctx = context.TODO()

func init() {

	var err error

	awsProfile, ok := os.LookupEnv("AWS_PROFILE")
	log.Infof("AWS_PROFILE: %s", awsProfile)

	if ok {
		log.Info("Use AWS profile")
		cfg, err = config.LoadDefaultConfig(ctx,
			config.WithSharedConfigProfile(awsProfile),
		)
		if err != nil {
			log.Fatalf("Error loading profile %v", err)
		}

	} else {
		log.Info("Use container role")
		cfg, err = config.LoadDefaultConfig(ctx)
		if err != nil {
			log.Fatalf("Error loading profile %v", err)
		}
	}
}

func main() {

	// Create S3 service client
	s3Client := s3.NewFromConfig(cfg)

	result, err := s3Client.ListBuckets(ctx, nil)
	if err != nil {
		log.Fatalf("Unable to list buckets, %v", err)
	}

	log.Debug("Buckets:")

	for _, b := range result.Buckets {
		log.Debugf("* %s", aws.ToString(b.Name))
	}

	param1 := os.Getenv("PARAM1")

	log.Info("Service started")

	log.Infof("Param1: %s", param1)

	s := utils.GetHello()

	log.Infof("Test from package: %s", s)

	log.Error("Test error message")

	for true {

		log.Warnf("Service is working")
		time.Sleep(5 * time.Second)

	}

}
