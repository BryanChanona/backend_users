package adapters

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/BryanChanona/backend_users/src/Mqtt/domain"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
)

var client mqtt.Client

// InitMQTT inicializa la conexión al broker MQTT
func InitMQTT() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://3.234.181.19:1883") // Dirección del broker MQTT
	opts.SetClientID("mqtt_publisher")
	opts.SetUsername("carlos")
	opts.SetPassword("carlos")

	// Crear cliente MQTT
	client = mqtt.NewClient(opts)

	// Intentar conectar
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error al conectar con MQTT: %v", token.Error())
	}
	fmt.Println("Conectado al broker MQTT")
}

// PublishHandler maneja la solicitud POST y publica en MQTT
func PublishHandler(c *gin.Context) {
	// Verifica si el cliente MQTT está conectado
	if client == nil || !client.IsConnected() {
		log.Println("El cliente MQTT no está conectado.")
		c.JSON(500, gin.H{"error": "Cliente MQTT no conectado"})
		return
	}

	// Leer los datos JSON de la solicitud
	var data domain.DeviceData
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Printf("Error al leer los datos: %v", err)
		c.JSON(400, gin.H{"error": "Error al leer los datos"})
		return
	}

	// Convertir los datos a JSON
	payload, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error al serializar los datos: %v", err)
		c.JSON(500, gin.H{"error": "Error al serializar los datos"})
		return
	}

	// Publicar en MQTT
	topic := "device.data"
	token := client.Publish(topic, 0, false, payload)
	token.Wait() // Esperar a que se complete la publicación

	if token.Error() != nil {
		log.Printf("Error al publicar el mensaje: %v", token.Error())
		c.JSON(500, gin.H{"error": "Error al publicar el mensaje en MQTT"})
		return
	}

	fmt.Println("Mensaje enviado a MQTT:", string(payload))

	// Responder con éxito
	c.JSON(200, gin.H{"message": "Datos enviados correctamente"})
}
