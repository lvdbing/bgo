package v1

import (
	"github.com/gin-gonic/gin"
)

type AccountApi struct {
}

func NewAccountApi() *AccountApi {
	return &AccountApi{}
}

// Register godoc
// @Summary     创建用户
// @Description 通过用户信息注册账号
// @Tags        account
// @Accept      json
// @Produce     json
// @Param       register body model.RegisterReq true "注册用户信息"
// @Success     200 {object} model.User "注册成功的用户信息"
// @Failure     400 {object} model.Error "Bad Request"
// @Failure     401 {object} model.Error "Unauthorized"
// @Failure     403 {object} model.Error "Forbidden"
// @Failure     404 {object} model.Error "Not Found"
// @Failure     500 {object} model.Error "Internal Server Error"
// @Router      /account/register [post]
func (api *AccountApi) Register(c *gin.Context) {

}

// Login godoc
// @Summary     用户登录
// @Description 用户登录
// @Tags        account
// @Accept      json
// @Produce     json
// @Param       login body model.LoginReq true "请求登录信息"
// @Success     200 {object} model.User "用户信息"
// @Failure     400 {object} model.Error "Bad Request"
// @Failure     401 {object} model.Error "Unauthorized"
// @Failure     403 {object} model.Error "Forbidden"
// @Failure     404 {object} model.Error "Not Found"
// @Failure     500 {object} model.Error "Internal Server Error"
// @Router      /account/login [post]
func (api *AccountApi) Login(c *gin.Context) {

}
