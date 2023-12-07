package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Campus-Hub/server/config"
	"github.com/go-playground/validator/v10"
)

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

// TokenData 带有 token 的 Data 结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) Response {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, e := range ve {
			field := config.T(fmt.Sprintf("Field.%s", e.Field))
			tag := config.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return Response{
				Status: 401,
				Msg:    fmt.Sprintf("%s%s", field, tag),
				Error:  fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return Response{
			Status: 401,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	return Response{
		Status: 401,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}

func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}
