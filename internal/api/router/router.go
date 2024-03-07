package router

import (
	"fmt"
	"io"
	"neversitup-test-template/internal/api/controllers"
	"neversitup-test-template/internal/api/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	app := gin.New()

	// Logging to a file.
	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())

	v1 := app.Group("api/v1")
	{
		v1.GET("/health-check", controllers.HealthCheck)
		hellow := v1.Group("/hellow")
		{
			hellow.GET("/neversitup", controllers.HellowNeverSitup)
		}
	}

	return app
}
