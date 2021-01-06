package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"fargate-boilerplate/pkg/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var cfg aws.Config
var ctx = context.TODO()

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	var err error

	awsProfile := os.Getenv("AWS_PROFILE")
	log.Printf("AWS_PROFILE: %s", awsProfile)

	if awsProfile != "" {
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
