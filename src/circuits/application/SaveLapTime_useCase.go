package application

import (
	"f1-hex-api/src/circuits/domain"
	"f1-hex-api/src/circuits/domain/entities"
)

type SaveLapTime struct {
	db domain.ICircuit
}

func NewSaveLapTime(db domain.ICircuit) *SaveLapTime {
	return &SaveLapTime{db: db}
}

func (slt *SaveLapTime) Execute(lapTime *entities.LapTime) error {
	return slt.db.GuardarTiempoVuelta(lapTime)
}
