package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lvdbing/bgo/global"
	"github.com/lvdbing/bgo/internal/pkg/errcode"
)

type Pager struct {
	Page     int `json:"page" form:"page"`
	Pagesize int `json:"pagesize" form:"pagesize"`
}

type RespList struct {
	Data     interface{} `json:"data"`
	Page     int         `json:"page"`
	Pagesize int         `json:"pagesize"`
	Total    int         `json:"total"`
}

type RespError struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Details []string `json:"details"`
}

type Response struct {
	Ctx *gin.Context
}

func NewResponse(c *gin.Context) *Response {
	return &Response{Ctx: c}
}

func (r *Response) GetPager() Pager {
	var p Pager
	_ = r.Ctx.ShouldBindJSON(&p)

	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Pagesize <= 0 {
		p.Pagesize = global.AppSetting.DefaultPagesize
	}
	if p.Pagesize > global.AppSetting.MaxPagesize {
		p.Pagesize = global.AppSetting.MaxPagesize
	}
	return p
}

func (r *Response) GetOffset(p Pager) int {
	offset := 0
	if p.Page > 0 {
		offset = (p.Page - 1) * p.Pagesize
	}
	return offset
}

func (r *Response) SendData(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) SendList(data interface{}, total int) {
	p := r.GetPager()
	var resp RespList
	resp.Data = data
	resp.Page = p.Page
	resp.Pagesize = p.Pagesize
	resp.Total = total
	r.Ctx.JSON(http.StatusOK, resp)
}

func (r *Response) SendError(err *errcode.Error) {
	var resp RespError
	resp.Code = err.Code()
	resp.Msg = err.Msg()
	resp.Details = err.Details()
	r.Ctx.JSON(err.StatusCode(), resp)
}
