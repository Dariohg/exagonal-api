package application

import (
	"f1-hex-api/src/drivers/domain"
	"f1-hex-api/src/drivers/domain/entities"
)

type CreateDriver struct {
	db domain.IDriver
}

func NewCreateDriver(db domain.IDriver) *CreateDriver {
	return &CreateDriver{db: db}
}

func (cd *CreateDriver) Execute(driver *entities.Driver) error {
	return cd.db.Guardar(driver)
}
