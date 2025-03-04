package infrastructure

import (
	"f1-hex-api/src/circuits/domain/entities"
	"f1-hex-api/src/core"
	"fmt"
	"log"
	"time"
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
              (nombre, pais, longitud, numero_vueltas, numero_curvas, tiempo_promedio_vuelta) 
              VALUES (?, ?, ?, ?, ?, ?)`

	result, err := mysql.conn.ExecutePreparedQuery(query,
		circuit.Nombre,
		circuit.Pais,
		circuit.Longitud,
		circuit.NumeroVueltas,
		circuit.NumeroCurvas,
		circuit.TiempoPromedioVuelta)

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
              numero_curvas, tiempo_promedio_vuelta, fecha_creacion FROM circuitos`
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
			&circuit.TiempoPromedioVuelta,
			&circuit.FechaCreacion); err != nil {
			return nil, fmt.Errorf("error al escanear circuito: %v", err)
		}
		circuits = append(circuits, circuit)
	}

	return circuits, nil
}

func (mysql *MySQL) ObtenerPorId(id int) (*entities.Circuit, error) {
	query := `SELECT id, nombre, pais, longitud, numero_vueltas, 
              numero_curvas, tiempo_promedio_vuelta, fecha_creacion 
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
			&circuit.TiempoPromedioVuelta,
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
                  numero_curvas = ?,
                  tiempo_promedio_vuelta = ?
              WHERE id = ?`

	result, err := mysql.conn.ExecutePreparedQuery(query,
		circuit.Nombre,
		circuit.Pais,
		circuit.Longitud,
		circuit.NumeroVueltas,
		circuit.NumeroCurvas,
		circuit.TiempoPromedioVuelta,
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
	// Esta consulta optimizada obtiene solo el tiempo más reciente para cada piloto
	query := `
    SELECT t1.id, t1.circuito_id, t1.conductor_id, t1.numero_vuelta, t1.tiempo, t1.timestamp,
           c.nombre_completo as nombre_piloto, c.nombre_equipo, c.numero_carro
    FROM tiempos_vuelta t1
    INNER JOIN (
        SELECT conductor_id, MAX(numero_vuelta) as ultima_vuelta
        FROM tiempos_vuelta
        WHERE circuito_id = ?
        GROUP BY conductor_id
    ) t2 ON t1.conductor_id = t2.conductor_id AND t1.numero_vuelta = t2.ultima_vuelta
    INNER JOIN conductores c ON t1.conductor_id = c.id
    WHERE t1.circuito_id = ?
    ORDER BY t1.tiempo ASC`

	rows := mysql.conn.FetchRows(query, circuitoID, circuitoID)
	defer rows.Close()

	var lapTimes []entities.LapTime
	for rows.Next() {
		var lapTime entities.LapTime
		var nombrePiloto, nombreEquipo string
		var numeroCarro int

		if err := rows.Scan(
			&lapTime.ID,
			&lapTime.CircuitoID,
			&lapTime.ConductorID,
			&lapTime.NumeroVuelta,
			&lapTime.Tiempo,
			&lapTime.Timestamp,
			&nombrePiloto,
			&nombreEquipo,
			&numeroCarro); err != nil {
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

// Obtener incidentes activos con ID mayor al último conocido
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

// Guardar un nuevo incidente
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

func (mysql *MySQL) ObtenerUltimoRecord(circuitoID int) (*entities.LapRecord, error) {
	// Obtener el tiempo promedio del circuito
	queryCircuito := `SELECT tiempo_promedio_vuelta FROM circuitos WHERE id = ?`
	rowsCircuito := mysql.conn.FetchRows(queryCircuito, circuitoID)
	defer rowsCircuito.Close()

	var tiempoPromedio float64
	if !rowsCircuito.Next() {
		return nil, fmt.Errorf("circuito no encontrado")
	}

	if err := rowsCircuito.Scan(&tiempoPromedio); err != nil {
		return nil, fmt.Errorf("error al leer tiempo promedio: %v", err)
	}

	// Consultar el mejor tiempo de vuelta actual en este circuito
	// Solo consideramos tiempos de vuelta de los últimos 10 minutos
	tiempoLimite := time.Now().Add(-10 * time.Minute)

	query := `
    SELECT t.id, t.circuito_id, t.conductor_id, c.nombre_completo, t.tiempo, t.timestamp
    FROM tiempos_vuelta t
    INNER JOIN conductores c ON t.conductor_id = c.id
    WHERE t.circuito_id = ? AND t.timestamp > ?
    ORDER BY t.tiempo ASC 
    LIMIT 1`

	rows := mysql.conn.FetchRows(query, circuitoID, tiempoLimite)
	defer rows.Close()

	if !rows.Next() {
		return nil, nil // No hay tiempos de vuelta recientes
	}

	var recordID, conductorID int
	var circuitID int
	var nombrePiloto string
	var tiempoVuelta float64
	var timestamp time.Time

	if err := rows.Scan(&recordID, &circuitID, &conductorID, &nombrePiloto, &tiempoVuelta, &timestamp); err != nil {
		return nil, fmt.Errorf("error al escanear tiempo de vuelta: %v", err)
	}

	// Verificar si es un récord (mejor que el tiempo promedio)
	if tiempoVuelta < tiempoPromedio {
		diferencia := tiempoPromedio - tiempoVuelta

		record := &entities.LapRecord{
			ID:               recordID,
			CircuitoID:       circuitID,
			ConductorID:      conductorID,
			NombrePiloto:     nombrePiloto,
			TiempoVuelta:     tiempoVuelta,
			DiferenciaTiempo: diferencia,
			Timestamp:        timestamp,
		}

		return record, nil
	}

	return nil, nil // No hay récord
}
