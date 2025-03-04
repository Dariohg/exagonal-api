package application

import (
	"f1-hex-api/src/circuits/domain"
	"f1-hex-api/src/circuits/domain/entities"
)

type CreateIncident struct {
	db domain.ICircuit
}

func NewCreateIncident(db domain.ICircuit) *CreateIncident {
	return &CreateIncident{db: db}
}

func (ci *CreateIncident) Execute(incident *entities.Incident) error {
	return ci.db.GuardarIncidente(incident)
}
