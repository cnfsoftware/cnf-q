package queueservice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueueService(t *testing.T) {
	svc := NewQueueService()
	assert.NotNil(t, svc)
	assert.NotNil(t, svc.manager)
}
