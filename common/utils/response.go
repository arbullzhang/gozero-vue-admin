package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code int 		 `json:"code"`
	Data interface{} `json:"data"`
	Msg  string 	 `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

type H map[string]interface{}

func Result(code int, data interface{}, msg string, w http.ResponseWriter) {
	// 开始时间
	retJson, err := json.Marshal(Response{
		code,
		data,
		msg,
	})
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(string(retJson)))
	}
}

func Ok(w http.ResponseWriter) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", w)
}

func OkWithMessage(message string, w http.ResponseWriter) {
	Result(SUCCESS, map[string]interface{}{}, message, w)
}

func OkWithData(data interface{}, w http.ResponseWriter) {
	Result(SUCCESS, data, "操作成功", w)
}

func OkWithDetailed(data interface{}, message string, w http.ResponseWriter) {
	Result(SUCCESS, data, message, w)
}

func Fail(w http.ResponseWriter) {
	Result(ERROR, map[string]interface{}{}, "操作失败", w)
}

func FailWithMessage(message string, w http.ResponseWriter) {
	Result(ERROR, map[string]interface{}{}, message, w)
}

func FailWithDetailed(data interface{}, message string, w http.ResponseWriter) {
	Result(ERROR, data, message, w)
}