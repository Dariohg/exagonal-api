package application

import (
	"f1-hex-api/src/circuits/domain"
	"f1-hex-api/src/circuits/domain/entities"
)

type CreateCircuit struct {
	db domain.ICircuit
}

func NewCreateCircuit(db domain.ICircuit) *CreateCircuit {
	return &CreateCircuit{db: db}
}

func (cc *CreateCircuit) Execute(circuit *entities.Circuit) error {
	return cc.db.Guardar(circuit)
}
