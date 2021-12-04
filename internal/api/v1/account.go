package v1

import (
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

// Register godoc
// @Summary     创建用户
// @Description 通过用户信息注册账号
// @Tags        account
// @Accept      json
// @Produce     json
// @Param       register body model.RegisterReq true "注册用户信息"
// @Success     200 {object} model.User "注册成功的用户信息"
// @Failure     400 {object} httputil.HTTPError "Bad Request"
// @Failure     401 {object} httputil.HTTPError "Unauthorized"
// @Failure     403 {object} httputil.HTTPError "Forbidden"
// @Failure     404 {object} httputil.HTTPError "Not Found"
// @Failure     500 {object} httputil.HTTPError "Internal Server Error"
// @Router      /account/register [post]
func (api *UserApi) Register(c *gin.Context) {

}

// Login godoc
// @Summary 用户登录
// @Description 用户登录
// @Tags account
// @Accept json
// @Produce json
// @Param login body model.LoginReq true "请求登录信息"
// @Success 200 {object} model.User "用户信息"
// @Failure 400 {object} httputil.HTTPError "Bad Request"
// @Failure 401 {object} httputil.HTTPError "Unauthorized"
// @Failure     403 {object} httputil.HTTPError "Forbidden"
// @Failure     404 {object} httputil.HTTPError "Not Found"
// @Failure 500 {object} httputil.HTTPError "Internal Server Error"
// @Router /account/login [post]
func (api *UserApi) Login(c *gin.Context) {

}
