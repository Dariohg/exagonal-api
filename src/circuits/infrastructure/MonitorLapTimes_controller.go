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

func (mltc *MonitorLapTimesController) Execute(c *gin.Context) {
	circuitoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de circuito inválido"})
		return
	}

	// Configurar el polling
	ticker := time.NewTicker(1 * time.Second) // Reducimos a 1 segundo para mayor reactividad
	defer ticker.Stop()

	maxPolls := 60 // Aumentamos a 1 minuto (60 segundos) para seguir la carrera más tiempo
	polls := 0

	for {
		select {
		case <-ticker.C:
			polls++

			lapTimes, err := mltc.useCase.Execute(circuitoID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"circuito_id":    circuitoID,
				"timestamp":      time.Now(),
				"poll_number":    polls,
				"tiempos_vuelta": lapTimes,
			})

			// Si alcanzamos el máximo de consultas, terminamos
			if polls >= maxPolls {
				return
			}
		}
	}
}
