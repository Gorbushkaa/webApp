package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webApp/models"

	"github.com/gin-gonic/gin"
)

type Timeout struct {
	Timeout string `json:"timeout"`
}


func PostSettingsPage(c *gin.Context) {
	var data Timeout
	bodyBytes, _ := c.GetRawData()
	json.Unmarshal(bodyBytes, &data)
	c.JSON(http.StatusOK,
		gin.H{})
	models.Config.RPM, _ = strconv.ParseFloat(data.Timeout, 64)
}
