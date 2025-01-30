package application

import (
	"f1-hex-api/src/circuits/domain"
	"f1-hex-api/src/circuits/domain/entities"
)

type GetCircuit struct {
	db domain.ICircuit
}

func NewGetCircuit(db domain.ICircuit) *GetCircuit {
	return &GetCircuit{db: db}
}

func (gc *GetCircuit) Execute(id int) (*entities.Circuit, error) {
	return gc.db.ObtenerPorId(id)
}
