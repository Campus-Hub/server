package api

import (
	"encoding/json"
	"fmt"

	"github.com/Campus-Hub/server/config"
	"github.com/Campus-Hub/server/internal/service"
	"github.com/go-playground/validator/v10"
)

// ErrorResponse 返回错误消息
func errorResponse(err error) service.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := config.T(fmt.Sprintf("Field.%s", e.Field))
			tag := config.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return service.Response{
				Status: 401,
				Msg:    fmt.Sprintf("%s%s", field, tag),
				Error:  err.Error(),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return service.Response{
			Status: 401,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	return service.Response{
		Status: 401,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
