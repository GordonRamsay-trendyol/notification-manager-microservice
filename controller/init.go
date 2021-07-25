package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GordonRamsay-trendyol/notification-manager-microservice/service"
	"github.com/gorilla/mux"
)

type Configuration struct {
	Host         string
	Port         int
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

func Start(config Configuration) {
	log.Println("HTTP Server configurations started..")

	userService := service.NewUserService()
	controller := NewUserController(userService)

	router := mux.NewRouter()

	router.HandleFunc("/user/notification", controller.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/user/notification/push-id", controller.UpdatePushID).Methods(http.MethodPut)
	router.HandleFunc("/user/notification/settings", controller.UpdateNotificationSettings).Methods(http.MethodPut)

	startServer(config, router)
}

func startServer(config Configuration, router *mux.Router) {
	serverAddress := fmt.Sprintf("%s:%d", config.Host, config.Port)

	server := http.Server{
		Handler:      router,
		Addr:         serverAddress,
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
	}

	log.Printf("HTTP Server configurations finished.. Server will start at address: %v\n", serverAddress)
	log.Fatal(server.ListenAndServe())
}
