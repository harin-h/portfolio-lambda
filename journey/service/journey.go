package service

type serviceResponse struct {
	Id             int    `json:"id" example:"1"`
	StartMonthYear string `json:"start_month_year" example:"07-2023"`
	EndMonthYear   string `json:"end_month_year" example:"08-2023"`
	Organization   string `json:"organization" example:"Chulalongkorn University"`
	Detail         string `json:"detail" example:"First Class Honor & Gold Medal"`
	PictureUrl     string `json:"picture_url"`
	SortValue      int    `json:"sort_value" example:"1"`
}

type AddServiceRequest struct {
	StartMonthYear string `json:"start_month_year" example:"07-2023"`
	EndMonthYear   string `json:"end_month_year" example:"08-2023"`
	Organization   string `json:"organization" example:"Chulalongkorn University"`
	Detail         string `json:"detail" example:"First Class Honor & Gold Medal"`
	PictureUrl     string `json:"picture_url"`
	SortValue      int    `json:"sort_value" example:"1"`
}

type UpdateServiceRequest struct {
	Id             int    `json:"id" example:"1"`
	StartMonthYear string `json:"start_month_year" example:"07-2023"`
	EndMonthYear   string `json:"end_month_year" example:"08-2023"`
	Organization   string `json:"organization" example:"Chulalongkorn University"`
	Detail         string `json:"detail" example:"First Class Honor & Gold Medal"`
	PictureUrl     string `json:"picture_url"`
	SortValue      int    `json:"sort_value" example:"1"`
}

type DeleteServiceRequest struct {
	Id int `json:"id" example:"1"`
}

type JourneyService interface {
	GetAllJourney() ([]serviceResponse, error)
	AddNewJourney([]AddServiceRequest) error
	UpdateJourney([]UpdateServiceRequest) error
	DeleteJourney([]DeleteServiceRequest) error
}
