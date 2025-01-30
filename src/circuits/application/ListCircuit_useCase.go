package application

import (
	"f1-hex-api/src/circuits/domain"
	"f1-hex-api/src/circuits/domain/entities"
)

type ListCircuit struct {
	db domain.ICircuit
}

func NewListCircuit(db domain.ICircuit) *ListCircuit {
	return &ListCircuit{db: db}
}

func (lc *ListCircuit) Execute() ([]entities.Circuit, error) {
	return lc.db.ObtenerTodos()
}
