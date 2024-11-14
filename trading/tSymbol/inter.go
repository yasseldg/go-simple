package tSymbol

import (
	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type Inter interface {
	InterModel

	String() string
	Log()

	IsValid() bool
	Model() InterModel
	Clone() Inter
}

// InterModel

type InterModel interface {
	rMongo.InterModelDateState

	Exchange() string
	Name() string
	Precision() int

	SetPrecision(int)
}

// InterIterLimited
type InterIterLimited interface {
	dIter.InterLimited[Inter]
}

type InterRepo interface {
	rMongo.InterRepo

	GetAll() ([]Inter, error)
	GetByName(name string) (Inter, error)
	GetByExchangeName(exchange, name string) (Inter, error)
}
