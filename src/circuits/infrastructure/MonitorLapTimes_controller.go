package infrastructure

import (
	"f1-hex-api/src/circuits/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type MonitorLapTimesController struct {
	useCase *application.MonitorLapTimes
}

func NewMonitorLapTimesController(useCase *application.MonitorLapTimes) *MonitorLapTimesController {
	return &MonitorLapTimesController{useCase: useCase}
}

// Execute maneja las solicitudes de monitoreo de tiempos de vuelta
// Este método utiliza short polling para obtener actualizaciones periódicas
func (mltc *MonitorLapTimesController) Execute(c *gin.Context) {
	// Obtener ID del circuito
	circuitoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de circuito inválido"})
		return
	}

	// Configuramos headers para evitar caché y optimizar la conexión
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate")
	c.Header("Content-Type", "application/json")

	// Obtener los tiempos de vuelta más recientes para cada piloto
	lapTimes, err := mltc.useCase.Execute(circuitoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respuesta inmediata con datos actuales
	c.JSON(http.StatusOK, gin.H{
		"timestamp":      time.Now(),
		"circuito_id":    circuitoID,
		"tiempos_vuelta": lapTimes,
	})
}
