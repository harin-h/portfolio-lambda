package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/harin-h/logs"
	"github.com/harin-h/portfolio-project-go-lambda/service"
	errs "github.com/harin-h/rest-api-err"

	"github.com/gin-gonic/gin"
)

func (h projectHandler) GetAllProjectDescript(ctx *gin.Context) {
	requestId := uuid.New().String()
	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] starting..."))
	res, err := h.projectServ.GetAllProjectDescript()
	if err != nil {
		handlerError(ctx, err)
		err := err.(errs.AppError)
		code := err.Code
		msg := err.Message
		logErr := err.LogError
		logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] finished - Status Code:", code, " Message:", msg, " Log Error:", logErr))
		return
	}
	ctx.JSON(http.StatusOK, res)
	code := http.StatusOK
	logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] finished - Status Code:", code, "\nResponse Body:", res))
}

func (h projectHandler) AddNewProjectDescript(ctx *gin.Context) {
	requestId := uuid.New().String()
	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	reqBody, _ := io.ReadAll(ctx.Request.Body)
	logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] starting...\nRequest Body:", reqBody))
	var req service.AddProjectDescriptServiceRequest
	if err := json.Unmarshal(reqBody, &req); err != nil {
		err := errs.AppError{Code: http.StatusBadRequest, Message: "incorrect request body", LogError: err.Error()}
		handlerError(ctx, err)
		code := err.Code
		msg := err.Message
		logErr := err.LogError
		logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] finished - Status Code:", code, " Message:", msg, " Log Error:", logErr))
		return
	}
	err := h.projectServ.AddNewProjectDescript(req)
	if err != nil {
		handlerError(ctx, err)
		err := err.(errs.AppError)
		code := err.Code
		msg := err.Message
		logErr := err.LogError
		logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] finished - Status Code:", code, " Message:", msg, " Log Error:", logErr))
		return
	}
	ctx.Status(http.StatusOK)
	code := http.StatusOK
	logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] finished - Status Code:", code))
}

func (h projectHandler) UpdateProjectDescript(ctx *gin.Context) {
	requestId := uuid.New().String()
	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	reqBody, _ := io.ReadAll(ctx.Request.Body)
	logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] starting...\nRequest Body:", reqBody))
	var req service.UpdateProjectDescriptServiceRequest
	if err := json.Unmarshal(reqBody, &req); err != nil {
		err := errs.AppError{Code: http.StatusBadRequest, Message: "incorrect request body", LogError: err.Error()}
		handlerError(ctx, err)
		code := err.Code
		msg := err.Message
		logErr := err.LogError
		logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] finished - Status Code:", code, " Message:", msg, " Log Error:", logErr))
		return
	}
	err := h.projectServ.UpdateProjectDescript(req)
	if err != nil {
		handlerError(ctx, err)
		err := err.(errs.AppError)
		code := err.Code
		msg := err.Message
		logErr := err.LogError
		logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] finished - Status Code:", code, " Message:", msg, " Log Error:", logErr))
		return
	}
	ctx.Status(http.StatusOK)
	code := http.StatusOK
	logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] finished - Status Code:", code))
}

func (h projectHandler) DeleteProjectDescript(ctx *gin.Context) {
	requestId := uuid.New().String()
	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	reqBody, _ := io.ReadAll(ctx.Request.Body)
	logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] starting...\nRequest Body:", reqBody))
	var req service.DeleteServiceRequest
	if err := json.Unmarshal(reqBody, &req); err != nil {
		err := errs.AppError{Code: http.StatusBadRequest, Message: "incorrect request body", LogError: err.Error()}
		handlerError(ctx, err)
		code := err.Code
		msg := err.Message
		logErr := err.LogError
		logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] finished - Status Code:", code, " Message:", msg, " Log Error:", logErr))
		return
	}
	err := h.projectServ.DeleteProjectDescript(req)
	if err != nil {
		handlerError(ctx, err)
		err := err.(errs.AppError)
		code := err.Code
		msg := err.Message
		logErr := err.LogError
		logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] finished - Status Code:", code, " Message:", msg, " Log Error:", logErr))
		return
	}
	ctx.Status(http.StatusOK)
	code := http.StatusOK
	logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] finished - Status Code:", code))
}
