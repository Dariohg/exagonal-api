package application

import (
	"f1-hex-api/src/circuits/domain"
	"f1-hex-api/src/circuits/domain/entities"
)

type UpdateCircuit struct {
	db domain.ICircuit
}

func NewUpdateCircuit(db domain.ICircuit) *UpdateCircuit {
	return &UpdateCircuit{db: db}
}

func (uc *UpdateCircuit) Execute(circuit *entities.Circuit) error {
	return uc.db.Actualizar(circuit)
}
