package application

import (
	"f1-hex-api/src/circuits/domain"
	"f1-hex-api/src/circuits/domain/entities"
)

type MonitorLapRecords struct {
	db domain.ICircuit
}

func NewMonitorLapRecords(db domain.ICircuit) *MonitorLapRecords {
	return &MonitorLapRecords{db: db}
}

func (mlr *MonitorLapRecords) Execute(circuitoID int) (*entities.LapRecord, error) {
	return mlr.db.ObtenerUltimoRecord(circuitoID)
}
