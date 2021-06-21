package controllers

import (
	"net/http"
	"strconv"
	"webApp/models"

	"github.com/gin-gonic/gin"
)

func PostStartStopPage(c *gin.Context) {
	models.Config.Requests_state, _ = strconv.ParseBool(c.Request.URL.Query().Get("request_state"))
	c.JSON(
		http.StatusOK,
		gin.H{},
	)
	go models.Run_requests()
	
}
