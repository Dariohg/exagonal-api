package application

import (
	"f1-hex-api/src/circuits/domain"
	"f1-hex-api/src/circuits/domain/entities"
)

type GetDriversByCircuit struct {
	db domain.ICircuit
}

func NewGetDriversByCircuit(db domain.ICircuit) *GetDriversByCircuit {
	return &GetDriversByCircuit{db: db}
}

func (gdc *GetDriversByCircuit) Execute(circuitoID int) ([]entities.CircuitDriver, error) {
	return gdc.db.ObtenerPilotosInscritos(circuitoID)
}
