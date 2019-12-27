package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler() (events.APIGatewayProxyResponse, error) {
	body := map[string]interface{}{
		"Message": "Go Serverless v1.0! Your function executed successfully!",
	}
	s, _ := json.Marshal(body)
	return events.APIGatewayProxyResponse{Body: string(s), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
