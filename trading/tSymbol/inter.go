package tSymbol

import (
	"github.com/yasseldg/go-simple/data/dIter"

	"github.com/yasseldg/mgm/v4"
)

type Inter interface {
	String() string
	Log()

	Exchange() string
	Name() string
	Precision() int
	SetPrecision(int)

	IsValid() bool
	Clone() Inter
	Model() InterModel
}
type Inters []Inter

// InterModel

type InterModel interface {
	mgm.ModelDateState

	Exchange() string
	Name() string
	Precision() int
}

// InterIterLimited
type InterIterLimited interface {
	dIter.InterLimited[Inter]
}
