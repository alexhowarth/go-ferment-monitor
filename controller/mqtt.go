package controller

import (
	"fmt"
	"net/http"

	"github.com/alexhowarth/go-ferment-monitor/models"
	"github.com/gin-gonic/gin"
)

type MQTTController struct{}

var mqttModel = new(models.MQTTModel)

func (ctr *MQTTController) GetMQTTConfig(c *gin.Context) {
	conf, err := mqttModel.GetMQTTConfig()
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{}) //"error": "No configuration"})
	} else {
		c.JSON(http.StatusOK, conf)
	}
}

// func (ctr *MQTTController) CreateMQTTConfig(c *gin.Context) {

// 	var conf models.MQTTConf
// 	c.BindJSON(&conf)

// 	res, err := mqttModel.CreateMQTTConfig(&conf)
// 	if err != nil {
// 		fmt.Printf("Error in create: %+v\n", err)
// 		c.AbortWithStatusJSON(404, gin.H{"error": "Unable to create config"})
// 	} else {
// 		c.JSON(http.StatusOK, res)
// 	}
// }

func (ctr *MQTTController) UpdateMQTTConfig(c *gin.Context) {

	var conf models.MQTTConf
	c.BindJSON(&conf)

	res, err := mqttModel.UpdateMQTTConfig(&conf)
	if err != nil {
		fmt.Printf("Error in update: %+v\n", err)
		c.AbortWithStatusJSON(404, gin.H{"error": "Unable to update config"})
	} else {
		c.JSON(http.StatusOK, res)
	}
}
