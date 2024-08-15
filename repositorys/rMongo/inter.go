package rMongo

import (
	"context"

	"github.com/yasseldg/go-simple/repositorys/rFilter"
	"github.com/yasseldg/go-simple/repositorys/rSort"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/bson"
)

type InterColl interface {
	Prefix() string
	Mgm() *mgm.Collection
	String() string
	Log()

	Create(model mgm.Model) error
	Update(model mgm.Model) error
	Upsert(model mgm.Model, filter rFilter.Filters) error
	Count() (int64, error)
	Find(models interface{}) error
	FindOne(model mgm.Model) error
	FindById(id interface{}, model mgm.Model) error

	Agregates(docs interface{}) error
	AgregatesWithCtx(ctx context.Context, docs interface{}) error
	AgregatesCount() ([]bson.M, error)
	AgregatesCountWithCtx(ctx context.Context) ([]bson.M, error)

	Drop(indexes ...Index) error
	Pipeline(p Pipeline) *Collection
	Filters(f rFilter.Filters) *Collection
	Sorts(s rSort.Sorts) *Collection
	Limit(l int64) *Collection
}
