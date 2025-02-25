package application

import (
	"f1-hex-api/src/circuits/domain"
)

type DeleteCircuit struct {
	db domain.ICircuit
}

func NewDeleteCircuit(db domain.ICircuit) *DeleteCircuit {
	return &DeleteCircuit{db: db}
}

func (dc *DeleteCircuit) Execute(id int) error {
	return dc.db.Eliminar(id)
}
