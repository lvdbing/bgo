package v1

import (
	"github.com/gin-gonic/gin"
)

var RoleApi = NewRoleApi()

type roleApi struct {
}

func NewRoleApi() *roleApi {
	return &roleApi{}
}

// Get godoc
// @Summary     获取角色信息
// @Description 获取角色信息
// @Tags        role
// @Accept      json
// @Produce     json
// @Param       id body int true "角色id"
// @Success     200 {object} model.Role "角色信息"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     404 {object} model.RespError "Not Found"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/role/get [get]
func (api *roleApi) Get(c *gin.Context) {

}

// List godoc
// @Summary     获取角色列表
// @Description 获取角色列表
// @Tags        role
// @Accept      json
// @Produce     json
// @Param       role body model.RoleReq true "查询条件"
// @Success     200 {object} []model.Role "角色列表"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     404 {object} model.RespError "Not Found"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/role/list [get]
func (api *roleApi) List(c *gin.Context) {

}

// Create godoc
// @Summary     新增角色
// @Description 新增角色
// @Tags        role
// @Accept      json
// @Produce     json
// @Param       role body model.Role true "角色信息"
// @Success     200 {object} model.Role "角色信息"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     404 {object} model.RespError "Not Found"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/role/create [post]
func (api *roleApi) Create(c *gin.Context) {

}

// Update godoc
// @Summary     修改角色
// @Description 修改角色
// @Tags        role
// @Accept      json
// @Produce     json
// @Param       role body model.Role true "角色信息"
// @Success     200 {object} model.Role "角色信息"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     404 {object} model.RespError "Not Found"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/role/update [put]
func (api *roleApi) Update(c *gin.Context) {

}

// Delete godoc
// @Summary     删除角色
// @Description 删除角色
// @Tags        role
// @Accept      json
// @Produce     json
// @Param       id body int true "角色ID"
// @Success     200 {string} string "成功"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     404 {object} model.RespError "Not Found"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/role/delete/:id [delete]
func (api *roleApi) Delete(c *gin.Context) {

}
