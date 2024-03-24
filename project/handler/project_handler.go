package handler

import (
	"github.com/harin-h/portfolio-project-go-lambda/service"
)

type projectHandler struct {
	projectServ service.ProjectService
}

func NewProjectHandler(projectServ service.ProjectService) projectHandler {
	return projectHandler{projectServ: projectServ}
}
