package adapter

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRouters() *gin.Engine {
	router := gin.New()
	//var logger = log.NewLogger("gin.log")
	//// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	//// By default gin.DefaultWriter = os.Stdout
	//router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//	// write to file by logger
	//	logger.WithFields(logrus.Fields{
	//		"client_ip": param.ClientIP,
	//		"proto":     param.Request.Proto,
	//		"UA":        param.Request.UserAgent(),
	//		"error_msg": param.ErrorMessage,
	//	}).Debugf("%s %s %d %s", param.Method, param.Path, param.StatusCode, param.Latency)
	//
	//	// your custom format
	//	return fmt.Sprintf("%s - [%s] \"%s %s %d %s %s \"%s\"\n",
	//		param.ClientIP,
	//		param.TimeStamp.Format(time.RFC3339),
	//		param.Method,
	//		param.Path,
	//		param.StatusCode,
	//		param.Latency,
	//		param.Request.Proto,
	//		// param.Request.UserAgent(),
	//		param.ErrorMessage,
	//	)
	//}))
	router.Use(gin.Recovery())

	api := router.Group("/apibpmn")
	api.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	api.POST("/instance", CreateInstance)
	api.GET("/instance/:instance_id", GetInstance)
	api.GET("/instance/:instance_id/variables", GetInstanceVariables)
	api.POST("/instance/:instance_id/variables", SetInstanceVariables)
	//api.GET("/instance/:instance_id/commands", GetAvailableCommands)
	//api.POST("/instance/:instance_id/command", CommandExecution)
	api.GET("/instance/:instance_id/transitions", GetInstanceTransitions)
	return router
}
