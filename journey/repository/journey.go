package repository

import "gorm.io/gorm"

type journey struct {
	gorm.Model
	StartMonthYear string
	EndMonthYear   string
	Organization   string
	Detail         string
	PictureUrl     string
	SortValue      int
}

type AddRepositoryRequest struct {
	StartMonthYear string
	EndMonthYear   string
	Organization   string
	Detail         string
	PictureUrl     string
	SortValue      int
}

type UpdateRepositoryRequest struct {
	Id             int
	StartMonthYear string
	EndMonthYear   string
	Organization   string
	Detail         string
	PictureUrl     string
	SortValue      int
}

type JourneyRepository interface {
	BeginTransaction() error
	CloseTransaction(error) error
	GetAllJourney() ([]journey, error)
	AddNewJourney(AddRepositoryRequest) error
	UpdateJourney(UpdateRepositoryRequest) error
	DeleteJourney(int) error
}
