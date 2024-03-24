package service

import (
	"net/http"
	"regexp"

	"github.com/harin-h/portfolio-project-go-lambda/repository"
	errs "github.com/harin-h/rest-api-err"
)

func (s projectService) GetAllGroup() ([]groupResponse, error) {
	groups, err := s.projectRepo.GetAllGroup()
	if err != nil {
		return nil, errs.InternalServerError(err.Error())
	}
	var servRes []groupResponse
	for _, v := range groups {
		group := groupResponse{
			Id:        int(v.ID),
			GroupName: v.GroupName,
			Detail:    v.Detail,
			SortValue: v.SortValue,
		}
		servRes = append(servRes, group)
	}
	return servRes, nil
}

func (s projectService) AddNewGroup(servReq AddGroupServiceRequest) (err error) {
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
	if !regexp.MatchString(servReq.GroupName) {
		return errs.AppError{Code: http.StatusBadRequest, Message: "incorrect body request", LogError: "inacceptable value (ex. wrong format)"}
	}
	repoReq := repository.AddGroupRepositoryRequest{
		GroupName: servReq.GroupName,
		Detail:    servReq.Detail,
		SortValue: servReq.SortValue,
	}
	if err := s.projectRepo.AddNewGroup(repoReq); err != nil {
		return errs.InternalServerError(err.Error())
	}
	return nil
}

func (s projectService) UpdateGroup(servReq []UpdateGroupServiceRequest) (err error) {
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
		if !regexp.MatchString(v.GroupName) {
			return errs.AppError{Code: http.StatusBadRequest, Message: "incorrect body request", LogError: "inacceptable value (ex. wrong format)"}
		}
		repoReq := repository.UpdateGroupRepositoryRequest{
			Id:        v.Id,
			GroupName: v.GroupName,
			Detail:    v.Detail,
			SortValue: v.SortValue,
		}
		if err := s.projectRepo.UpdateGroup(repoReq); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}

func (s projectService) DeleteGroup(servReq DeleteServiceRequest) (err error) {
	defer func() {
		if errCommit := s.projectRepo.CloseTransaction(err); errCommit != nil && err != nil {
			err := err.(errs.AppError)
			err.LogError += errCommit.Error()
		}
	}()
	if err := s.projectRepo.BeginTransaction(); err != nil {
		return errs.InternalServerError(err.Error())
	}
	if err := s.projectRepo.DeleteGroup(servReq.Id); err != nil {
		return errs.InternalServerError(err.Error())
	}
	if err := s.projectRepo.DeleteGroupProjectByGroupId(servReq.Id); err != nil {
		return errs.InternalServerError(err.Error())
	}
	return nil
}
