package service

import (
	"net/http"
	"regexp"

	"github.com/harin-h/portfolio-project-go-lambda/repository"
	errs "github.com/harin-h/rest-api-err"
)

func (s projectService) GetAllProjectTag() ([]projectTagResponse, error) {
	projectTags, err := s.projectRepo.GetAllProjectTag()
	if err != nil {
		return nil, errs.InternalServerError(err.Error())
	}
	var servRes []projectTagResponse
	for _, v := range projectTags {
		projectTag := projectTagResponse{
			Id:        int(v.ID),
			ProjectId: v.ProjectId,
			Main:      v.Main,
			Sub:       v.Sub,
			SortValue: v.SortValue,
		}
		servRes = append(servRes, projectTag)
	}
	return servRes, nil
}

func (s projectService) AddNewProjectTag(servReq []AddProjectTagServiceRequest) (err error) {
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
	for _, v := range servReq {
		if !regexp.MatchString(v.Main) {
			return errs.AppError{Code: http.StatusBadRequest, Message: "incorrect body request", LogError: "inacceptable value (ex. wrong format)"}
		}
		repoReq := repository.AddProjectTagRepositoryRequest{
			ProjectId: v.ProjectId,
			Main:      v.Main,
			Sub:       v.Sub,
			SortValue: v.SortValue,
		}
		if err := s.projectRepo.AddNewProjectTag(repoReq); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}

func (s projectService) UpdateProjectTag(servReq []UpdateProjectTagServiceRequest) (err error) {
	defer func() {
		if errCommit := s.projectRepo.CloseTransaction(err); errCommit != nil && err != nil {
			err := err.(errs.AppError)
			err.LogError += errCommit.Error()
		}
	}()
	if err := s.projectRepo.BeginTransaction(); err != nil {
		return errs.InternalServerError(err.Error())
	}
	for _, v := range servReq {
		repoReq := repository.UpdateProjectTagRepositoryRequest{
			Id:        v.Id,
			SortValue: v.SortValue,
		}
		if err := s.projectRepo.UpdateProjectTag(repoReq); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}

func (s projectService) DeleteProjectTag(servReq []DeleteServiceRequest) (err error) {
	defer func() {
		if errCommit := s.projectRepo.CloseTransaction(err); errCommit != nil && err != nil {
			err := err.(errs.AppError)
			err.LogError += errCommit.Error()
		}
	}()
	if err := s.projectRepo.BeginTransaction(); err != nil {
		return errs.InternalServerError(err.Error())
	}
	for _, v := range servReq {
		if err := s.projectRepo.DeleteProjectTagById(v.Id); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}
