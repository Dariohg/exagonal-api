package domain

import "f1-hex-api/src/drivers/domain/entities"

type IDriver interface {
	Guardar(*entities.Driver) error
	ObtenerTodos() ([]entities.Driver, error)
	ObtenerPorId(id int) (*entities.Driver, error)
	Actualizar(*entities.Driver) error
	Eliminar(id int) error
}
