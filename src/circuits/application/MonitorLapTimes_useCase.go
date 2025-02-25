package application

import (
	"f1-hex-api/src/circuits/domain"
	"f1-hex-api/src/circuits/domain/entities"
)

type MonitorLapTimes struct {
	db domain.ICircuit
}

func NewMonitorLapTimes(db domain.ICircuit) *MonitorLapTimes {
	return &MonitorLapTimes{db: db}
}

func (mlt *MonitorLapTimes) Execute(circuitoID int) ([]entities.LapTime, error) {
	return mlt.db.ObtenerTiemposVuelta(circuitoID)
}
