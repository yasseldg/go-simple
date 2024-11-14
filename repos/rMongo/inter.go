package rMongo

import (
	"context"

	"github.com/yasseldg/go-simple/repos/rIndex"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/client"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/collection"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/database"
)

type Inter interface {
	Log()

	SetDebug(bool)

	GetColl(ctx context.Context, env, conn_name, db_name, coll_name string, indexes ...rIndex.Inter) (collection.Inter, error)
}

type InterRepo interface {
	collection.Inter
}

type InterAdmin interface {
	Inter

	GetClient(env, conn_name string) (*client.Base, error)
}

type InterAdminClient interface {
	client.InterAdmin
}

type InterAdminDatabase interface {
	database.InterAdmin
}
