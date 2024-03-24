package handler

import "github.com/gin-gonic/gin"

type JourneyHandler interface {
	GetAllJourney(*gin.Context)
	AddNewJourney(*gin.Context)
	UpdateJourney(*gin.Context)
	DeleteJourney(*gin.Context)
}
