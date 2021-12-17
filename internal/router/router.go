package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lvdbing/bgo/docs"
	"github.com/lvdbing/bgo/global"

	v1 "github.com/lvdbing/bgo/internal/api/v1"
	"github.com/lvdbing/bgo/internal/helper/limiter"
	"github.com/lvdbing/bgo/internal/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type RouterGroup struct {
}

func NewRouter() *gin.Engine {
	r := gin.New()
	// if global.ServerSetting.RunMode == "debug" {
	// 	r.Use(gin.Logger())
	// 	r.Use(gin.Recovery())
	// } else {
	// 	r.Use(middleware.AccessLog()) // 访问日志记录
	// 	r.Use(middleware.Recovery())  // 异常处理
	// }
	r.Use(middleware.AccessLog())    // 访问日志记录
	r.Use(middleware.Recovery())     // 异常处理
	r.Use(middleware.Cors())         // 跨域
	r.Use(middleware.Translations()) // 国际化

	var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
		limiter.BucketRule{
			Key:          "/auth",
			FillInterval: global.LimiterSetting.FillInterval,
			Capacity:     global.LimiterSetting.Capacity,
			Quantum:      global.LimiterSetting.Quantum,
		},
	)
	r.Use(middleware.RateLimiter(methodLimiters))                      // 接口限流控制
	r.Use(middleware.ContextTimeout(global.AppSetting.ContextTimeout)) // 超时控制
	r.Use(middleware.Tracing())                                        // 链路追踪

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadPath)) // 提供静态资源的访问。

	// 开放的api，不需要鉴权。
	pubRouter := r.Group("/api/v1")
	// pubRouter.Use(middleware.Cors()) // 跨域
	pubRouter.POST("/register", v1.AccountApi.Register)
	pubRouter.POST("/login", v1.AccountApi.Login)
	pubRouter.POST("/upload", v1.UploadApi.UploadFile) // 公开的上传文件api。

	// 私密的api，需要鉴权。
	priRouter := r.Group("/api/v1")
	priRouter.Use(middleware.JWT()) // JWT鉴权
	groupPrimaryRouter(priRouter)

	return r
}

func groupPrimaryRouter(priRouter *gin.RouterGroup) {
	priRouter.POST("/priupload", v1.UploadApi.UploadFile) // 需要鉴权的上传文件api。

	accRouter := priRouter.Group("/account")
	accRouter.GET("/get/:id", v1.AccountApi.Get)          // 查询用户信息
	accRouter.GET("/list", v1.AccountApi.List)            // 获取用户列表
	accRouter.POST("/create", v1.AccountApi.Create)       // 创建用户
	accRouter.PUT("/update", v1.AccountApi.Update)        // 更新用户
	accRouter.DELETE("/delete/:id", v1.AccountApi.Delete) // 删除用户

	roleRouter := priRouter.Group("/role")
	roleRouter.GET("/get/:id", v1.RoleApi.Get)          // 查询角色信息
	roleRouter.GET("/list", v1.RoleApi.List)            // 获取角色列表
	roleRouter.POST("/create", v1.RoleApi.Create)       // 新增角色
	roleRouter.PUT("/update", v1.RoleApi.Update)        // 更新角色
	roleRouter.DELETE("/delete/:id", v1.RoleApi.Delete) // 删除角色

	permitRouter := priRouter.Group("/permit")
	permitRouter.GET("/get/:id", v1.PermitApi.Get)          // 查询权限信息
	permitRouter.GET("/list", v1.PermitApi.List)            // 获取权限列表
	permitRouter.POST("/create", v1.PermitApi.Create)       // 新增权限
	permitRouter.PUT("/update", v1.PermitApi.Update)        // 更新权限
	permitRouter.DELETE("/delete/:id", v1.PermitApi.Delete) // 删除权限
}
