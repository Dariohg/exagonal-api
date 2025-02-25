package entities

import "time"

type Incident struct {
	ID            int       `json:"id"`
	CircuitoID    int       `json:"circuito_id"`
	TipoIncidente string    `json:"tipo_incidente"`
	Descripcion   string    `json:"descripcion,omitempty"`
	ConductorID   *int      `json:"conductor_id,omitempty"`
	Estado        string    `json:"estado"`
	Timestamp     time.Time `json:"timestamp"`
}
