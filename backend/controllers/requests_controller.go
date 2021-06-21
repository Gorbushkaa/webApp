package controllers

import (
	"net/http"
	"webApp/db"
	"webApp/settings"

	"github.com/gin-gonic/gin"
)

func RequestsPage(c *gin.Context) {
	requests := db.GetRequests()
	c.JSON(http.StatusOK,
		gin.H{"data": requests})
}

func GetLastRequestsTime(c *gin.Context) {
	delay_request := db.GetLastRequestByUrl(settings.DELAY_URL)
	status_request := db.GetLastRequestByUrl(settings.STATUS_URL)
	c.JSON(
		http.StatusOK,
		gin.H{
			"status_request": status_request.Req_time.Seconds(),
			"delay_request":  delay_request.Req_time.Seconds(),
		},
	)

}
