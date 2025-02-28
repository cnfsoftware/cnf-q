package queueservice

import (
	"cnf-q/pkg/queue"
	"github.com/gin-gonic/gin"
	"net/http"
)

type QueueService struct {
	manager     *queue.QueueManager
	port        string
	accessToken string
	tlsCertFile string
	tlsKeyFile  string
}

type Option func(*QueueService) *QueueService

func NewQueueService(opts ...Option) *QueueService {
	srv := &QueueService{
		manager: queue.NewQueueManager(),
		port:    "8080",
	}

	for _, opt := range opts {
		srv = opt(srv)
	}

	return srv
}

func WithPort(port string) Option {
	return func(srv *QueueService) *QueueService {
		if port == "" {
			return srv
		}

		srv.port = port
		return srv
	}
}

func WithAccessToken(accessToken string) Option {
	return func(srv *QueueService) *QueueService {
		if accessToken == "" {
			return srv
		}

		srv.accessToken = accessToken
		return srv
	}
}

func WithTLS(tlsCertFile, tlsKeyFile string) Option {
	return func(srv *QueueService) *QueueService {
		if tlsCertFile != "" && tlsKeyFile != "" {
			srv.tlsCertFile = tlsCertFile
			srv.tlsKeyFile = tlsKeyFile
		}
		return srv
	}
}

func (qs *QueueService) Run() error {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	if qs.accessToken != "" {
		r.Use(qs.accessTokenMiddleware())
	}

	r.POST("/queue/:name/push", qs.pushHandler)
	r.GET("/queue/:name/pop", qs.popHandler)
	r.GET("/queue/:name/peek", qs.peekHandler)
	r.GET("/queues", qs.listQueuesHandler)

	if qs.tlsCertFile != "" && qs.tlsKeyFile != "" {
		return r.RunTLS(":"+qs.port, qs.tlsCertFile, qs.tlsKeyFile)
	}

	return r.Run(":" + qs.port)
}

func (qs *QueueService) accessTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-Auth-Token") != qs.accessToken {
			c.JSON(http.StatusForbidden, gin.H{"error": "access token is incorrect"})
			c.Abort()
			return
		}
	}
}
