package service

import (
	"net/http"
	"regexp"

	"github.com/harin-h/portfolio-project-go-lambda/repository"
	errs "github.com/harin-h/rest-api-err"
)

func (s projectService) GetAllProjectPicture() ([]projectPictureResponse, error) {
	projectPictures, err := s.projectRepo.GetAllProjectPicture()
	if err != nil {
		return nil, errs.InternalServerError(err.Error())
	}
	var servRes []projectPictureResponse
	for _, v := range projectPictures {
		projectPicture := projectPictureResponse{
			Id:         int(v.ID),
			ProjectId:  v.ProjectId,
			PictureUrl: v.PictureUrl,
			SortValue:  v.SortValue,
		}
		servRes = append(servRes, projectPicture)
	}
	return servRes, nil
}

func (s projectService) AddNewProjectPicture(servReq []AddProjectPictureServiceRequest) (err error) {
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
		if !regexp.MatchString(v.PictureUrl) {
			return errs.AppError{Code: http.StatusBadRequest, Message: "incorrect body request", LogError: "inacceptable value (ex. wrong format)"}
		}
		repoReq := repository.AddProjectPictureRepositoryRequest{
			ProjectId:  v.ProjectId,
			PictureUrl: v.PictureUrl,
			SortValue:  v.SortValue,
		}
		if err := s.projectRepo.AddNewProjectPicture(repoReq); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}

func (s projectService) UpdateProjectPicture(servReq []UpdateProjectPictureServiceRequest) (err error) {
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
		repoReq := repository.UpdateProjectPictureRepositoryRequest{
			Id:        v.Id,
			SortValue: v.SortValue,
		}
		if err := s.projectRepo.UpdateProjectPicture(repoReq); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}

func (s projectService) DeleteProjectPicture(servReq []DeleteServiceRequest) (err error) {
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
		if err := s.projectRepo.DeleteProjectPictureById(v.Id); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}
