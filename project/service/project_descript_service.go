package service

import (
	"net/http"
	"regexp"

	"github.com/harin-h/portfolio-project-go-lambda/repository"
	errs "github.com/harin-h/rest-api-err"
)

func (s projectService) GetAllProjectDescript() ([]projectDescriptResponse, error) {
	projectDescripts, err := s.projectRepo.GetAllProjectDescript()
	if err != nil {
		return nil, errs.InternalServerError(err.Error())
	}
	var servRes []projectDescriptResponse
	for _, v := range projectDescripts {
		projectDescript := projectDescriptResponse{
			Id:             int(v.ID),
			ProjectName:    v.ProjectName,
			About:          v.About,
			WebsiteUrl:     v.WebsiteUrl,
			GithubUrl:      v.GithubUrl,
			DockerImageUrl: v.DockerImageUrl,
		}
		servRes = append(servRes, projectDescript)
	}
	return servRes, nil
}

func (s projectService) AddNewProjectDescript(servReq AddProjectDescriptServiceRequest) (err error) {
	defer func() {
		if errCommit := s.projectRepo.CloseTransaction(err); errCommit != nil && err != nil {
			err := err.(errs.AppError)
			err.LogError += errCommit.Error()
		}
	}()
	regexp, _ := regexp.Compile(`\S+`)
	if err := s.projectRepo.BeginTransaction(); err != nil {
		return errs.InternalServerError(err.Error())
	}
	if !regexp.MatchString(servReq.ProjectName) {
		return errs.AppError{Code: http.StatusBadRequest, Message: "incorrect body request", LogError: "inacceptable value (ex. wrong format)"}
	}
	repoReq := repository.AddProjectDescriptRepositoryRequest{
		ProjectName:    servReq.ProjectName,
		About:          servReq.About,
		WebsiteUrl:     servReq.WebsiteUrl,
		GithubUrl:      servReq.GithubUrl,
		DockerImageUrl: servReq.DockerImageUrl,
	}
	if err := s.projectRepo.AddNewProjectDescript(repoReq); err != nil {
		return errs.InternalServerError(err.Error())
	}
	return nil
}

func (s projectService) UpdateProjectDescript(servReq UpdateProjectDescriptServiceRequest) (err error) {
	defer func() {
		if errCommit := s.projectRepo.CloseTransaction(err); errCommit != nil && err != nil {
			err := err.(errs.AppError)
			err.LogError += errCommit.Error()
		}
	}()
	regexp, _ := regexp.Compile(`\S+`)
	if err := s.projectRepo.BeginTransaction(); err != nil {
		return errs.InternalServerError(err.Error())
	}
	if !regexp.MatchString(servReq.ProjectName) {
		return errs.AppError{Code: http.StatusBadRequest, Message: "incorrect body request", LogError: "inacceptable value (ex. wrong format)"}
	}
	repoReq := repository.UpdateProjectDescriptRepositoryRequest{
		Id:             servReq.Id,
		ProjectName:    servReq.ProjectName,
		About:          servReq.About,
		WebsiteUrl:     servReq.WebsiteUrl,
		GithubUrl:      servReq.GithubUrl,
		DockerImageUrl: servReq.DockerImageUrl,
	}
	if err := s.projectRepo.UpdateProjectDescript(repoReq); err != nil {
		return errs.InternalServerError(err.Error())
	}
	return nil
}

func (s projectService) DeleteProjectDescript(servReq DeleteServiceRequest) (err error) {
	defer func() {
		if errCommit := s.projectRepo.CloseTransaction(err); errCommit != nil && err != nil {
			err := err.(errs.AppError)
			err.LogError += errCommit.Error()
		}
	}()
	if err := s.projectRepo.BeginTransaction(); err != nil {
		return errs.InternalServerError(err.Error())
	}
	if err := s.projectRepo.DeleteProjectDescript(servReq.Id); err != nil {
		return errs.InternalServerError(err.Error())
	}
	if err := s.projectRepo.DeleteProjectPictureByProjectId(servReq.Id); err != nil {
		return errs.InternalServerError(err.Error())
	}
	if err := s.projectRepo.DeleteProjectTagByProjectId(servReq.Id); err != nil {
		return errs.InternalServerError(err.Error())
	}
	if err := s.projectRepo.DeleteProjectTopicByProjectId(servReq.Id); err != nil {
		return errs.InternalServerError(err.Error())
	}
	if err := s.projectRepo.DeleteGroupProjectByProjectId(servReq.Id); err != nil {
		return errs.InternalServerError(err.Error())
	}
	return nil
}
