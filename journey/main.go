package main

import (
	"context"
	"flag"

	"github.com/harin-h/portfolio-journey-go-lambda/handler"
	"github.com/harin-h/portfolio-journey-go-lambda/repository"
	"github.com/harin-h/portfolio-journey-go-lambda/service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var ginLambda *ginadapter.GinLambda

func main() {

	databaseUrl := flag.String("DATABASE", "postgres://postgresql_portfolio_wpec_user:uoTApYmU1hekviTf7UXU2ejezCv46yu3@dpg-cnfqd7n109ks738k8q0g-a.oregon-postgres.render.com/postgresql_portfolio_wpec", "Database Url")

	flag.Parse()

	db, err := gorm.Open(postgres.Open(*databaseUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	repo := repository.NewJourneyRepositoryDB(db)
	serv := service.NewJourneyService(repo)
	hdlr := handler.NewJourneyHandler(serv)

	r := gin.Default()

	r.GET("/journey", hdlr.GetAllJourney)
	r.POST("/journey", hdlr.AddNewJourney)
	r.PUT("/journey", hdlr.UpdateJourney)
	r.DELETE("/journey", hdlr.DeleteJourney)

	ginLambda = ginadapter.New(r)

	lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}
