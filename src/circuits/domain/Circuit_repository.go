package domain

import "f1-hex-api/src/circuits/domain/entities"

type ICircuit interface {
	Guardar(*entities.Circuit) error
	ObtenerTodos() ([]entities.Circuit, error)
	ObtenerPorId(id int) (*entities.Circuit, error)
	Actualizar(*entities.Circuit) error
	Eliminar(id int) error
	InscribirPiloto(circuitoID int, conductorID int) error
	ObtenerPilotosInscritos(circuitoID int) ([]entities.CircuitDriver, error)

	ObtenerTiemposVuelta(circuitoID int) ([]entities.LapTime, error)
	GuardarTiempoVuelta(lapTime *entities.LapTime) error

	ObtenerUltimoRecord(circuitoID int) (*entities.LapRecord, error)

	ObtenerIncidentesActivos(circuitoID int, ultimoID int) ([]entities.Incident, error)
	GuardarIncidente(incident *entities.Incident) error
}
