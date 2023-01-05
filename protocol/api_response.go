package protocol

import (
	"fmt"
	"lecture/WBABEProject-23/logger"

	"github.com/gin-gonic/gin"
)

type ApiResponse[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

func (e *ApiResponse[any]) Response(c *gin.Context) {
	c.JSON(e.Code, e)
}

func Success(code int) *ApiResponse[any] {
	return SuccessDataAndCustom("success", code, "ok")
}

func SuccessData[T any](data T, code int) *ApiResponse[any] {
	return SuccessDataAndCustom(data, code, "ok")
}

func SuccessDataAndCustom[T any](data T, code int, MSG string) *ApiResponse[any] {
	return &ApiResponse[any]{
		Code:    code,
		Data:    data,
		Message: MSG,
	}
}

func Fail(e error, code int) *ApiResponse[any] {
	if e == nil {
		return FailCustomMessage(e, "Should use Custom error message.", code)
	}
	return FailCustomMessage(e, e.Error(), code)
}

func FailCustomMessage(e error, MSG string, code int) *ApiResponse[any] {
	errLogMSG := fmt.Sprintf("Error: %v, Message: %s", e, MSG)
	logger.Error(errLogMSG)
	return &ApiResponse[any]{
		Code:    code,
		Data:    nil,
		Message: MSG,
	}
}
