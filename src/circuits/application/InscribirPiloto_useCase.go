package application

import (
	"f1-hex-api/src/circuits/domain"
)

type InscribirPiloto struct {
	db domain.ICircuit
}

func NewInscribirPiloto(db domain.ICircuit) *InscribirPiloto {
	return &InscribirPiloto{db: db}
}

func (ip *InscribirPiloto) Execute(circuitoID int, conductorID int) error {
	return ip.db.InscribirPiloto(circuitoID, conductorID)
}
