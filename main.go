package main

import (
	"github.com/lvdbing/bgo/internal/router"
)

// @title           bgo后台管理系统
// @version         1.0
// @description     使用Go和Angular创建的后台管理系统
// @termsOfService  http://beanlv.top/

// @contact.name   Bean
// @contact.url    http://beanlv.top/
// @contact.email  lvduanbing@126.com

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	r := router.NewRouter()

	r.Run(":8080")
}
