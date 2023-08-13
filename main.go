package main

import (
	"context"
	"fmt"
	"github.com/sjqzhang/gdi"
	"go-web-mini/config"
	"go-web-mini/global"
	"go-web-mini/middleware"
	"go-web-mini/repository"
	"go-web-mini/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Swagger Example API
// @version v1
// @description This is a sample server Petstore server.
// @termsOfService http://localhost
// @BasePath /api
// @schemes http https
//@securityDefinitions.apikey JWT
//@in header
//@name Authorization
func main() {

	// 加载配置文件到全局配置结构体
	config.InitConfig()

	// 初始化日志
	global.InitLogger()

	// 初始化数据库(mysql)
	global.InitMysql()

	global.InitRedis()

	// 初始化casbin策略管理器
	global.InitCasbinEnforcer()

	// 初始化Validator数据校验
	global.InitValidate()

	// 初始化mysql数据
	global.InitData()

	//bs,err:=gdiEmbededFiles.ReadFile("config/config.go")
	//fmt.Println(string(bs),err)

	gdi.GenGDIRegisterFile(true)

	// 注册所有路由
	r := routes.InitRoutes()

	// 操作日志中间件处理日志时没有将日志发送到rabbitmq或者kafka中, 而是发送到了channel中
	// 这里开启3个goroutine处理channel将日志记录到数据库
	logRepository := gdi.Get(&repository.OperationLogRepository{}).(*repository.OperationLogRepository)
	for i := 0; i < 3; i++ {
		go logRepository.SaveOperationLogChannel(nil, middleware.OperationLogChan)
	}

	host := "localhost"
	port := config.Conf.System.Port

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Log.Fatalf("listen: %s\n", err)
		}
	}()

	global.Log.Info(fmt.Sprintf("Server is running at %s:%d/%s", host, port, config.Conf.System.UrlPathPrefix))

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Log.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Log.Fatal("Server forced to shutdown:", err)
	}

	global.Log.Info("Server exiting!")

}
