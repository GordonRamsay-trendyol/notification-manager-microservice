package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

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

	router := mux.NewRouter()
	router.PathPrefix("/user/notification")

	router.HandleFunc("/push-id", updatePushID).Methods(http.MethodPut)
	router.HandleFunc("/settings", updateNotificationSettings).Methods(http.MethodPut)

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

func updatePushID(w http.ResponseWriter, r *http.Request) {
	// Get push id from path variables or request parameters
	pathVariables := mux.Vars(r)
	log.Println(pathVariables)
}

func updateNotificationSettings(w http.ResponseWriter, r *http.Request) {

}
