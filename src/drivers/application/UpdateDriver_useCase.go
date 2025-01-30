package application

import (
	"f1-hex-api/src/drivers/domain"
	"f1-hex-api/src/drivers/domain/entities"
)

type UpdateDriver struct {
	db domain.IDriver
}

func NewUpdateDriver(db domain.IDriver) *UpdateDriver {
	return &UpdateDriver{db: db}
}

func (ud *UpdateDriver) Execute(driver *entities.Driver) error {
	return ud.db.Actualizar(driver)
}
