package routes

import (
	"fmt"
	"github.com/sjqzhang/gdi"
	//"fmt"
	"github.com/gin-gonic/gin"
	"go-web-mini/common"
	"go-web-mini/config"
	"go-web-mini/middleware"
	"time"
)

// 初始化
func InitRoutes() *gin.Engine {

	routerMap, _ := gdi.GetRouterInfo("controller")

	ctrls, err := gdi.AutoRegisterByPackagePatten(`controller*`)

	gdi.Init() //初始化对象池

	//设置模式
	gin.SetMode(config.Conf.System.Mode)

	// 创建带有默认中间件的路由:
	// 日志与恢复中间件

	// 创建不带中间件的路由:
	r := gin.New()
	r.Use(gin.Recovery())

	// 启用限流中间件
	// 默认每50毫秒填充一个令牌，最多填充200个
	fillInterval := time.Duration(config.Conf.RateLimit.FillInterval)
	capacity := config.Conf.RateLimit.Capacity
	r.Use(middleware.RateLimitMiddleware(time.Millisecond*fillInterval, capacity))

	// 启用全局跨域中间件
	r.Use(middleware.CORSMiddleware())

	// 启用操作日志中间件
	r.Use(middleware.OperationLogMiddleware())

	// 初始化JWT认证中间件
	authMiddleware, err := middleware.InitAuth()
	if err != nil {
		common.Log.Panicf("初始化JWT中间件失败：%v", err)
		panic(fmt.Sprintf("初始化JWT中间件失败：%v", err))
	}

	// 路由分组
	apiGroup := r.Group("/" + config.Conf.System.UrlPathPrefix)

	// 注册路由
	//InitBaseRoutes(apiGroup, authMiddleware) // 注册基础路由, 不需要jwt认证中间件,不需要casbin中间件
	//InitUserRoutes(apiGroup, authMiddleware)         // 注册用户路由, jwt认证中间件,casbin鉴权中间件
	//InitRoleRoutes(apiGroup, authMiddleware)         // 注册角色路由, jwt认证中间件,casbin鉴权中间件
	//InitMenuRoutes(apiGroup, authMiddleware)         // 注册菜单路由, jwt认证中间件,casbin鉴权中间件
	//InitApiRoutes(apiGroup, authMiddleware)          // 注册接口路由, jwt认证中间件,casbin鉴权中间件
	//InitOperationLogRoutes(apiGroup, authMiddleware) // 注册操作日志路由, jwt认证中间件,casbin鉴权中间件

	apiGroup.POST("/base/login", authMiddleware.LoginHandler)
	apiGroup.POST("/base/logout", authMiddleware.LogoutHandler)
	apiGroup.POST("/base/refreshToken", authMiddleware.RefreshHandler)

	common.Log.Info("初始化路由完成！")

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

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for _, o := range ctrls { //自动绑定路由
		ctrlName := o.Elem().Type().Name()

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

			if o.Method(i).Type().NumIn() > 2 {
				continue
			}

			if o.Method(i).Type().NumIn() == 2 {
				if o.Method(i).Type().In(0).String() != "*gin.Context" {
					continue
				}
			}

			var mds []gin.HandlerFunc
			for _, m := range v.Middlewares {
				if f, ok := middlewaresMaping[m]; ok {
					mds = append(mds, f)
				}
			}


			switch x.(type) {
			case func(*gin.Context):
				{
					mds = append(mds, x.(func(*gin.Context)))
					r.Handle(v.Method, v.Uri, mds...)

				}
			default:
				mds = append(mds, middleware.BinderMiddleware(o.Method(i)))
				r.Handle(v.Method, v.Uri, mds...)
			}
		}

	}

	return r
}
