package queueservice

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (qs *QueueService) peekHandler(c *gin.Context) {
	queueName := c.Param("name")
	q := qs.manager.GetQueue(queueName)

	item, err := q.Peek()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", item)
}
