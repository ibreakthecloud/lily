package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ibreakthecloud/lily/pkg/models"

	kfka "github.com/ibreakthecloud/lily/pkg/kafka"
)

func monteCarloIncident(c *gin.Context) {
	var incident models.DataIssue
	if err := c.ShouldBindJSON(&incident); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	kfka.Producer.ProduceMonteCarlo(incident)
	c.JSON(http.StatusOK, gin.H{"message": "Incident sent"})
}

func annotateData(c *gin.Context) {
	var annotation models.Annotation

	if err := c.ShouldBindJSON(&annotation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	kfka.Producer.ProduceDataAnnotation(annotation)
	c.JSON(http.StatusOK, gin.H{"message": "Annotation sent"})
}
