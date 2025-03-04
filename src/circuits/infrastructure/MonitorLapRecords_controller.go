package infrastructure

import (
	"f1-hex-api/src/circuits/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type MonitorLapRecordsController struct {
	useCase *application.MonitorLapRecords
}

func NewMonitorLapRecordsController(useCase *application.MonitorLapRecords) *MonitorLapRecordsController {
	return &MonitorLapRecordsController{useCase: useCase}
}

func (mlrc *MonitorLapRecordsController) Execute(c *gin.Context) {
	circuitoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de circuito inválido"})
		return
	}

	// Configurar el short polling
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	maxPolls := 30 // Máximo 1 minuto (30 * 2 segundos)
	polls := 0

	for {
		select {
		case <-ticker.C:
			polls++

			record, err := mlrc.useCase.Execute(circuitoID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			response := gin.H{
				"circuito_id": circuitoID,
				"timestamp":   time.Now(),
				"poll_number": polls,
			}

			if record != nil {
				response["record_detectado"] = true
				response["record"] = record
				response["mensaje"] = "¡Récord de vuelta detectado!"
			} else {
				response["record_detectado"] = false
				response["mensaje"] = "No se ha detectado ningún récord de vuelta aún."
			}

			c.JSON(http.StatusOK, response)

			// Terminamos si alcanzamos el máximo de polls o si encontramos un récord
			if polls >= maxPolls || record != nil {
				return
			}
		}
	}
}
