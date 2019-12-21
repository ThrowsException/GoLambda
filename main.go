package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type myEvent struct {
	Name string `json:"name"`
}

func handleRequest(ctx context.Context, name myEvent) (string, error) {
	return fmt.Sprintf("Goodbye %s!", name.Name), nil
}

func main() {
	lambda.Start(handleRequest)
}
