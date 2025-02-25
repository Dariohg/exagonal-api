CREATE DATABASE IF NOT EXISTS f1_database;
USE f1_database;

CREATE TABLE IF NOT EXISTS conductores (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nombre_completo VARCHAR(100) NOT NULL,
    nacionalidad VARCHAR(50) NOT NULL,
    nombre_equipo VARCHAR(50) NOT NULL,
    numero_carro INT NOT NULL,
    edad INT NOT NULL,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS circuitos (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nombre VARCHAR(100) NOT NULL,
    pais VARCHAR(50) NOT NULL,
    longitud DECIMAL(10,3) NOT NULL,
    numero_vueltas INT NOT NULL,
    numero_curvas INT NOT NULL,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS circuito_conductor (
   circuito_id INT,
   conductor_id INT,
   FOREIGN KEY (circuito_id) REFERENCES circuitos(id),
    FOREIGN KEY (conductor_id) REFERENCES conductores(id),
    PRIMARY KEY (circuito_id, conductor_id)
    );

CREATE TABLE IF NOT EXISTS conductor_estado (
    id INT PRIMARY KEY AUTO_INCREMENT,
    conductor_id INT NOT NULL,
    esta_corriendo BOOLEAN DEFAULT FALSE,
    ultima_verificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (conductor_id) REFERENCES conductores(id)
    );

CREATE TABLE IF NOT EXISTS tiempos_vuelta (
    id INT PRIMARY KEY AUTO_INCREMENT,
    circuito_id INT NOT NULL,
    conductor_id INT NOT NULL,
    numero_vuelta INT NOT NULL,
    tiempo DECIMAL(10,3) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (circuito_id) REFERENCES circuitos(id),
    FOREIGN KEY (conductor_id) REFERENCES conductores(id)
    );

CREATE TABLE IF NOT EXISTS posiciones_carrera (
    id INT PRIMARY KEY AUTO_INCREMENT,
    conductor_id INT NOT NULL,
    posicion INT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (conductor_id) REFERENCES conductores(id)
    );
