package application

import (
	"f1-hex-api/src/drivers/domain"
	"f1-hex-api/src/drivers/domain/entities"
)

type GetDriver struct {
	db domain.IDriver
}

func NewGetDriver(db domain.IDriver) *GetDriver {
	return &GetDriver{db: db}
}

func (gd *GetDriver) Execute(id int) (*entities.Driver, error) {
	return gd.db.ObtenerPorId(id)
}
