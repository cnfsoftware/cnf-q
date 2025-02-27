package queueservice

import (
	"cnf-q/pkg/queue"
	"github.com/gin-gonic/gin"
	"sync"
)

type QueueService struct {
	manager    *queue.QueueManager
	bufferPool *sync.Pool
}

func NewQueueService() *QueueService {
	return &QueueService{
		manager: queue.NewQueueManager(),
	}
}

func (qs *QueueService) Run() error {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.POST("/queue/:name/push", qs.pushHandler)
	r.GET("/queue/:name/pop", qs.popHandler)
	r.GET("/queue/:name/peek", qs.peekHandler)
	r.GET("/queues", qs.listQueuesHandler)

	return r.Run(":8080")
}
