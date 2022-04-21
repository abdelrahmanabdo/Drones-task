package router

import (
	"github.com/gin-gonic/gin"
	"github.com/abdelrahmanabdo/drones-task/controllers/drone"
	"github.com/abdelrahmanabdo/drones-task/controllers/medication"
)

func SetupRoutes() {
	router := gin.Default()

	// Drones
	router.POST("/drones", drone.createDrone)
	router.GET("/drones/:drone_id", drone.getDrone)
	router.GET("/drones", drone.getDrones)

	// Medications
	router.POST("/medications", medication.createMedication)
	router.GET("/medications/:medication_id", medication.getMedication)
	router.GET("/medications", medication.getMedications)

	router.Run(":8080")
}