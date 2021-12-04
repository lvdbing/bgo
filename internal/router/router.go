package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lvdbing/bgo/docs"
	v1 "github.com/lvdbing/bgo/internal/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/account/register", v1.NewAccountApi().Register)
		apiv1.POST("/account/login", v1.NewAccountApi().Login)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
