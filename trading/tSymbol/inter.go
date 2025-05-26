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
	interCommon

	Precision() int
	LaunchTime() int64
	MinOrder() float64
	Location() string

	SetPrecision(int)
	SetLocation(string)

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

// privates

type interCommon interface {
	Exchange() string
	Name() string
	OwnerName() string

	SetOwnerName(string)
	ModifyName(string)
}
