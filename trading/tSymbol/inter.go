package tSymbol

import (
	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type Inter interface {
	InterModel

	sLog.InterStringLog

	GetInterModel() InterModel
	Clone() Inter
}

type InterModel interface {
	rMongo.InterModelDateState

	Exchange() string
	Name() string
	Precision() int
	LaunchTime() int64
	MinOrder() float64

	SetPrecision(int)

	GetConfig(any) error
	SetConfig(any) error
}

type InterIterLimited interface {
	dIter.InterLimited[Inter]
}

type InterRepo interface {
	rMongo.InterRepo

	GroupsRepo() rMongo.InterRepo

	GetByID(rMongo.ObjectID) (Inter, error)
	GetByExchangeNames(exchange string, names ...string) (InterIterLimited, error)
}
