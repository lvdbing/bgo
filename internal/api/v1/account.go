package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lvdbing/bgo/global"
	"github.com/lvdbing/bgo/internal/model"
	"github.com/lvdbing/bgo/internal/pkg/errcode"
	"github.com/lvdbing/bgo/internal/service"
)

var AccountApi = NewAccountApi()

type accountApi struct {
}

func NewAccountApi() *accountApi {
	return &accountApi{}
}

// Register godoc
// @Summary     创建用户
// @Description 通过用户信息注册账号
// @Tags        account
// @Accept      json
// @Produce     json
// @Param       register body model.RegisterReq true "注册用户信息"
// @Success     200 {object} model.User "注册成功的用户信息"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/register [post]
func (api *accountApi) Register(c *gin.Context) {
	resp := model.NewResponse(c)
	var req model.RegisterReq
	valid, errs := model.ValidAndBind(c, &req)
	if !valid {
		global.Logger.Errorf("register.ValidAndBind err: %v", errs)
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		resp.SendError(errResp)
		return
	}
	// _ = c.ShouldBindJSON(&req)
	// fmt.Println(req)

	svc := service.NewService(c.Request.Context())
	user, err := svc.Register(&req)
	if err != nil {
		resp.SendError(errcode.CustomerError.WithDetails(err.Error()))
		return
	}
	resp.SendData(*user)
}

// Login godoc
// @Summary     用户登录
// @Description 用户登录
// @Tags        account
// @Accept      json
// @Produce     json
// @Param       login body model.LoginReq true "请求登录信息"
// @Success     200 {object} model.User "用户信息"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/login [post]
func (api *accountApi) Login(c *gin.Context) {

}

// Get godoc
// @Summary     查询用户信息
// @Description 查询用户信息
// @Tags        account
// @Accept      json
// @Produce     json
// @Param       id body int true "用户id"
// @Success     200 {object} model.User "用户信息"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/account/get [get]
func (api *accountApi) Get(c *gin.Context) {

}

// List godoc
// @Summary     获取用户列表
// @Description 获取用户列表
// @Tags        account
// @Accept      json
// @Produce     json
// @Param       user body model.UserReq true "查询条件"
// @Success     200 {object} []model.User "用户列表"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/account/list [get]
func (api *accountApi) List(c *gin.Context) {

}

// Create godoc
// @Summary     新增用户
// @Description 新增用户
// @Tags        account
// @Accept      json
// @Produce     json
// @Param       user body model.User true "用户信息"
// @Success     200 {object} model.User "用户信息"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/account/create [post]
func (api *accountApi) Create(c *gin.Context) {

}

// Update godoc
// @Summary     修改用户
// @Description 修改用户
// @Tags        account
// @Accept      json
// @Produce     json
// @Param       user body model.User true "用户信息"
// @Success     200 {object} model.User "用户信息"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/account/update [put]
func (api *accountApi) Update(c *gin.Context) {

}

// Delete godoc
// @Summary     删除用户
// @Description 删除用户
// @Tags        account
// @Accept      json
// @Produce     json
// @Param       id body int true "用户ID"
// @Success     200 {string} string "成功"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/account/delete/:id [delete]
func (api *accountApi) Delete(c *gin.Context) {

}
