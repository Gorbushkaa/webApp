package main

import (
	controllers "webApp/controllers"
)

func initializeRoutes() {
	router.GET("/requests", controllers.RequestsPage)
	router.GET("/get_last_requests", controllers.GetLastRequestsTime)
	router.POST("/settings", controllers.PostSettingsPage)
	router.POST("/", controllers.PostStartStopPage)

}
