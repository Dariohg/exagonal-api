package entities

import "time"

type Driver struct {
	ID             int       `json:"id"`
	NombreCompleto string    `json:"nombre_completo"`
	Nacionalidad   string    `json:"nacionalidad"`
	NombreEquipo   string    `json:"nombre_equipo"`
	NumeroCarro    int       `json:"numero_carro"`
	Edad           int       `json:"edad"`
	FechaCreacion  time.Time `json:"fecha_creacion,omitempty"`
}
