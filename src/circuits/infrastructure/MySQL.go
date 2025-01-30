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
