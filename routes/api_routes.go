package routes

//func InitApiRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
//	apiController := gdi.Get(&controller.ApiController{}).(*controller.ApiController)
//	router := r.Group("/api")
//	// 开启jwt认证中间件
//	router.Use(authMiddleware.MiddlewareFunc())
//	// 开启casbin鉴权中间件
//	router.Use(middleware.CasbinMiddleware())
//	{
//		router.GET("/list", apiController.GetApis)
//		router.GET("/tree", apiController.GetApiTree)
//		router.POST("/create", apiController.CreateApi)
//		router.PATCH("/update/:apiId", apiController.UpdateApiById)
//		router.DELETE("/delete/batch", apiController.BatchDeleteApiByIds)
//	}
//
//	return r
//}
