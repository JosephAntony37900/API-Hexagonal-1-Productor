package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JosephAntony37900/API-Hexagonal-1-Productor/Order/application"
)

type CreateOrderRequest struct {
	Usuario_id        int    `json:"usuario_id"`
	Producto          string `json:"producto"`
	Pais             string `json:"pais"`
	Entidad_federativa string `json:"entidad_federativa"`
	Cp               string `json:"cp"`
}

type CreateOrderController struct {
	useCase *application.CreateOrder
}

func NewCreateOrderController(useCase *application.CreateOrder) *CreateOrderController {
	return &CreateOrderController{useCase: useCase}
}

func (c *CreateOrderController) Handle(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest

	// Decodificar el JSON de la solicitud
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error en el formato de la solicitud", http.StatusBadRequest)
		return
	}

	err := c.useCase.Run(req.Usuario_id, req.Producto, req.Pais, req.Entidad_federativa, req.Cp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Pedido creado con éxito y enviado a la cola 'order.created'"}`))
}