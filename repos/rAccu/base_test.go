package rAccu

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type mockRepo struct {
	rMongo.InterRepo
}

func (m *mockRepo) Clone() rMongo.InterRepo {
	return m
}

func (m *mockRepo) Create(model rMongo.InterModel) error {
	return nil
}

func (m *mockRepo) CreateMany(models []rMongo.InterModel) error {
	return nil
}

func TestNew(t *testing.T) {
	mockColl := &mockRepo{}
	limit := 10
	accu := New(mockColl, limit)

	assert.NotNil(t, accu)
	assert.Equal(t, limit, accu.Limit())
	assert.Equal(t, 0, accu.Count())
	assert.True(t, accu.Empty())
}

func TestAdd(t *testing.T) {
	mockColl := &mockRepo{}
	limit := 2
	accu := New(mockColl, limit)

	model := &rMongo.Model{}
	accu.Add(model)

	assert.Equal(t, 1, accu.Count())
	assert.False(t, accu.Empty())

	accu.Add(model)
	assert.Equal(t, 0, accu.Count())
	assert.True(t, accu.Empty())
}

func TestSave(t *testing.T) {
	mockColl := &mockRepo{}
	limit := 2
	accu := New(mockColl, limit)

	model := &rMongo.Model{}
	accu.Add(model)
	accu.Add(model)

	assert.Equal(t, 0, accu.Count())
	assert.True(t, accu.Empty())
}

func TestClone(t *testing.T) {
	mockColl := &mockRepo{}
	limit := 2
	accu := New(mockColl, limit)

	clonedAccu := accu.Clone()

	assert.NotNil(t, clonedAccu)
	assert.Equal(t, limit, clonedAccu.Limit())
	assert.Equal(t, 0, clonedAccu.Count())
	assert.True(t, clonedAccu.Empty())
}
