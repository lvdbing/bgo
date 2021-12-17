package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lvdbing/bgo/global"
	"github.com/lvdbing/bgo/pkg/errcode"
)

type Pager struct {
	Page     int `json:"page" form:"page"`         // 页码
	Pagesize int `json:"pagesize" form:"pagesize"` // 每页数量
}

type RespList struct {
	Data     interface{} `json:"data"`     // 响应数据
	Page     int         `json:"page"`     // 页码
	Pagesize int         `json:"pagesize"` // 每页数量
	Total    int         `json:"total"`    // 总数量
}

type RespError struct {
	Code    int      `json:"code"`    // 错误码
	Msg     string   `json:"msg"`     // 错误消息
	Details []string `json:"details"` // 错误详情
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

func (r *Response) SendError(err error, details ...string) {
	var resp RespError
	respErr, ok := err.(*errcode.Error)
	if ok {
		resp.Code = respErr.Code()
		resp.Msg = respErr.Msg()
		if len(details) > 0 {
			respErr = respErr.WithDetails(details...)
		}
		resp.Details = respErr.Details()
		r.Ctx.JSON(respErr.StatusCode(), resp)
	} else {
		resp.Msg = err.Error()
		resp.Code = errcode.CustomerError.Code()
		resp.Details = details
		r.Ctx.JSON(http.StatusInternalServerError, resp)
	}
}
