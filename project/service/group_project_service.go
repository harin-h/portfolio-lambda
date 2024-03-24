package service

import (
	"github.com/harin-h/portfolio-project-go-lambda/repository"
	errs "github.com/harin-h/rest-api-err"
)

func (s projectService) GetAllGroupProject() ([]groupProjectResponse, error) {
	groupProjects, err := s.projectRepo.GetAllGroupProject()
	if err != nil {
		return nil, errs.InternalServerError(err.Error())
	}
	var servRes []groupProjectResponse
	for _, v := range groupProjects {
		groupProject := groupProjectResponse{
			Id:        int(v.ID),
			GroupId:   v.GroupId,
			ProjectId: v.ProjectId,
			SortValue: v.SortValue,
		}
		servRes = append(servRes, groupProject)
	}
	return servRes, nil
}

func (s projectService) AddGroupProject(servReq []AddGroupProjectServiceRequest) (err error) {
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
		repoReq := repository.AddGroupProjectRepositoryRequest{
			GroupId:   v.GroupId,
			ProjectId: v.ProjectId,
			SortValue: v.SortValue,
		}
		if err := s.projectRepo.AddGroupProject(repoReq); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}

func (s projectService) UpdateGroupProject(servReq []UpdateGroupProjectServiceRequest) (err error) {
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
		repoReq := repository.UpdateGroupProjectRepositoryRequest{
			Id:        v.Id,
			SortValue: v.SortValue,
		}
		if err := s.projectRepo.UpdateGroupProject(repoReq); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}

func (s projectService) DeleteGroupProject(servReq []DeleteServiceRequest) (err error) {
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
		if err := s.projectRepo.DeleteGroupProjectById(v.Id); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}
