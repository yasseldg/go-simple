package client

import (
	"context"
	"time"

	"github.com/yasseldg/go-simple/repos/rIndex"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/collection"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/database"
)

type Inter interface {
	Databases() database.Map
	Env() string
	GetColl(ctx context.Context, env, db_name, coll_name string, indexes ...rIndex.Inter) (*collection.Base, error)
}

type InterAdmin interface {
	GetDatabase(db_name string) (*database.Base, error)

	Timeout() *time.Duration
	NumberSessionsInProgress() int
	ListDatabaseNames(ctx context.Context, filter interface{}) ([]string, error)
}
