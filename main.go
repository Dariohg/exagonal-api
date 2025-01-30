package main

import (
	"f1-hex-api/src/circuits/infrastructure"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// Crear el router de Gin
	r := gin.Default()

	// Configurar CORS si es necesario
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Ruta de prueba
	r.GET("/prueba", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensaje": "API F1 funcionando correctamente",
		})
	})

	// Configurar las rutas de circuitos
	infrastructure.ConfigureCircuitRoutes(r)

	// Obtener el puerto del archivo .env
	puerto := os.Getenv("PUERTO")
	if puerto == "" {
		puerto = "8080"
	}

	// Iniciar el servidor
	log.Printf("Servidor iniciando en el puerto %s", puerto)
	if err := r.Run(":" + puerto); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
