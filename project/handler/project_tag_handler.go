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

func (h projectHandler) GetAllProjectTag(ctx *gin.Context) {
	requestId := uuid.New().String()
	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] starting..."))
	res, err := h.projectServ.GetAllProjectTag()
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

func (h projectHandler) AddNewProjectTag(ctx *gin.Context) {
	requestId := uuid.New().String()
	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	reqBody, _ := io.ReadAll(ctx.Request.Body)
	logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] starting...\nRequest Body:", reqBody))
	var req []service.AddProjectTagServiceRequest
	if err := json.Unmarshal(reqBody, &req); err != nil {
		err := errs.AppError{Code: http.StatusBadRequest, Message: "incorrect request body", LogError: err.Error()}
		handlerError(ctx, err)
		code := err.Code
		msg := err.Message
		logErr := err.LogError
		logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] finished - Status Code:", code, " Message:", msg, " Log Error:", logErr))
		return
	}
	err := h.projectServ.AddNewProjectTag(req)
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

func (h projectHandler) UpdateProjectTag(ctx *gin.Context) {
	requestId := uuid.New().String()
	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	reqBody, _ := io.ReadAll(ctx.Request.Body)
	logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] starting...\nRequest Body:", reqBody))
	var req []service.UpdateProjectTagServiceRequest
	if err := json.Unmarshal(reqBody, &req); err != nil {
		err := errs.AppError{Code: http.StatusBadRequest, Message: "incorrect request body", LogError: err.Error()}
		handlerError(ctx, err)
		code := err.Code
		msg := err.Message
		logErr := err.LogError
		logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] finished - Status Code:", code, " Message:", msg, " Log Error:", logErr))
		return
	}
	err := h.projectServ.UpdateProjectTag(req)
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

func (h projectHandler) DeleteProjectTag(ctx *gin.Context) {
	requestId := uuid.New().String()
	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	reqBody, _ := io.ReadAll(ctx.Request.Body)
	logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] starting...\nRequest Body:", reqBody))
	var req []service.DeleteServiceRequest
	if err := json.Unmarshal(reqBody, &req); err != nil {
		err := errs.AppError{Code: http.StatusBadRequest, Message: "incorrect request body", LogError: err.Error()}
		handlerError(ctx, err)
		code := err.Code
		msg := err.Message
		logErr := err.LogError
		logs.Info(fmt.Sprint("[ID-", requestId, "|", method, "|", path, "] finished - Status Code:", code, " Message:", msg, " Log Error:", logErr))
		return
	}
	err := h.projectServ.DeleteProjectTag(req)
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
