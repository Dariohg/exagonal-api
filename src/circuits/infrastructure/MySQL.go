package infrastructure

import (
	"f1-hex-api/src/circuits/domain/entities"
	"f1-hex-api/src/core"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Guardar(circuit *entities.Circuit) error {
	query := `INSERT INTO circuitos 
              (nombre, pais, longitud, numero_vueltas, numero_curvas) 
              VALUES (?, ?, ?, ?, ?)`

	result, err := mysql.conn.ExecutePreparedQuery(query,
		circuit.Nombre,
		circuit.Pais,
		circuit.Longitud,
		circuit.NumeroVueltas,
		circuit.NumeroCurvas)

	if err != nil {
		return fmt.Errorf("error al guardar el circuito: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener el ID del circuito insertado: %v", err)
	}

	circuit.ID = int(id)
	return nil
}

func (mysql *MySQL) ObtenerTodos() ([]entities.Circuit, error) {
	query := `SELECT id, nombre, pais, longitud, numero_vueltas, 
              numero_curvas, fecha_creacion FROM circuitos`
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var circuits []entities.Circuit
	for rows.Next() {
		var circuit entities.Circuit
		if err := rows.Scan(
			&circuit.ID,
			&circuit.Nombre,
			&circuit.Pais,
			&circuit.Longitud,
			&circuit.NumeroVueltas,
			&circuit.NumeroCurvas,
			&circuit.FechaCreacion); err != nil {
			return nil, fmt.Errorf("error al escanear circuito: %v", err)
		}
		circuits = append(circuits, circuit)
	}

	return circuits, nil
}

func (mysql *MySQL) ObtenerPorId(id int) (*entities.Circuit, error) {
	query := `SELECT id, nombre, pais, longitud, numero_vueltas, 
              numero_curvas, fecha_creacion 
              FROM circuitos WHERE id = ?`

	rows := mysql.conn.FetchRows(query, id)
	defer rows.Close()

	var circuit entities.Circuit
	if rows.Next() {
		err := rows.Scan(
			&circuit.ID,
			&circuit.Nombre,
			&circuit.Pais,
			&circuit.Longitud,
			&circuit.NumeroVueltas,
			&circuit.NumeroCurvas,
			&circuit.FechaCreacion)
		if err != nil {
			return nil, fmt.Errorf("error al escanear circuito: %v", err)
		}
		return &circuit, nil
	}
	return nil, fmt.Errorf("circuito no encontrado")
}

func (mysql *MySQL) Actualizar(circuit *entities.Circuit) error {
	query := `UPDATE circuitos 
              SET nombre = ?, 
                  pais = ?, 
                  longitud = ?, 
                  numero_vueltas = ?, 
                  numero_curvas = ?
              WHERE id = ?`

	result, err := mysql.conn.ExecutePreparedQuery(query,
		circuit.Nombre,
		circuit.Pais,
		circuit.Longitud,
		circuit.NumeroVueltas,
		circuit.NumeroCurvas,
		circuit.ID)

	if err != nil {
		return fmt.Errorf("error al actualizar el circuito: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró el circuito para actualizar")
	}

	return nil
}

func (mysql *MySQL) Eliminar(id int) error {
	query := "DELETE FROM circuitos WHERE id = ?"

	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar el circuito: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró el circuito para eliminar")
	}

	return nil
}

func (mysql *MySQL) InscribirPiloto(circuitoID int, conductorID int) error {
	// Primero verificar que el circuito existe
	checkCircuito := "SELECT id FROM circuitos WHERE id = ?"
	rowsCircuito := mysql.conn.FetchRows(checkCircuito, circuitoID)
	if !rowsCircuito.Next() {
		return fmt.Errorf("el circuito con ID %d no existe", circuitoID)
	}
	rowsCircuito.Close()

	// Verificar que el conductor existe
	checkConductor := "SELECT id FROM conductores WHERE id = ?"
	rowsConductor := mysql.conn.FetchRows(checkConductor, conductorID)
	if !rowsConductor.Next() {
		return fmt.Errorf("el conductor con ID %d no existe", conductorID)
	}
	rowsConductor.Close()

	// Verificar que la combinación no existe
	checkCombinacion := "SELECT circuito_id FROM circuito_conductor WHERE circuito_id = ? AND conductor_id = ?"
	rowsCombinacion := mysql.conn.FetchRows(checkCombinacion, circuitoID, conductorID)
	if rowsCombinacion.Next() {
		return fmt.Errorf("el conductor ya está inscrito en este circuito")
	}
	rowsCombinacion.Close()

	// Si todo está bien, hacer la inserción
	query := `INSERT INTO circuito_conductor (circuito_id, conductor_id) 
              VALUES (?, ?)`

	_, err := mysql.conn.ExecutePreparedQuery(query, circuitoID, conductorID)
	if err != nil {
		return fmt.Errorf("error al inscribir piloto: %v", err)
	}

	return nil
}

func (mysql *MySQL) ObtenerPilotosInscritos(circuitoID int) ([]entities.CircuitDriver, error) {
	query := `SELECT cc.circuito_id, cc.conductor_id, c.nombre_completo, c.nombre_equipo, c.numero_carro
              FROM circuito_conductor cc
              INNER JOIN conductores c ON cc.conductor_id = c.id
              WHERE cc.circuito_id = ?`

	rows := mysql.conn.FetchRows(query, circuitoID)
	defer rows.Close()

	var drivers []entities.CircuitDriver
	for rows.Next() {
		var driver entities.CircuitDriver
		if err := rows.Scan(
			&driver.CircuitoID,
			&driver.ConductorID,
			&driver.NombreConductor,
			&driver.NombreEquipo,
			&driver.NumeroCarro); err != nil {
			return nil, fmt.Errorf("error al escanear conductor: %v", err)
		}
		drivers = append(drivers, driver)
	}

	return drivers, nil
}

func (mysql *MySQL) ObtenerTiemposVuelta(circuitoID int) ([]entities.LapTime, error) {
	// Esta consulta obtiene el tiempo más reciente para cada piloto
	query := `
    SELECT t1.id, t1.circuito_id, t1.conductor_id, t1.numero_vuelta, t1.tiempo, t1.timestamp 
    FROM tiempos_vuelta t1
    INNER JOIN (
        SELECT conductor_id, MAX(timestamp) as max_timestamp
        FROM tiempos_vuelta
        WHERE circuito_id = ?
        GROUP BY conductor_id
    ) t2 ON t1.conductor_id = t2.conductor_id AND t1.timestamp = t2.max_timestamp
    WHERE t1.circuito_id = ?
    ORDER BY t1.tiempo ASC` // Ordenamos por el mejor tiempo

	rows := mysql.conn.FetchRows(query, circuitoID, circuitoID)
	defer rows.Close()

	var lapTimes []entities.LapTime
	for rows.Next() {
		var lapTime entities.LapTime
		if err := rows.Scan(
			&lapTime.ID,
			&lapTime.CircuitoID,
			&lapTime.ConductorID,
			&lapTime.NumeroVuelta,
			&lapTime.Tiempo,
			&lapTime.Timestamp); err != nil {
			return nil, fmt.Errorf("error al escanear tiempo de vuelta: %v", err)
		}
		lapTimes = append(lapTimes, lapTime)
	}

	return lapTimes, nil
}

func (mysql *MySQL) GuardarTiempoVuelta(lapTime *entities.LapTime) error {
	query := `INSERT INTO tiempos_vuelta 
              (circuito_id, conductor_id, numero_vuelta, tiempo) 
              VALUES (?, ?, ?, ?)`

	result, err := mysql.conn.ExecutePreparedQuery(query,
		lapTime.CircuitoID,
		lapTime.ConductorID,
		lapTime.NumeroVuelta,
		lapTime.Tiempo)

	if err != nil {
		return fmt.Errorf("error al guardar tiempo de vuelta: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener ID del tiempo de vuelta: %v", err)
	}

	lapTime.ID = int(id)
	return nil
}

func (mysql *MySQL) ObtenerPosiciones(circuitoID int) ([]entities.Position, error) {
	query := `
    SELECT p1.id, p1.circuito_id, p1.conductor_id, p1.posicion, p1.timestamp 
    FROM posiciones_carrera p1
    INNER JOIN (
        SELECT conductor_id, MAX(timestamp) as max_timestamp
        FROM posiciones_carrera
        WHERE circuito_id = ?
        GROUP BY conductor_id
    ) p2 ON p1.conductor_id = p2.conductor_id AND p1.timestamp = p2.max_timestamp
    WHERE p1.circuito_id = ?
    ORDER BY p1.posicion ASC`

	rows := mysql.conn.FetchRows(query, circuitoID, circuitoID)
	defer rows.Close()

	var positions []entities.Position
	for rows.Next() {
		var position entities.Position
		if err := rows.Scan(
			&position.ID,
			&position.CircuitoID,
			&position.ConductorID,
			&position.Posicion,
			&position.Timestamp); err != nil {
			return nil, fmt.Errorf("error al escanear posición: %v", err)
		}
		positions = append(positions, position)
	}

	return positions, nil
}

func (mysql *MySQL) GuardarPosicion(position *entities.Position) error {
	query := `INSERT INTO posiciones_carrera 
              (circuito_id, conductor_id, posicion) 
              VALUES (?, ?, ?)`

	result, err := mysql.conn.ExecutePreparedQuery(query,
		position.CircuitoID,
		position.ConductorID,
		position.Posicion)

	if err != nil {
		return fmt.Errorf("error al guardar posición: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener ID de la posición: %v", err)
	}

	position.ID = int(id)
	return nil
}

func (mysql *MySQL) ObtenerIncidentesActivos(circuitoID int, ultimoID int) ([]entities.Incident, error) {
	query := `SELECT id, circuito_id, tipo_incidente, descripcion, conductor_id, 
              estado, timestamp 
              FROM incidentes_pista 
              WHERE circuito_id = ? AND id > ? AND estado = 'ACTIVO'
              ORDER BY timestamp DESC`

	rows := mysql.conn.FetchRows(query, circuitoID, ultimoID)
	defer rows.Close()

	var incidents []entities.Incident
	for rows.Next() {
		var incident entities.Incident
		if err := rows.Scan(
			&incident.ID,
			&incident.CircuitoID,
			&incident.TipoIncidente,
			&incident.Descripcion,
			&incident.ConductorID,
			&incident.Estado,
			&incident.Timestamp); err != nil {
			return nil, fmt.Errorf("error al escanear incidente: %v", err)
		}
		incidents = append(incidents, incident)
	}

	return incidents, nil
}

func (mysql *MySQL) GuardarIncidente(incident *entities.Incident) error {
	query := `INSERT INTO incidentes_pista 
              (circuito_id, tipo_incidente, descripcion, conductor_id, estado) 
              VALUES (?, ?, ?, ?, ?)`

	result, err := mysql.conn.ExecutePreparedQuery(query,
		incident.CircuitoID,
		incident.TipoIncidente,
		incident.Descripcion,
		incident.ConductorID,
		incident.Estado)

	if err != nil {
		return fmt.Errorf("error al guardar incidente: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener ID del incidente: %v", err)
	}

	incident.ID = int(id)
	return nil
}
