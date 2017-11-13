package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/revel/revel"
)

// revel.Controller をラップした ApiV1Controller を定義する
type ApiV1Controller struct {
	*revel.Controller
}

// エラーの際に返す Json 用の構造体
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 正常な際に返す Json 用の構造体(今回は1種類で統一する)
type Response struct {
	Results interface{} `json:"results"`
}

// 引数として渡されて interface にリクエストの Json の値を格納する
func (c ApiV1Controller) BindParams(s interface{}) error {
	return json.Unmarshal(c.Params.JSON, s)
}

// Bad Request Error を返すやつ
func (c ApiV1Controller) HandleBadRequestError(s string) revel.Result {
	c.Response.Status = http.StatusBadRequest
	r := ErrorResponse{c.Response.Status, s}
	return c.RenderJSON(r)
}

// Not Found Error を返すやつ
func (c ApiV1Controller) HandleNotFoundError(s string) revel.Result {
	c.Response.Status = http.StatusNotFound
	r := ErrorResponse{c.Response.Status, s}
	return c.RenderJSON(r)
}

// Internal Server Error を返すやつ
func (c ApiV1Controller) HandleInternalServerError(s string) revel.Result {
	c.Response.Status = http.StatusInternalServerError
	r := ErrorResponse{c.Response.Status, s}
	return c.RenderJSON(r)
}
