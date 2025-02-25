package application

import (
	"f1-hex-api/src/circuits/domain"
	"f1-hex-api/src/circuits/domain/entities"
)

type MonitorPositions struct {
	db domain.ICircuit
}

func NewMonitorPositions(db domain.ICircuit) *MonitorPositions {
	return &MonitorPositions{db: db}
}

func (mp *MonitorPositions) Execute(circuitoID int) ([]entities.Position, error) {
	return mp.db.ObtenerPosiciones(circuitoID)
}
