package entities

import "time"

type LapRecord struct {
	ID               int       `json:"id"`
	CircuitoID       int       `json:"circuito_id"`
	ConductorID      int       `json:"conductor_id"`
	NombrePiloto     string    `json:"nombre_piloto,omitempty"`
	TiempoVuelta     float64   `json:"tiempo_vuelta"`
	DiferenciaTiempo float64   `json:"diferencia_tiempo"`
	Timestamp        time.Time `json:"timestamp"`
}
