package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

var (
	logGroup string = os.Getenv("LOG_GROUP")
)

func fetchData() {
	// fetch data from Cloud Watch for a particular log group
	fmt.Println(logGroup)
}

func analyzeTraffic() {

}

func reportAddresses() {

}

func handler() (string, error) {

	return "Hello", nil
}

func main() {
	fmt.Println("Registering handler ...")
	lambda.Start(handler)
}
