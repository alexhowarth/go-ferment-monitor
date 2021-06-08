package controller

import (
	"fmt"
	"net/http"

	"github.com/alexhowarth/go-ferment-monitor/models"
	"github.com/gin-gonic/gin"
)

type TiltController struct{}

var tiltModel = new(models.TiltModel)

func (ctr *TiltController) GetTilts(c *gin.Context) {
	//t := make([]models.Tilt, 0)
	t, _ := tiltModel.GetTilts()
	c.JSON(http.StatusOK, t)
}

func (ctr *TiltController) GetTilt(c *gin.Context) {
	//tiltModel.ID = c.Param("id")
	//var tilt models.Tilt
	t, err := tiltModel.GetTilt(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Record not found!"})
	} else {
		c.JSON(http.StatusOK, t)
	}
}

func (ctr *TiltController) CreateTilt(c *gin.Context) {
	//var tilt models.Tilt
	conf := models.TiltConf{}
	c.BindJSON(conf) // this is probably stupid - one per method? not sure why it's global
	t, err := tiltModel.CreateTilt(conf)
	if err != nil {
		fmt.Printf("Error in create: %+v\n", err)
		c.AbortWithStatusJSON(404, gin.H{"error": "Unable to create record"})
	} else {
		c.JSON(http.StatusOK, t)
	}
}

// func (ctr *TiltController) UpdateTilt(c *gin.Context) {
// 	// t, err := tiltModel.GetTilt(c.Param("id"))
// 	// if err != nil {
// 	// 	c.AbortWithStatusJSON(404, gin.H{"error": "Record not found!"})
// 	// }

// }

func (ctr *TiltController) UpdateTilt(c *gin.Context) {

	var conf models.TiltConf
	c.BindJSON(&conf)

	res, err := tiltModel.UpdateTilt(&conf)
	if err != nil {
		fmt.Printf("Error in update: %+v\n", err)
		c.AbortWithStatusJSON(404, gin.H{"error": "Unable to update tilt config"})
	} else {
		c.JSON(http.StatusOK, res)
	}
}
