package infrastructure

import (
	"f1-hex-api/src/core"
	"f1-hex-api/src/drivers/domain/entities"
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

func (mysql *MySQL) Guardar(driver *entities.Driver) error {
	query := `INSERT INTO conductores 
              (nombre_completo, nacionalidad, nombre_equipo, numero_carro, edad) 
              VALUES (?, ?, ?, ?, ?)`

	result, err := mysql.conn.ExecutePreparedQuery(query,
		driver.NombreCompleto,
		driver.Nacionalidad,
		driver.NombreEquipo,
		driver.NumeroCarro,
		driver.Edad)

	if err != nil {
		return fmt.Errorf("error al guardar el conductor: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener el ID del conductor insertado: %v", err)
	}

	driver.ID = int(id)
	return nil
}

func (mysql *MySQL) ObtenerTodos() ([]entities.Driver, error) {
	query := `SELECT id, nombre_completo, nacionalidad, nombre_equipo, 
              numero_carro, edad, fecha_creacion FROM conductores`

	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var drivers []entities.Driver
	for rows.Next() {
		var driver entities.Driver
		if err := rows.Scan(
			&driver.ID,
			&driver.NombreCompleto,
			&driver.Nacionalidad,
			&driver.NombreEquipo,
			&driver.NumeroCarro,
			&driver.Edad,
			&driver.FechaCreacion); err != nil {
			return nil, fmt.Errorf("error al escanear conductor: %v", err)
		}
		drivers = append(drivers, driver)
	}

	return drivers, nil
}

func (mysql *MySQL) ObtenerPorId(id int) (*entities.Driver, error) {
	query := `SELECT id, nombre_completo, nacionalidad, nombre_equipo, 
              numero_carro, edad, fecha_creacion 
              FROM conductores WHERE id = ?`

	rows := mysql.conn.FetchRows(query, id)
	defer rows.Close()

	var driver entities.Driver
	if rows.Next() {
		err := rows.Scan(
			&driver.ID,
			&driver.NombreCompleto,
			&driver.Nacionalidad,
			&driver.NombreEquipo,
			&driver.NumeroCarro,
			&driver.Edad,
			&driver.FechaCreacion)
		if err != nil {
			return nil, fmt.Errorf("error al escanear conductor: %v", err)
		}
		return &driver, nil
	}
	return nil, fmt.Errorf("conductor no encontrado")
}

func (mysql *MySQL) Actualizar(driver *entities.Driver) error {
	query := `UPDATE conductores 
              SET nombre_completo = ?, 
                  nacionalidad = ?, 
                  nombre_equipo = ?, 
                  numero_carro = ?, 
                  edad = ?
              WHERE id = ?`

	result, err := mysql.conn.ExecutePreparedQuery(query,
		driver.NombreCompleto,
		driver.Nacionalidad,
		driver.NombreEquipo,
		driver.NumeroCarro,
		driver.Edad,
		driver.ID)

	if err != nil {
		return fmt.Errorf("error al actualizar el conductor: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró el conductor para actualizar")
	}

	return nil
}

func (mysql *MySQL) Eliminar(id int) error {
	query := "DELETE FROM conductores WHERE id = ?"

	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar el conductor: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró el conductor para eliminar")
	}

	return nil
}
