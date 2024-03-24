package handler

import (
	errs "github.com/harin-h/rest-api-err"

	"github.com/gin-gonic/gin"
)

func handlerError(ctx *gin.Context, err error) {
	switch e := err.(type) {
	case errs.AppError:
		ctx.String(e.Code, e.Message)
	}
}
