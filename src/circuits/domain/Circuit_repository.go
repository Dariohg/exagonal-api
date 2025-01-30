package domain

import "f1-hex-api/src/circuits/domain/entities"

type ICircuit interface {
	Guardar(*entities.Circuit) error
	ObtenerTodos() ([]entities.Circuit, error)
	ObtenerPorId(id int) (*entities.Circuit, error)
	Actualizar(*entities.Circuit) error
	Eliminar(id int) error
}
