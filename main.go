package main

import (
	"hao-admin/global"
	"hao-admin/internal/model"
	"hao-admin/internal/routers"
	"hao-admin/pkg/logger"
	"hao-admin/pkg/settings"
	"hao-admin/pkg/tracer"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.SetupSetting err:%v", err)
	}
	err = setupLogger()
	err = setUpDBEngine()
	if err != nil {
		log.Fatalf("init.setUpDBEngine err:%v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err:%v", err)
	}
	//global.Logger.Infof("%s: go-book/%s", "eddycjy", "blog_service")
}

// @title 博客系统
// @version 1.0
// @description 博客系统的详细描述
// @termsOfService http://127.0.0.1:8080/swagger/index.html
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("s.ListenAndServe err:%v", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown:", err)
	}
	log.Println("server exiting")

}

func setupSetting() error {
	setting, err := settings.NewSettings()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("DataBase", &global.DataBaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	global.AppSetting.DefaultContextTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
func setUpDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DataBaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("hao-admin", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil

}
