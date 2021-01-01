package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"fargate-boilerplate/pkg/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	var sess *session.Session

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	awsProfile := os.Getenv("AWS_PROFILE")
	log.Printf("AWS_PROFILE: %s", awsProfile)

	if awsProfile != "" {
		log.Println("Use AWS profile")
		sess = session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
			Profile:           awsProfile,
		}))
	} else {
		log.Println("Use container role")
		sess = session.Must(session.NewSession())
	}

	// Create S3 service client
	svc := s3.New(sess)

	result, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatalf("Unable to list buckets, %v", err)
	}

	log.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
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
