package rMongo

import (
	"github.com/yasseldg/mgm/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// mgm interfaces

type InterModel interface {
	mgm.Model

	ID() ObjectID
}

type InterDate interface {
	mgm.Date
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
	Model          `bson:",inline"`
	mgm.DateFields `bson:",inline"`
}

type ModelState struct {
	Model          `bson:",inline"`
	mgm.StateField `bson:",inline"`
}

type ModelDateState struct {
	Model           `bson:",inline"`
	mgm.DateFields  `bson:",inline"`
	mgm.StateFields `bson:",inline"`
}

// ObjectID es un alias de string para trabajar con IDs de MongoDB
type ObjectID = primitive.ObjectID

func (m *Model) ID() ObjectID {
	return m.IDField.ID
}
