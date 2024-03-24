package service

import (
	"net/http"
	"regexp"

	"github.com/harin-h/portfolio-profile-go-lambda/repository"
	errs "github.com/harin-h/rest-api-err"
)

type profileService struct {
	profileRepo repository.ProfileRepository
}

func NewProfileService(profileRepo repository.ProfileRepository) profileService {
	return profileService{profileRepo: profileRepo}
}

func (s profileService) GetProfile() ([]serviceResponse, error) {
	profiles, err := s.profileRepo.GetAllMetric()
	if err != nil {
		return nil, errs.InternalServerError(err.Error())
	}
	var servRes []serviceResponse
	for _, v := range profiles {
		servRes = append(servRes, serviceResponse{Id: int(v.ID), MetricName: v.MetricName, MetricValue: v.MetricValue, SortValue: v.SortValue})
	}
	return servRes, nil
}

func (s profileService) AddMetricProfile(servReq []AddServiceRequest) (err error) {
	defer func() {
		if errCommit := s.profileRepo.CloseTransaction(err); errCommit != nil && err != nil {
			err := err.(errs.AppError)
			err.LogError += errCommit.Error()
		}
	}()
	if err := s.profileRepo.BeginTransaction(); err != nil {
		return errs.InternalServerError(err.Error())
	}
	regexp, _ := regexp.Compile(`\S+`)
	for _, v := range servReq {
		if !regexp.MatchString(v.MetricName) || !regexp.MatchString(v.MetricValue) {
			return errs.AppError{Code: http.StatusBadRequest, Message: "incorrect body request", LogError: "inacceptable value (ex. wrong format)"}
		}
		addRepoReq := repository.AddRepositoryRequest{MetricName: v.MetricName, MetricValue: v.MetricValue, SortValue: v.SortValue}
		if err := s.profileRepo.AddNewMetric(addRepoReq); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}

func (s profileService) UpdateMetricProfile(servReq []UpdateServiceRequest) (err error) {
	defer func() {
		if errCommit := s.profileRepo.CloseTransaction(err); errCommit != nil && err != nil {
			err := err.(errs.AppError)
			err.LogError += errCommit.Error()
		}
	}()
	if err := s.profileRepo.BeginTransaction(); err != nil {
		return errs.InternalServerError(err.Error())
	}
	regexp, _ := regexp.Compile(`\S+`)
	for _, v := range servReq {
		if !regexp.MatchString(v.MetricValue) {
			return errs.AppError{Code: http.StatusBadRequest, Message: "incorrect body request", LogError: "inacceptable value (ex. wrong format)"}
		}
		updateRepoReq := repository.UpdateRepositoryRequest{Id: v.Id, MetricValue: v.MetricValue, SortValue: v.SortValue}
		if err := s.profileRepo.UpdateMetric(updateRepoReq); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}

func (s profileService) DeleteMetricProfile(servReq []DeleteServiceRequest) (err error) {
	defer func() {
		if errCommit := s.profileRepo.CloseTransaction(err); errCommit != nil && err != nil {
			err := err.(errs.AppError)
			err.LogError += errCommit.Error()
		}
	}()
	if err := s.profileRepo.BeginTransaction(); err != nil {
		return errs.InternalServerError(err.Error())
	}
	for _, v := range servReq {
		if err := s.profileRepo.DeleteMetric(v.Id); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}
