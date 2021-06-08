package server

import (
	"github.com/alexhowarth/go-ferment-monitor/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "PUT", "POST", "OPTIONS"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		//AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "http://localhost:5000"
		// },
		//MaxAge: 12 * time.Hour,
	}))
	router.Use(static.Serve("/", static.LocalFile("./web/public", true)))

	v1 := router.Group("v1")
	{
		mqttGroup := v1.Group("mqtt")
		{
			mqttController := &controller.MQTTController{}
			mqttGroup.GET("/", mqttController.GetMQTTConfig)
			//mqttGroup.POST("/", mqttController.CreateMQTTConfig)
			mqttGroup.PUT("/", mqttController.UpdateMQTTConfig)
			//mqttGroup.PUT("/restart", mqttController.RestartMQTTConsumer)
		}
		tg := v1.Group("tilt")
		{
			tiltController := &controller.TiltController{}
			tg.GET("/", tiltController.GetTilts)
			tg.PUT("/", tiltController.UpdateTilt)
		}

	}

	return router
}
