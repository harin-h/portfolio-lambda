package service

import "github.com/harin-h/portfolio-project-go-lambda/repository"

type projectService struct {
	projectRepo repository.ProfileRepository
}

func NewProjectService(projectRepo repository.ProfileRepository) projectService {
	return projectService{projectRepo: projectRepo}
}
