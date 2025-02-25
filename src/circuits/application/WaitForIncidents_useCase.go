package application

import (
	"f1-hex-api/src/circuits/domain"
	"f1-hex-api/src/circuits/domain/entities"
)

type WaitForIncidents struct {
	db domain.ICircuit
}

func NewWaitForIncidents(db domain.ICircuit) *WaitForIncidents {
	return &WaitForIncidents{db: db}
}

func (wfi *WaitForIncidents) Execute(circuitoID int, ultimoID int) ([]entities.Incident, error) {
	return wfi.db.ObtenerIncidentesActivos(circuitoID, ultimoID)
}
