package v1

import "github.com/gin-gonic/gin"

var PermitApi = NewPermitApi()

type permitApi struct {
}

func NewPermitApi() *permitApi {
	return &permitApi{}
}

// Get godoc
// @Summary     获取权限信息
// @Description 获取权限信息
// @Tags        permit
// @Accept      json
// @Produce     json
// @Param       id body int true "权限id"
// @Success     200 {object} model.Permit "权限信息"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     404 {object} model.RespError "Not Found"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/permit/get [get]
func (api *permitApi) Get(c *gin.Context) {

}

// List godoc
// @Summary     获取权限列表
// @Description 获取权限列表
// @Tags        permit
// @Accept      json
// @Produce     json
// @Param       permit body model.PermitReq true "查询条件"
// @Success     200 {object} []model.Permit "权限列表"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     404 {object} model.RespError "Not Found"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/permit/list [get]
func (api *permitApi) List(c *gin.Context) {

}

// Create godoc
// @Summary     新增权限
// @Description 新增权限
// @Tags        permit
// @Accept      json
// @Produce     json
// @Param       permit body model.Permit true "权限信息"
// @Success     200 {object} model.Permit "权限信息"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     404 {object} model.RespError "Not Found"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/permit/create [post]
func (api *permitApi) Create(c *gin.Context) {

}

// Update godoc
// @Summary     修改权限
// @Description 修改权限
// @Tags        permit
// @Accept      json
// @Produce     json
// @Param       permit body model.Permit true "权限信息"
// @Success     200 {object} model.Permit "权限信息"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     404 {object} model.RespError "Not Found"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/permit/update [put]
func (api *permitApi) Update(c *gin.Context) {

}

// Delete godoc
// @Summary     删除权限
// @Description 删除权限
// @Tags        permit
// @Accept      json
// @Produce     json
// @Param       id body int true "权限ID"
// @Success     200 {string} string "成功"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     404 {object} model.RespError "Not Found"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/permit/delete/:id [delete]
func (api *permitApi) Delete(c *gin.Context) {

}
