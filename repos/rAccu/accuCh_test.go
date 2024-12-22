package rAccu

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type mockInterRepo struct {
	rMongo.InterRepo
}

func (m *mockInterRepo) Clone() rMongo.InterRepo {
	return m
}

func TestNewAccuCH(t *testing.T) {
	mockRepo := &mockInterRepo{}
	saveCh := make(chan struct{}, 1)
	accu := NewAccuCH(mockRepo, 10, saveCh)

	assert.NotNil(t, accu)
	assert.Equal(t, 10, accu.Limit())
	assert.Equal(t, saveCh, accu.SaveCh)
}

func TestAccuChStart(t *testing.T) {
	mockRepo := &mockInterRepo{}
	saveCh := make(chan struct{}, 1)
	accu := NewAccuCH(mockRepo, 10, saveCh)

	saveCh <- struct{}{}
	close(saveCh)

	accu.Save()
	assert.True(t, accu.Empty())
}
