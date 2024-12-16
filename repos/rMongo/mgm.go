package rMongo

import (
	"time"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// mgm interfaces

type InterModel interface {
	mgm.Model

	ID() ObjectID
}

type InterDate interface {
	mgm.Date

	CreatedAt() time.Time
	UpdatedAt() time.Time
}

type InterState interface {
	mgm.State
}

type InterModelDate interface {
	InterModel
	InterDate
}

type InterModelState interface {
	InterModel
	InterState
}

type InterModelDateState interface {
	InterModel
	InterDate
	InterState
}

// mgm models

type Model struct {
	mgm.DefaultModel `bson:",inline"`
}

type ModelDate struct {
	Model      `bson:",inline"`
	DateFields `bson:",inline"`
}

type ModelState struct {
	Model          `bson:",inline"`
	mgm.StateField `bson:",inline"`
}

type ModelDateState struct {
	Model           `bson:",inline"`
	DateFields      `bson:",inline"`
	mgm.StateFields `bson:",inline"`
}

// Alias bson
type M = bson.M
type D = bson.D
type A = bson.A

// Alias ObjectID
type ObjectID = primitive.ObjectID

func (m *Model) ID() ObjectID {
	return m.IDField.ID
}

// DateFields

type DateFields struct {
	mgm.DateFields `bson:",inline"`
}

func (m *DateFields) CreatedAt() time.Time {
	return m.DateFields.CreatedAt
}

func (m *DateFields) UpdatedAt() time.Time {
	return m.DateFields.UpdatedAt
}
