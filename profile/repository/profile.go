package repository

import "gorm.io/gorm"

type profile struct {
	gorm.Model
	MetricName  string
	MetricValue string
	SortValue   int
}

type AddRepositoryRequest struct {
	MetricName  string
	MetricValue string
	SortValue   int
}

type UpdateRepositoryRequest struct {
	Id          int
	MetricValue string
	SortValue   int
}

type ProfileRepository interface {
	BeginTransaction() error
	CloseTransaction(error) error
	GetAllMetric() ([]profile, error)
	AddNewMetric(AddRepositoryRequest) error
	UpdateMetric(UpdateRepositoryRequest) error
	DeleteMetric(int) error
}
