package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lvdbing/bgo/global"
	"github.com/lvdbing/bgo/internal/helper/setting"
	"github.com/lvdbing/bgo/internal/helper/tracer"
	"github.com/lvdbing/bgo/internal/model"
	"github.com/lvdbing/bgo/internal/router"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// @title           bgo后台管理系统
// @version         1.0
// @description     使用Go和Angular创建的后台管理系统
// @termsOfService  http://beanlv.top/

// @contact.name   Bean
// @contact.url    http://beanlv.top/

// @host      localhost:8000
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

	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			global.Logger.Logger.Fatalf("ListenAndServe err: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Logger.Logger.Info("Shuting down server...")
	log.Println("Shuting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		global.Logger.Logger.Fatal("Shutdown err: %v", err)
	}
	global.Logger.Logger.Info("Server exiting...")
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

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func setupSetting() error {
	configFiles := []setting.ConfigFile{
		{Name: "password", Path: "config/", Type: "yaml"},
		{Name: "config", Path: "config/", Type: "yaml"},
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
	global.AppSetting.ContextTimeout *= time.Second

	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second

	err = s.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("RateLimiter", &global.LimiterSetting)
	if err != nil {
		return err
	}
	global.LimiterSetting.FillInterval *= time.Second

	return nil
}

func setupLogger() error {
	filename := global.AppSetting.LogPath + "/" + global.AppSetting.LogFilename + global.AppSetting.LogFileExt
	global.Logger = global.NewLogger(&lumberjack.Logger{
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

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"bgo-admin",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
