package entities

import "time"

type Position struct {
	ID          int       `json:"id"`
	CircuitoID  int       `json:"circuito_id"`
	ConductorID int       `json:"conductor_id"`
	Posicion    int       `json:"posicion"`
	Timestamp   time.Time `json:"timestamp"`
}
