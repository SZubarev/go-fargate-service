package main

import (
	"log"
	"os"
	"time"

	"fargate-boilerplate/pkg/utils"
)

func main() {

	param1 := os.Getenv("PARAM1")

	log.Println("Service started")

	log.Printf("Param1: %s", param1)

	s := utils.GetHello()

	log.Printf("Hello world %s", s)

	for true {

		log.Printf("Service is working")
		time.Sleep(2 * time.Second)

	}

}
