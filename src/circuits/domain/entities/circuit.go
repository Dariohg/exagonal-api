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

func NewCircuit(nombre string, pais string, longitud float64, numeroVueltas int, numeroCurvas int) *Circuit {
	return &Circuit{
		Nombre:        nombre,
		Pais:          pais,
		Longitud:      longitud,
		NumeroVueltas: numeroVueltas,
		NumeroCurvas:  numeroCurvas,
	}
}

func (c *Circuit) GetNombre() string {
	return c.Nombre
}

func (c *Circuit) SetNombre(nombre string) {
	c.Nombre = nombre
}
