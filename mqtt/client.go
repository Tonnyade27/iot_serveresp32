package mqtt

import (
	"encoding/json"
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/username/iot-server/db"
)

type SensorData struct {
	Name     string `json:"name"`
	Power    string `json:"power,omitempty"`
	Voltage  string `json:"voltage"`
	Current  string `json:"current"`
	PowerAC  string `json:"power ac,omitempty"`
}

var Client mqtt.Client

func InitMQTT() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(os.Getenv("MQTT_BROKER"))
	opts.SetClientID("go_mqtt_client")
	opts.OnConnect = func(c mqtt.Client) {
		fmt.Println("Connected to MQTT broker")
	}
	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		fmt.Printf("Connection lost: %v\n", err)
	}

	Client = mqtt.NewClient(opts)
	if token := Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func Subscribe(topic string) {
	Client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received from %s: %s\n", msg.Topic(), msg.Payload())

		var sensorList []SensorData
		err := json.Unmarshal(msg.Payload(), &sensorList)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

		for _, sensor := range sensorList {
			power := sensor.Power
			if power == "" {
				power = sensor.PowerAC
			}
			db.InsertSensor(sensor.Name, power, sensor.Voltage, sensor.Current)
			fmt.Printf("Saved: %s\n", sensor.Name)
		}
	})
}
