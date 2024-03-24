package service

import (
	"net/http"
	"regexp"

	"github.com/harin-h/portfolio-project-go-lambda/repository"
	errs "github.com/harin-h/rest-api-err"
)

func (s projectService) GetProjectTopicByProjectId(Id int) ([]projectTopicResponse, error) {
	projectTopics, err := s.projectRepo.GetProjectTopicByProjectId(Id)
	if err != nil {
		return nil, errs.InternalServerError(err.Error())
	}
	var servRes []projectTopicResponse
	for _, v := range projectTopics {
		projectTopic := projectTopicResponse{
			Id:        int(v.ID),
			ProjectId: v.ProjectId,
			TopicName: v.TopicName,
			Detail:    v.Detail,
			SortValue: v.SortValue,
		}
		servRes = append(servRes, projectTopic)
	}
	return servRes, nil
}

func (s projectService) AddNewProjectTopic(servReq AddProjectTopicServiceRequest) (err error) {
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
	if !regexp.MatchString(servReq.TopicName) {
		return errs.AppError{Code: http.StatusBadRequest, Message: "incorrect body request", LogError: "inacceptable value (ex. wrong format)"}
	}
	repoReq := repository.AddProjectTopicRepositoryRequest{
		ProjectId: servReq.ProjectId,
		TopicName: servReq.TopicName,
		Detail:    servReq.Detail,
		SortValue: servReq.SortValue,
	}
	if err := s.projectRepo.AddNewProjectTopic(repoReq); err != nil {
		return errs.InternalServerError(err.Error())
	}
	return nil
}

func (s projectService) UpdateProjectTopic(servReq []UpdateProjectTopicServiceRequest) (err error) {
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
		if !regexp.MatchString(v.TopicName) {
			return errs.AppError{Code: http.StatusBadRequest, Message: "incorrect body request", LogError: "inacceptable value (ex. wrong format)"}
		}
		repoReq := repository.UpdateProjectTopicRepositoryRequest{
			Id:        v.Id,
			TopicName: v.TopicName,
			Detail:    v.Detail,
			SortValue: v.SortValue,
		}
		if err := s.projectRepo.UpdateProjectTopic(repoReq); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}

func (s projectService) DeleteProjectTopic(servReq DeleteServiceRequest) (err error) {
	defer func() {
		if errCommit := s.projectRepo.CloseTransaction(err); errCommit != nil && err != nil {
			err := err.(errs.AppError)
			err.LogError += errCommit.Error()
		}
	}()
	if err := s.projectRepo.BeginTransaction(); err != nil {
		return errs.InternalServerError(err.Error())
	}
	if err := s.projectRepo.DeleteProjectTopicById(servReq.Id); err != nil {
		return errs.InternalServerError(err.Error())
	}
	return nil
}
