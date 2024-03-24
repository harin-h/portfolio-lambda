package handler

import "github.com/gin-gonic/gin"

type ProjectHandler interface {
	GetAllProjectDescript(*gin.Context)
	AddNewProjectDescript(*gin.Context)
	UpdateProjectDescript(*gin.Context)
	DeleteProjectDescript(*gin.Context)
	GetAllProjectTag(*gin.Context)
	AddNewProjectTag(*gin.Context)
	UpdateProjectTag(*gin.Context)
	DeleteProjectTag(*gin.Context)
	GetAllProjectPicture(*gin.Context)
	AddNewProjectPicture(*gin.Context)
	UpdateProjectPicture(*gin.Context)
	DeleteProjectPicture(*gin.Context)
	GetProjectTopic(*gin.Context)
	AddNewProjectTopic(*gin.Context)
	UpdateProjectTopic(*gin.Context)
	DeleteProjectTopic(*gin.Context)
	GetAllGroup(*gin.Context)
	AddNewGroup(*gin.Context)
	UpdateGroup(*gin.Context)
	DeleteGroup(*gin.Context)
	GetAllGroupProject(*gin.Context)
	AddGroupProject(*gin.Context)
	UpdateGroupProject(*gin.Context)
	DeleteGroupProject(*gin.Context)
}
