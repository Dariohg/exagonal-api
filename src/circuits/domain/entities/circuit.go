package entities

import "time"

type Circuit struct {
	ID            int       `json:"id"`
	Nombre        string    `json:"nombre"`
	Pais          string    `json:"pais"`
	Longitud      float64   `json:"longitud"`
	NumeroVueltas int       `json:"numero_vueltas"`
	NumeroCurvas  int       `json:"numero_curvas"`
	FechaCreacion time.Time `json:"fecha_creacion,omitempty"`
}
