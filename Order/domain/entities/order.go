package entities

type Order struct {
	Id int
	Usuario_id int
	Producto string
	Estado string //estado del evento
	Pais string
	Entidad_federativa string
	Cp string
}