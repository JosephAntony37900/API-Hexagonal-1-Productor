package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/domain/entities"
	"github.com/streadway/amqp"
)

type OrderRepoMySQL struct {
	db      *sql.DB
	channel *amqp.Channel // Agregar el canal de RabbitMQ como campo
}

// Constructor del repositorio
func NewOrderRepoMySQL(db *sql.DB, channel *amqp.Channel) *OrderRepoMySQL {
	return &OrderRepoMySQL{db: db, channel: channel}
}

func (r *OrderRepoMySQL) Save(order entities.Order) error {
	query := `
		INSERT INTO Pedidos (Usuario_id, Producto, Estado, Pais, Entidad_federativa, Cp)
		VALUES (?, ?, 'Pendiente', ?, ?, ?)
	`

	_, err := r.db.Exec(query, order.Usuario_id, order.Producto, order.Pais, order.Entidad_federativa, order.Cp)
	if err != nil {
		return fmt.Errorf("error al guardar el pedido en la BD: %w", err)
	}

	//  Publicar evento en la cola "order.created" para que la API Consumidora lo procese
	err = r.PublishOrderCreated(order)
	if err != nil {
		return fmt.Errorf("error al publicar evento en la cola: %w", err)
	}

	return nil
}

// Método para obtener los pedidos de un usuario por su ID
func (r *OrderRepoMySQL) FindByUserID(usuarioID int) ([]entities.Order, error) {
	query := `
		SELECT Producto, Estado, Pais 
		FROM Pedidos 
		WHERE Usuario_id = ?
	`

	rows, err := r.db.Query(query, usuarioID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener pedidos del usuario %d: %w", usuarioID, err)
	}
	defer rows.Close()

	var orders []entities.Order
	for rows.Next() {
		var order entities.Order
		if err := rows.Scan(&order.Producto, &order.Estado, &order.Pais); err != nil {
			return nil, fmt.Errorf("error al escanear pedidos: %w", err)
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error en la iteración de pedidos: %w", err)
	}

	return orders, nil
}

// Método para eliminar un pedido por su ID
func (r *OrderRepoMySQL) Delete(orderID int) error {
	query := `DELETE FROM Pedidos WHERE id = ?`

	result, err := r.db.Exec(query, orderID)
	if err != nil {
		return fmt.Errorf("error al eliminar el pedido %d: %w", orderID, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró el pedido con ID %d", orderID)
	}

	fmt.Printf(" Pedido %d eliminado correctamente\n", orderID)
	return nil
}

// Método para actualizar un pedido
func (r *OrderRepoMySQL) Update(order entities.Order) error {
    query := `UPDATE Pedidos SET Usuario_id = ?, Producto = ?, Estado = ?, Pais = ?, Entidad_federativa = ?, Cp = ? WHERE id = ?`

    result, err := r.db.Exec(query, order.Usuario_id, order.Producto, order.Estado, order.Pais, order.Entidad_federativa, order.Cp, order.Id)
    if err != nil {
        return fmt.Errorf("error al actualizar el pedido %d: %w", order.Id, err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error al obtener filas afectadas: %w", err)
    }

    if rowsAffected == 0 {
        return fmt.Errorf("no se encontró el pedido con ID %d", order.Id)
    }

    fmt.Printf(" Pedido %d actualizado correctamente\n", order.Id)
    return nil
}


// Método para encontrar un pedido por su ID
func (r *OrderRepoMySQL) FindByID(orderID int) (*entities.Order, error) {
    query := `SELECT id, Usuario_id, Producto, Estado, Pais, Entidad_federativa, Cp FROM Pedidos WHERE id = ?`

    row := r.db.QueryRow(query, orderID)
    var order entities.Order
    if err := row.Scan(&order.Id, &order.Usuario_id, &order.Producto, &order.Estado, &order.Pais, &order.Entidad_federativa, &order.Cp); err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("pedido no encontrado con ID %d", orderID)
        }
        return nil, fmt.Errorf("error al obtener el pedido con ID %d: %w", orderID, err)
    }

    return &order, nil
}

//publicar event
func (r *OrderRepoMySQL) PublishOrderCreated(order entities.Order) error {
	// Convertir la orden a JSON
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("error al convertir la orden a JSON: %w", err)
	}

	// Publicar el mensaje en la cola "created.order"
	err = r.channel.Publish(
		"",               // exchange
		"created.order",  // queue name
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        orderJSON,
		},
	)
	if err != nil {
		return fmt.Errorf("error al publicar el mensaje en RabbitMQ: %w", err)
	}

	fmt.Printf(" Evento 'order.created' publicado para el pedido %d\n", order.Id)
	return nil
}
