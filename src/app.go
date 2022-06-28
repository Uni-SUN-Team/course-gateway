package src

import (
	"strings"
	"unisun/api/course-listener/docs"
	"unisun/api/course-listener/src/component"
	"unisun/api/course-listener/src/controller"
	"unisun/api/course-listener/src/route"
	"unisun/api/course-listener/src/service"
	"unisun/api/course-listener/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io

// @license.name  MIT License Copyright (c) 2022 Uni-SUN-Team
// @license.url   https://api.unisun.dynu.com/course-listenner/api/license

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func App() *gin.Engine {
	docs.SwaggerInfo.Title = "COURSE LISTENER API"
	docs.SwaggerInfo.Description = ""
	docs.SwaggerInfo.Version = viper.GetString("app.version")
	docs.SwaggerInfo.Host = viper.GetString("app.host")
	docs.SwaggerInfo.BasePath = strings.Join([]string{viper.GetString("app.context_path"), viper.GetString("app.root_path")}, "")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	healCheckController := controller.NewControllerHealthCheckAdapter()
	utilAdap := utils.NewUtilsHTTPRequestAdapter()
	serviceAdap := service.NewServiceConsumerAdapter(utilAdap)
	componentAdap := component.NewServiceConsumerAdap(serviceAdap)
	controllerAdap := controller.NewControllerConsumerAdapter(serviceAdap, componentAdap)
	routeAdap := route.NewRouteConsumerAdapter(controllerAdap)

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	g := r.Group(strings.Join([]string{viper.GetString("app.context_path"), viper.GetString("app.root_path"), "/v1"}, ""))
	{
		g.GET("/healcheck", healCheckController.HealthCheckHandler)
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		g.StaticFile("/license", "./LICENSE")
		routeAdap.Consumer(g)
	}
	return r
}
