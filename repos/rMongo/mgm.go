package rMongo

import "github.com/yasseldg/mgm/v4"

// mgm interfaces

type InterModel interface {
	mgm.Model
}

type InterDate interface {
	mgm.Date
}

type InterState interface {
	mgm.State
}

type InterModelDate interface {
	mgm.ModelDate
}

type InterModelState interface {
	mgm.ModelState
}

type InterModelDateState interface {
	mgm.ModelDateState
}

// mgm models

type DefaultModel struct {
	mgm.DefaultModel `bson:",inline"`
}

type DefaultModelDate struct {
	mgm.DefaultModelDate `bson:",inline"`
}

type DefaultModelState struct {
	mgm.DefaultModelState `bson:",inline"`
}

type DefaultModelDateState struct {
	mgm.DefaultModelDateState `bson:",inline"`
}
