package adapters

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/streadway/amqp"
	"github.com/joho/godotenv"
)

var conn *amqp.Connection
var channel *amqp.Channel

// Inicializa la conexión a RabbitMQ
func InitRabbitMQ() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}

	username := os.Getenv("Name")
	password := os.Getenv("PasswordQueue")

	rabbitURL := fmt.Sprintf("amqp://%s:%s@98.85.106.157:5672/", username, password)
	conn, err = amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %s", err)
	}

	channel, err = conn.Channel()
	if err != nil {
		log.Fatalf("Error al abrir un canal en RabbitMQ: %s", err)
	}

	// Declarar la cola "created.order"
	_, err = channel.QueueDeclare(
		"created.order", // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		log.Fatalf("Error al declarar la cola 'created.order': %s", err)
	}

	log.Println("Conectado a RabbitMQ exitosamente")
}

// Cierra la conexión y el canal
func CloseRabbitMQ() {
	if channel != nil {
		channel.Close()
	}
	if conn != nil {
		conn.Close()
	}
}

// Obtener el canal de RabbitMQ
func GetChannel() *amqp.Channel {
	return channel
}