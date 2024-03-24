package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/harin-h/logs"
	"github.com/harin-h/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayCustomAuthorizerRequestTypeRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	requestId := uuid.New().String()
	logs.Info(fmt.Sprint("[ID-", requestId, " | Authorization] starting..."))
	res := events.APIGatewayCustomAuthorizerResponse{PrincipalID: "Editor"}
	policy := events.APIGatewayCustomAuthorizerPolicy{Version: "2012-10-17"}
	statement := events.IAMPolicyStatement{Action: []string{"execute-api:Invoke"}, Resource: []string{
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/DELETE/group",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/POST/group",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/PUT/group",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/DELETE/group/project",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/POST/group/project",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/PUT/group/project",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/DELETE/project/descript",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/POST/project/descript",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/PUT/project/descript",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/DELETE/project/tag",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/POST/project/tag",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/PUT/project/tag",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/DELETE/project/topic",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/POST/project/topic",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/PUT/project/topic",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/DELETE/project/picture",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/POST/project/picture",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/PUT/project/picture",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/DELETE/journey",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/POST/journey",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/PUT/journey",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/DELETE/profile",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/POST/profile",
		"arn:aws:execute-api:ap-southeast-2:767397842364:f55ffuid2i/*/PUT/profile",
	}}
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		statement.Effect = "Deny"
		policy.Statement = []events.IAMPolicyStatement{statement}
		res.PolicyDocument = policy
		logs.Info(fmt.Sprint("[ID-", requestId, " | Authorization] finished -> Secret Key is not found (Unauthorized)"))
		return res, fmt.Errorf("secret key is not found")
	}
	tokenReq, prs := request.Headers["Cookie"]
	tokenString := strings.TrimPrefix(tokenReq, "token=")
	if !prs {
		statement.Effect = "Deny"
		policy.Statement = []events.IAMPolicyStatement{statement}
		res.PolicyDocument = policy
		logs.Info(fmt.Sprint("[ID-", requestId, " | Authorization] finished -> Unauthorized"))
		return res, nil
	}
	isValid, err := utils.ValidateToken(tokenString, secretKey)
	if err != nil || !isValid {
		statement.Effect = "Deny"
		policy.Statement = []events.IAMPolicyStatement{statement}
		res.PolicyDocument = policy
		logs.Info(fmt.Sprint("[ID-", requestId, " | Authorization] finished -> Unauthorized"))
		return res, nil
	}
	statement.Effect = "Allow"
	policy.Statement = []events.IAMPolicyStatement{statement}
	res.PolicyDocument = policy
	logs.Info(fmt.Sprint("[ID-", requestId, " | Authorization] finished -> Authorized"))
	return res, nil
}
