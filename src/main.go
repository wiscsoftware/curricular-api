package main

import (
	"curricular-api/api"
	"flag"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"net/http"
	"os"
)

func main() {

	// local deployment will set up OAuth2/JWT setup for demonstration,
	// AWS/Lambda deployment will use AWS APIGateway for API authorization.

	local := flag.Bool("local", false, "local deployment")
	port := flag.Int("port", 8080, "port")

	flag.Parse()

	numOfRecords := 20
	c := api.NewConfiguration(numOfRecords)
	handlers := api.Handlers(c)

	if *local {
		fmt.Println("Curricular API server is listing on port:", *port)

		err := http.ListenAndServe(fmt.Sprintf(":%d", *port), handlers)
		if err != nil {
			fmt.Println("error starting the sever", err)
			os.Exit(1)
		}
	} else {
		// AWS API Gateway
		// https://github.com/awslabs/aws-lambda-go-api-proxy/
		lambda.Start(httpadapter.New(handlers).ProxyWithContext)
	}
}
