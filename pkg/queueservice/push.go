package queueservice

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (qs *QueueService) pushHandler(c *gin.Context) {
	queueName := c.Param("name")
	q := qs.manager.GetQueue(queueName)

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer c.Request.Body.Close()

	q.Push(body)
	c.Status(http.StatusCreated)
}
