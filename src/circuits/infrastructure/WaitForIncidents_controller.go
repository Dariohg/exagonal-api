package infrastructure

import (
	"f1-hex-api/src/circuits/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type WaitForIncidentsController struct {
	useCase *application.WaitForIncidents
}

func NewWaitForIncidentsController(useCase *application.WaitForIncidents) *WaitForIncidentsController {
	return &WaitForIncidentsController{useCase: useCase}
}

func (wfic *WaitForIncidentsController) Execute(c *gin.Context) {
	circuitoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de circuito inválido"})
		return
	}

	// Obtener el último ID visto por el cliente
	ultimoIDStr := c.DefaultQuery("ultimo_id", "0")
	ultimoID, err := strconv.Atoi(ultimoIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID último inválido"})
		return
	}

	// Configurar el long polling
	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			// Si pasaron 30 segundos sin incidentes nuevos
			c.JSON(http.StatusOK, gin.H{
				"mensaje":     "No hay nuevos incidentes",
				"circuito_id": circuitoID,
				"timestamp":   time.Now(),
			})
			return

		case <-ticker.C:
			// Cada 500ms verificamos si hay incidentes nuevos
			incidents, err := wfic.useCase.Execute(circuitoID, ultimoID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Si encontramos nuevos incidentes, los devolvemos inmediatamente
			if len(incidents) > 0 {
				c.JSON(http.StatusOK, gin.H{
					"circuito_id": circuitoID,
					"timestamp":   time.Now(),
					"incidentes":  incidents,
				})
				return
			}
			// Si no hay incidentes nuevos, seguimos esperando
		}
	}
}
