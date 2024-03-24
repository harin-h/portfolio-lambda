package service

import (
	"net/http"
	"regexp"

	"github.com/harin-h/portfolio-journey-go-lambda/repository"

	errs "github.com/harin-h/rest-api-err"
)

type journeyService struct {
	journeyRepo repository.JourneyRepository
}

func NewJourneyService(journeyRepo repository.JourneyRepository) journeyService {
	return journeyService{journeyRepo: journeyRepo}
}

func (s journeyService) GetAllJourney() ([]serviceResponse, error) {
	journeys, err := s.journeyRepo.GetAllJourney()
	if err != nil {
		return nil, errs.InternalServerError(err.Error())
	}
	var servRes []serviceResponse
	for _, v := range journeys {
		journey := serviceResponse{
			Id:             int(v.ID),
			StartMonthYear: v.StartMonthYear,
			EndMonthYear:   v.EndMonthYear,
			Organization:   v.Organization,
			Detail:         v.Detail,
			PictureUrl:     v.PictureUrl,
			SortValue:      v.SortValue,
		}
		servRes = append(servRes, journey)
	}
	return servRes, nil
}

func (s journeyService) AddNewJourney(servReq []AddServiceRequest) (err error) {
	defer func() {
		if errCommit := s.journeyRepo.CloseTransaction(err); errCommit != nil && err != nil {
			err := err.(errs.AppError)
			err.LogError += errCommit.Error()
		}
	}()
	if err := s.journeyRepo.BeginTransaction(); err != nil {
		return errs.InternalServerError(err.Error())
	}
	regexpMonthYear, _ := regexp.Compile(`(0[123456789]|1[012])-(19\d{2}|20\d{2})`)
	regexpString, _ := regexp.Compile(`\S+`)
	for _, v := range servReq {
		if !regexpMonthYear.MatchString(v.StartMonthYear) || !regexpMonthYear.MatchString(v.EndMonthYear) || !regexpString.MatchString(v.Organization) {
			err = errs.AppError{Code: http.StatusBadRequest, Message: "incorrect body request", LogError: "inacceptable value (wrong format)"}
			return err
		}
		addRepoReq := repository.AddRepositoryRequest{
			StartMonthYear: v.StartMonthYear,
			EndMonthYear:   v.EndMonthYear,
			Organization:   v.Organization,
			Detail:         v.Detail,
			PictureUrl:     v.PictureUrl,
			SortValue:      v.SortValue,
		}
		if err := s.journeyRepo.AddNewJourney(addRepoReq); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}

func (s journeyService) UpdateJourney(servReq []UpdateServiceRequest) (err error) {
	defer func() {
		if errCommit := s.journeyRepo.CloseTransaction(err); errCommit != nil && err != nil {
			err := err.(errs.AppError)
			err.LogError += errCommit.Error()
		}
	}()
	if err := s.journeyRepo.BeginTransaction(); err != nil {
		return errs.InternalServerError(err.Error())
	}
	regexpMonthYear, _ := regexp.Compile(`(0[123456789]|1[012])-(19\d{2}|20\d{2})`)
	regexpString, _ := regexp.Compile(`\S+`)
	for _, v := range servReq {
		if !regexpMonthYear.MatchString(v.StartMonthYear) || !regexpMonthYear.MatchString(v.EndMonthYear) || !regexpString.MatchString(v.Organization) {
			return errs.AppError{Code: http.StatusBadRequest, Message: "incorrect body request", LogError: "inacceptable value (wrong format)"}
		}
		updateRepoReq := repository.UpdateRepositoryRequest{
			Id:             v.Id,
			StartMonthYear: v.StartMonthYear,
			EndMonthYear:   v.EndMonthYear,
			Organization:   v.Organization,
			Detail:         v.Detail,
			PictureUrl:     v.PictureUrl,
			SortValue:      v.SortValue,
		}
		if err := s.journeyRepo.UpdateJourney(updateRepoReq); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}

func (s journeyService) DeleteJourney(servReq []DeleteServiceRequest) (err error) {
	defer func() {
		if errCommit := s.journeyRepo.CloseTransaction(err); errCommit != nil && err != nil {
			err := err.(errs.AppError)
			err.LogError += errCommit.Error()
		}
	}()
	if err := s.journeyRepo.BeginTransaction(); err != nil {
		return errs.InternalServerError(err.Error())
	}
	for _, v := range servReq {
		if err := s.journeyRepo.DeleteJourney(v.Id); err != nil {
			return errs.InternalServerError(err.Error())
		}
	}
	return nil
}
