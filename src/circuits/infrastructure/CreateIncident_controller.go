package infrastructure

import (
	"f1-hex-api/src/circuits/application"
	"f1-hex-api/src/circuits/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CreateIncidentController struct {
	useCase *application.CreateIncident
}

func NewCreateIncidentController(useCase *application.CreateIncident) *CreateIncidentController {
	return &CreateIncidentController{useCase: useCase}
}

func (cic *CreateIncidentController) Execute(c *gin.Context) {
	circuitoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de circuito inválido"})
		return
	}

	var incident entities.Incident
	if err := c.ShouldBindJSON(&incident); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Establecer el ID del circuito y valores predeterminados
	incident.CircuitoID = circuitoID
	incident.Timestamp = time.Now()

	// Si no se especifica un estado, establecer como "ACTIVO"
	if incident.Estado == "" {
		incident.Estado = "ACTIVO"
	}

	if err := cic.useCase.Execute(&incident); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, incident)
}
