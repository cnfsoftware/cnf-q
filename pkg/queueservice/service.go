package queueservice

import (
	"cnf-q/pkg/queue"
	"github.com/gin-gonic/gin"
)

type QueueService struct {
	manager *queue.QueueManager
}

func NewQueueService() *QueueService {
	return &QueueService{
		manager: queue.NewQueueManager(),
	}
}

func (qs *QueueService) Run() error {
	r := gin.Default()

	r.POST("/queue/:name/push", qs.pushHandler)
	r.GET("/queue/:name/pop", qs.popHandler)
	r.GET("/queue/:name/peek", qs.peekHandler)
	r.GET("/queues", qs.listQueuesHandler)

	return r.Run(":8080")
}
