package application

import (
	"f1-hex-api/src/circuits/domain"
	"f1-hex-api/src/circuits/domain/entities"
)

type SavePosition struct {
	db domain.ICircuit
}

func NewSavePosition(db domain.ICircuit) *SavePosition {
	return &SavePosition{db: db}
}

func (sp *SavePosition) Execute(position *entities.Position) error {
	return sp.db.GuardarPosicion(position)
}
