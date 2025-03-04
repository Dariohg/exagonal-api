package entities

import "time"

type PilotoInfo struct {
	Nombre      string `json:"nombre"`
	Equipo      string `json:"equipo"`
	NumeroCarro int    `json:"numero_carro"`
}

type LapTime struct {
	ID           int        `json:"id"`
	CircuitoID   int        `json:"circuito_id"`
	ConductorID  int        `json:"conductor_id"`
	NumeroVuelta int        `json:"numero_vuelta"`
	Tiempo       float64    `json:"tiempo"`
	Timestamp    time.Time  `json:"timestamp"`
	DatosPiloto  PilotoInfo `json:"datos_piloto,omitempty"`
}
