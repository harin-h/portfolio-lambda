package service

type serviceResponse struct {
	Id          int    `json:"id" example:"1"`
	MetricName  string `json:"metric_name" example:"Name"`
	MetricValue string `json:"metric_value" example:"Harin Harisombut"`
	SortValue   int    `json:"sort_value" example:"1"`
}

type AddServiceRequest struct {
	MetricName  string `json:"metric_name" example:"Name"`
	MetricValue string `json:"metric_value" example:"Harin Harisombut"`
	SortValue   int    `json:"sort_value" example:"1"`
}

type UpdateServiceRequest struct {
	Id          int    `json:"id" example:"1"`
	MetricValue string `json:"metric_value" example:"Harin Harisombut"`
	SortValue   int    `json:"sort_value" example:"1"`
}

type DeleteServiceRequest struct {
	Id int `json:"id" example:"1"`
}

type ProfileService interface {
	GetProfile() ([]serviceResponse, error)
	AddMetricProfile([]AddServiceRequest) error
	UpdateMetricProfile([]UpdateServiceRequest) error
	DeleteMetricProfile([]DeleteServiceRequest) error
}
