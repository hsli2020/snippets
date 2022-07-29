// https://www.abilityrush.com/working-with-aws-lambda-golang-and-cache/
package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   string `json:"age"`
	City  string `json:"city"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var userCache = map[string]Request{}

func HandleRequest(ctx context.Context, req Request) (Response, error) {

	if req.Email == "" || req.Name == "" || req.Age == "" || req.City == "" {
		return Response{Status: "error", Message: "all fields are required"}, nil
	}

	if isCached(req.Email) {
		return Response{
			Status:  "success",
			Message: "User already present in cache of this container",
		}, nil
	}

	//cache the user
	userCache[req.Email] = req

	return Response{
		Status:  "success",
		Message: "User cached",
	}, nil
}

func isCached(email string) bool {
	if _, ok := userCache[email]; ok {
		return true
	}
	return false
}

func main() {
	lambda.Start(HandleRequest)
}
