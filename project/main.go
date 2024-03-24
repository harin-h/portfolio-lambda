package main

import (
	"context"
	"flag"

	"github.com/harin-h/portfolio-project-go-lambda/handler"
	"github.com/harin-h/portfolio-project-go-lambda/repository"
	"github.com/harin-h/portfolio-project-go-lambda/service"

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

	repo := repository.NewProjectRepositoryDB(db)
	serv := service.NewProjectService(repo)
	hdlr := handler.NewProjectHandler(serv)

	r := gin.Default()

	r.GET("/project/descript", hdlr.GetAllProjectDescript)
	r.POST("/project/descript", hdlr.AddNewProjectDescript)
	r.PUT("/project/descript", hdlr.UpdateProjectDescript)
	r.DELETE("/project/descript", hdlr.DeleteProjectDescript)
	r.GET("/project/tag", hdlr.GetAllProjectTag)
	r.POST("/project/tag", hdlr.AddNewProjectTag)
	r.PUT("/project/tag", hdlr.UpdateProjectTag)
	r.DELETE("/project/tag", hdlr.DeleteProjectTag)
	r.GET("/project/picture", hdlr.GetAllProjectPicture)
	r.POST("/project/picture", hdlr.AddNewProjectPicture)
	r.PUT("/project/picture", hdlr.UpdateProjectPicture)
	r.DELETE("/project/picture", hdlr.DeleteProjectPicture)
	r.GET("/project/topic/:id", hdlr.GetProjectTopic)
	r.POST("/project/topic", hdlr.AddNewProjectTopic)
	r.PUT("/project/topic", hdlr.UpdateProjectTopic)
	r.DELETE("/project/topic", hdlr.DeleteProjectTopic)
	r.GET("/group", hdlr.GetAllGroup)
	r.POST("/group", hdlr.AddNewGroup)
	r.PUT("/group", hdlr.UpdateGroup)
	r.DELETE("/group", hdlr.DeleteGroup)
	r.GET("/group/project", hdlr.GetAllGroupProject)
	r.POST("/group/project", hdlr.AddGroupProject)
	r.PUT("/group/project", hdlr.UpdateGroupProject)
	r.DELETE("/group/project", hdlr.DeleteGroupProject)

	ginLambda = ginadapter.New(r)

	lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}
