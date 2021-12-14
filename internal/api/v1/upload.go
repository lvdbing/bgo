package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lvdbing/bgo/global"
	"github.com/lvdbing/bgo/internal/model"
	"github.com/lvdbing/bgo/internal/pkg/errcode"
	"github.com/lvdbing/bgo/internal/service"
)

var UploadApi = NewUploadApi()

type uploadApi struct {
}

func NewUploadApi() *uploadApi {
	return &uploadApi{}
}

// UploadFile godoc
// @Summary     上传文件
// @Description 上传文件
// @Tags        upload
// @Accept      multipart/form-data
// @Produce     multipart/form-data
// @Param       file body object true "上传文件"
// @Param       type body int true "上传文件类型"
// @Success     200 {object} string "上传成功的文件链接"
// @Failure     400 {object} model.RespError "Bad Request"
// @Failure     401 {object} model.RespError "Unauthorized"
// @Failure     403 {object} model.RespError "Forbidden"
// @Failure     500 {object} model.RespError "Internal Server Error"
// @Router      /api/v1/upload [post]
func (api *uploadApi) UploadFile(c *gin.Context) {
	resp := model.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		errResp := errcode.InvalidParams.WithDetails(err.Error())
		resp.SendError(errResp)
		return
	}
	fileType, _ := strconv.Atoi(c.PostForm("type"))
	if fileHeader == nil || fileType <= 0 {
		resp.SendError(errcode.InvalidParams)
		return
	}
	svc := service.NewService(c.Request.Context())
	fileInfo, err := svc.UploadFile(model.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("upload file err: %v", err)
		errResp := errcode.InvalidParams
		resp.SendError(errResp)
		return
	}
	resp.SendData(fileInfo.AccessUrl)
}
