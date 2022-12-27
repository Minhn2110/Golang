package main

import (
	"context"
	"log"

	pb "github.com/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 7,
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %d\n", res.Result)
}
