package main

import (
	"io"
	"log"
	"github.com/abdelrahmanabdo/drones-task/router"
)

func StartApplication() {
	router.SetupRoutes()
}

func main() {
	StartApplication()
}