package rAccu

import (
	"github.com/yasseldg/go-simple/data/dAccu"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type Inter interface {
	dAccu.Inter

	Log(string)

	Coll() rMongo.InterRepo
	Clone() Inter

	Add(rMongo.InterModel)
}
