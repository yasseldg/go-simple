package rMongo

import (
	"context"

	"github.com/yasseldg/go-simple/repos/rIndex"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/collection"
)

type Inter interface {
	Log()

	SetDebug(bool)

	GetColl(ctx context.Context, env, conn_name, db_name, coll_name string, indexes ...rIndex.Inter) (collection.Inter, error)
}

type InterColl interface {
	collection.Inter
}
