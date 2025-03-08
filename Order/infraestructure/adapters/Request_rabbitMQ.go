package adapters

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/streadway/amqp"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/domain/entities"
	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/domain/repository"
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

	// Declarar las colas necesarias
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

	_, err = channel.QueueDeclare(
		"order.confirmed", // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		log.Fatalf("Error al declarar la cola 'order.confirmed': %s", err)
	}

	_, err = channel.QueueDeclare(
		"order.rejected", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Fatalf("Error al declarar la cola 'order.rejected': %s", err)
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

// ConsumeConfirmedOrders consume mensajes de la cola "order.confirmed"
func ConsumeConfirmedOrders(repo repository.OrderRepository) {
	// Consumir mensajes de la cola "order.confirmed"
	msgs, err := channel.Consume(
		"order.confirmed", // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		log.Fatalf("Error al registrar el consumidor para 'order.confirmed': %s", err)
	}

	log.Println("✅ Consumidor de 'order.confirmed' iniciado correctamente")

	go func() {
		for d := range msgs {
			var order entities.Order
			if err := json.Unmarshal(d.Body, &order); err != nil {
				log.Printf("Error al decodificar el mensaje de 'order.confirmed': %s", err)
				continue
			}

			log.Printf("Mensaje recibido en 'order.confirmed': %+v", order)

			// Actualizar el estado del pedido a "Enviado"
			order.Estado = "Enviado"
			if err := repo.Update(order); err != nil {
				log.Printf("Error al actualizar el pedido %d: %s", order.Id, err)
				continue
			}

			log.Printf("✅ Pedido %d confirmado y actualizado a 'Enviado'", order.Id)
		}
	}()
}

// ConsumeRejectedOrders consume mensajes de la cola "order.rejected"
func ConsumeRejectedOrders(repo repository.OrderRepository) {
	// Consumir mensajes de la cola "order.rejected"
	msgs, err := channel.Consume(
		"order.rejected", // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		log.Fatalf("Error al registrar el consumidor para 'order.rejected': %s", err)
	}

	log.Println("✅ Consumidor de 'order.rejected' iniciado correctamente")

	go func() {
		for d := range msgs {
			var order entities.Order
			if err := json.Unmarshal(d.Body, &order); err != nil {
				log.Printf("Error al decodificar el mensaje de 'order.rejected': %s", err)
				continue
			}

			log.Printf("Mensaje recibido en 'order.rejected': %+v", order)

			// Actualizar el estado del pedido a "Fallido"
			order.Estado = "Fallido"
			if err := repo.Update(order); err != nil {
				log.Printf("Error al actualizar el pedido %d: %s", order.Id, err)
				continue
			}

			log.Printf("❌ Pedido %d rechazado y actualizado a 'Fallido'", order.Id)
		}
	}()
}