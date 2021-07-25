package main

import (
	"log"
	"time"

	"github.com/GordonRamsay-trendyol/notification-manager-microservice/controller"
)

var (
	controllerConfig = controller.Configuration{
		Host:         "localhost",
		Port:         8888,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
)

func main() {
	log.SetPrefix("[notification-manager] ")

	log.Println("Application started...")
	controller.Start(controllerConfig)
}
