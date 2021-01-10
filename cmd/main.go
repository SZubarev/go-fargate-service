package main

import (
	"context"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"fargate-boilerplate/pkg/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var cfg aws.Config
var ctx = context.TODO()

func init() {

	configLogrus()

	var err error

	awsProfile, ok := os.LookupEnv("AWS_PROFILE")
	log.Printf("AWS_PROFILE: %s", awsProfile)

	if ok {
		log.Println("Use AWS profile")
		cfg, err = config.LoadDefaultConfig(ctx,
			config.WithSharedConfigProfile(awsProfile),
		)
		if err != nil {
			log.Fatalf("Error loading profile %v", err)
		}

	} else {
		log.Println("Use container role")
		cfg, err = config.LoadDefaultConfig(ctx)
		if err != nil {
			log.Fatalf("Error loading profile %v", err)
		}
	}
}

func configLogrus() {

	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05.000"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)

	logLevelStr, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		logLevelStr = "debug"
	}

	logLevel, err := log.ParseLevel(logLevelStr)
	if err != nil {
		logLevel = log.DebugLevel
	}

	log.SetLevel(logLevel)
}

func main() {

	// Create S3 service client
	s3Client := s3.NewFromConfig(cfg)

	result, err := s3Client.ListBuckets(ctx, nil)
	if err != nil {
		log.Fatalf("Unable to list buckets, %v", err)
	}

	log.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s\n", aws.ToString(b.Name))
	}

	param1 := os.Getenv("PARAM1")

	log.Println("Service started")

	log.Printf("Param1: %s", param1)

	s := utils.GetHello()

	log.Printf("Hello world %s", s)

	for true {

		log.Printf("Service is working")
		time.Sleep(5 * time.Second)

	}

}
