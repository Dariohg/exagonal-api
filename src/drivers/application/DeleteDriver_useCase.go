package application

import "f1-hex-api/src/drivers/domain"

type DeleteDriver struct {
	db domain.IDriver
}

func NewDeleteDriver(db domain.IDriver) *DeleteDriver {
	return &DeleteDriver{db: db}
}

func (dd *DeleteDriver) Execute(id int) error {
	return dd.db.Eliminar(id)
}
