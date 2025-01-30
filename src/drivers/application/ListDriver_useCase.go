package application

import (
	"f1-hex-api/src/drivers/domain"
	"f1-hex-api/src/drivers/domain/entities"
)

type ListDriver struct {
	db domain.IDriver
}

func NewListDriver(db domain.IDriver) *ListDriver {
	return &ListDriver{db: db}
}

func (ld *ListDriver) Execute() ([]entities.Driver, error) {
	return ld.db.ObtenerTodos()
}
