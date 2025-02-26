package queueservice

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (qs *QueueService) listQueuesHandler(c *gin.Context) {
	queues := qs.manager.ListQueues()
	c.JSON(http.StatusOK, gin.H{"queues": queues})
}
