package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/harin-h/logs"
	"github.com/harin-h/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func main() {

	r := gin.Default()

	r.POST("/login", GetAllJourney)

	ginLambda = ginadapter.New(r)

	lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func GetAllJourney(ctx *gin.Context) {
	requestId := uuid.New().String()
	reqBody, _ := io.ReadAll(ctx.Request.Body)
	logs.Info(fmt.Sprint("[ID-", requestId, "] starting...\nRequest Body:", string(reqBody)))
	hashedPassword := os.Getenv("HASHED_PASSWORD")
	secretKey := os.Getenv("SECRET_KEY")
	if hashedPassword == "" || secretKey == "" {
		code := http.StatusInternalServerError
		msg := "unexpected error"
		logError := "hashed password or secret key is not found"
		ctx.String(code, msg)
		logs.Info(fmt.Sprint("[ID-", requestId, "] finished - Status Code:", code, " & Message:", msg, " & Log Error:", logError))
		return
	}
	req := map[string]string{}
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		code := http.StatusBadRequest
		msg := "incorrect request body"
		logs.Info(fmt.Sprint("[ID-", requestId, "] finished - Status Code:", code, " & Message:", msg))
		return
	}
	isValid := utils.ComparePassword(req["password"], hashedPassword)
	if !isValid {
		code := http.StatusUnauthorized
		msg := "wrong password"
		logs.Info(fmt.Sprint("[ID-", requestId, "] finished - Status Code:", code, " & Message:", msg))
		return
	}
	token, err := utils.GenerateToken(secretKey)
	if err != nil {
		code := http.StatusInternalServerError
		msg := "unexpected error"
		logError := "generating token fail"
		logs.Info(fmt.Sprint("[ID-", requestId, "] finished - Status Code:", code, " & Message:", msg, " & Log Error:", logError))
		return
	}
	code := http.StatusOK
	ctx.Status(code)
	ctx.SetCookie("token", token, 7200, "", "", true, true)
	logs.Info(fmt.Sprint("[ID-", requestId, "] finished - Status Code:", code))
}
