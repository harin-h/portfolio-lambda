package handler

import "github.com/gin-gonic/gin"

type ProfileHandler interface {
	GetProfile(*gin.Context)
	AddMetricProfile(*gin.Context)
	UpdateMetricProfile(*gin.Context)
	DeleteMetricProfile(*gin.Context)
}
