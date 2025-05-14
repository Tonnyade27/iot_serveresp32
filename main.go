package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/Tonnyade27/iot-server/db"
	"github.com/Tonnyade27/iot-server/mqtt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()
	mqtt.InitMQTT()
	mqtt.Subscribe(os.Getenv("MQTT_TOPIC"))

	r := mux.NewRouter()
	r.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
		msg := r.URL.Query().Get("msg")
		if msg == "" {
			http.Error(w, "Message is required", http.StatusBadRequest)
			return
		}
		mqtt.Publish(os.Getenv("MQTT_TOPIC"), msg)
		fmt.Fprintf(w, "Published message: %s", msg)
	}).Methods("GET")

	fmt.Println("HTTP server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

