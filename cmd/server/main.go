package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"cnf-q/pkg/queue"
)

var manager = queue.NewQueueManager()

func main() {
	r := gin.Default()

	r.POST("/queue/:name/push", push)
	r.GET("/queue/:name/pop", pop)
	r.GET("/queue/:name/peek", peek)
	r.GET("/queues", listQueues)

	r.Run(":8080")
}

func push(c *gin.Context) {
	queueName := c.Param("name")
	q := manager.GetQueue(queueName)

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer c.Request.Body.Close()

	q.Push(body)
	c.JSON(http.StatusOK, gin.H{"message": "Add to queue"})
}

func pop(c *gin.Context) {
	queueName := c.Param("name")
	q := manager.GetQueue(queueName)

	item, err := q.Pop()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", item)
}

func peek(c *gin.Context) {
	queueName := c.Param("name")
	q := manager.GetQueue(queueName)

	item, err := q.Peek()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", item)
}

func listQueues(c *gin.Context) {
	queues := manager.ListQueues()
	c.JSON(http.StatusOK, gin.H{"queues": queues})
}
