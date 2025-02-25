package entities

type CircuitDriver struct {
	CircuitoID      int    `json:"circuito_id"`
	ConductorID     int    `json:"conductor_id"`
	NombreConductor string `json:"nombre_conductor,omitempty"`
	NombreEquipo    string `json:"nombre_equipo,omitempty"`
	NumeroCarro     int    `json:"numero_carro,omitempty"`
}
