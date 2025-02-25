package infrastructure

import (
	"f1-hex-api/src/circuits/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type MonitorPositionsController struct {
	useCase *application.MonitorPositions
}

func NewMonitorPositionsController(useCase *application.MonitorPositions) *MonitorPositionsController {
	return &MonitorPositionsController{useCase: useCase}
}

func (mpc *MonitorPositionsController) Execute(c *gin.Context) {
	circuitoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de circuito inválido"})
		return
	}

	// Configurar el polling
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	maxPolls := 5 // Número máximo de consultas
	polls := 0

	for {
		select {
		case <-ticker.C:
			polls++

			positions, err := mpc.useCase.Execute(circuitoID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"circuito_id": circuitoID,
				"timestamp":   time.Now(),
				"poll_number": polls,
				"posiciones":  positions,
			})

			// Si alcanzamos el máximo de consultas, terminamos
			if polls >= maxPolls {
				return
			}
		}
	}
}
