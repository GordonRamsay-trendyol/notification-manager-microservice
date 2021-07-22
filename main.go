package main

import (
	"log"
	"time"

	"github.com/GordonRamsay-Trendyol/notification-manager/controller"
	"github.com/GordonRamsay-Trendyol/notification-manager/messaging"
)

var (
	controllerConfig = controller.Configuration{
		Host:         "localhost",
		Port:         8888,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	messagingConfig = messaging.Configuration{
		BootstrapServers: []string{},
		Topics:           []string{},
		Partitions:       0,
	}
)

func main() {
	log.SetPrefix("[notification-manager] ")

	log.Println("Application started...")
	go messaging.Start(messagingConfig)
	controller.Start(controllerConfig)
	time.Sleep(1 * time.Second)
}
