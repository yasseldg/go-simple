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

func TestAccuCh(t *testing.T) {
	mockRepo := &mockInterRepo{}
	saveCh := make(chan struct{}, 1)
	accu := NewAccuCH(mockRepo, 10, saveCh)

	assert.NotNil(t, accu)
	assert.Equal(t, 10, accu.Limit())
	assert.Equal(t, saveCh, accu.SaveCh)

	saveCh <- struct{}{}
	close(saveCh)

	accu.Save()
	assert.True(t, accu.Empty())
}

func TestBase(t *testing.T) {
	mockColl := &mockInterRepo{}
	limit := 10
	accu := New(mockColl, limit)

	assert.NotNil(t, accu)
	assert.Equal(t, limit, accu.Limit())
	assert.Equal(t, 0, accu.Count())
	assert.True(t, accu.Empty())

	model := &rMongo.Model{}
	accu.Add(model)

	assert.Equal(t, 1, accu.Count())
	assert.False(t, accu.Empty())

	accu.Add(model)
	assert.Equal(t, 0, accu.Count())
	assert.True(t, accu.Empty())

	clonedAccu := accu.Clone()

	assert.NotNil(t, clonedAccu)
	assert.Equal(t, limit, clonedAccu.Limit())
	assert.Equal(t, 0, clonedAccu.Count())
	assert.True(t, clonedAccu.Empty())
}
