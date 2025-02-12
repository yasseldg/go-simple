package tSymbol

import (
	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type Inter interface {
	InterModel

	sLog.InterStringLog

	Model() InterModel
	Clone() Inter
}

type InterModel interface {
	rMongo.InterModelDateState

	Exchange() string
	Name() string
	Precision() int

	SetPrecision(int)
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
