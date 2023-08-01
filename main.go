package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sjqzhang/gdi"
	"go-web-mini/common"
	"go-web-mini/config"
	"go-web-mini/middleware"
	"go-web-mini/repository"
	"go-web-mini/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// 加载配置文件到全局配置结构体
	config.InitConfig()

	// 初始化日志
	common.InitLogger()

	// 初始化数据库(mysql)
	common.InitMysql()

	// 初始化casbin策略管理器
	common.InitCasbinEnforcer()

	// 初始化Validator数据校验
	common.InitValidate()

	// 初始化mysql数据
	common.InitData()

	//bs,err:=gdiEmbededFiles.ReadFile("config/config.go")
	//fmt.Println(string(bs),err)

	gdi.GenGDIRegisterFile(true)

	middlewaresMaping := make(map[string]gin.HandlerFunc)

	//middlewaresMaping["ratelimit"]=middleware.RateLimitMiddleware(time.Millisecond*time.Duration(config.Conf.RateLimit.FillInterval),config.Conf.RateLimit.Capacity)
	middlewaresMaping["cors"] = middleware.CORSMiddleware()
	middlewaresMaping["operationlog"] = middleware.OperationLogMiddleware()
	auth, err := middleware.InitAuth()
	if err != nil {
		panic(err)
	}
	middlewaresMaping["auth"] = auth.MiddlewareFunc()
	middlewaresMaping["casbin"] = middleware.CasbinMiddleware()
	middlewaresMaping["transition"] = middleware.TransitionMiddleware()

	routerMap, _ := gdi.GetRouterInfo("controller")

	ctrls, err := gdi.AutoRegisterByPackagePatten(`controller*`)

	if err != nil {
		fmt.Println(err)
		return
	}

	r := gin.Default()

	for _, o := range ctrls {
		ctrlName := o.Elem().Type().Name()
		if ctrlName=="NewsController" {
			//fmt.Println(ctrlName)
		}

		for i := 0; i < o.NumMethod(); i++ {
			if o.NumMethod() == 0 {
				continue
			}
			if !o.Type().Method(i).IsExported() {
				continue
			}
			var v gdi.RouterInfo
			var ok bool
			methodName := o.Type().Method(i).Name
			key := fmt.Sprintf("%v.%v", ctrlName, methodName)
			if v, ok = routerMap[key]; !ok {
				v.Method = "POST"
				v.Uri = fmt.Sprintf("/%v/%v", ctrlName, methodName)
				v.Handler = methodName
			}

			x := o.Method(i).Interface()

			switch x.(type) {
			case func(*gin.Context):
				{
					var mds []gin.HandlerFunc

					for _, m := range v.Middlewares {
						if f, ok := middlewaresMaping[m]; ok {
							mds = append(mds, f)
						}
					}
					mds = append(mds, x.(func(*gin.Context)))

					r.Handle(v.Method, v.Uri, mds...)

				}

			}
		}

	}

	gdi.Init()

	// 操作日志中间件处理日志时没有将日志发送到rabbitmq或者kafka中, 而是发送到了channel中
	// 这里开启3个goroutine处理channel将日志记录到数据库
	logRepository := repository.NewOperationLogRepository()
	for i := 0; i < 3; i++ {
		go logRepository.SaveOperationLogChannel(nil, middleware.OperationLogChan)
	}

	// 注册所有路由
	//r := routes.InitRoutes()

	routes.InitRoutes(r, auth)

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
			common.Log.Fatalf("listen: %s\n", err)
		}
	}()

	common.Log.Info(fmt.Sprintf("Server is running at %s:%d/%s", host, port, config.Conf.System.UrlPathPrefix))

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	common.Log.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		common.Log.Fatal("Server forced to shutdown:", err)
	}

	common.Log.Info("Server exiting!")

}
