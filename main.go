package main

import (
	"log"
	"net/http"
	"time"

	"github.com/lvdbing/bgo/global"
	"github.com/lvdbing/bgo/internal/model"
	"github.com/lvdbing/bgo/internal/pkg/logger"
	"github.com/lvdbing/bgo/internal/pkg/setting"
	"github.com/lvdbing/bgo/internal/router"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
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
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        r,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func setupSetting() error {
	configFiles := []setting.ConfigFile{
		{Name: "config", Path: "config/", Type: "yaml"},
		{Name: "password", Path: "config/", Type: "yaml"},
	}
	s, err := setting.NewSetting(configFiles...)
	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger() error {
	filename := global.AppSetting.LogPath + "/" + global.AppSetting.LogFilename + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     31,   //days
		Compress:   true, // disabled by default
		LocalTime:  true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupDBEngine() error {
	var err error
	global.UserDB, err = model.NewDBEngine(global.DatabaseSetting, "user")

	return err
}
