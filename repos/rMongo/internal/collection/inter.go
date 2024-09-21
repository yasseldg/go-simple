package collection

import (
	"context"

	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rIndex"
	"github.com/yasseldg/go-simple/repos/rSort"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Inter interface {
	InterBase
	InterOper
	InterAgregate

	Clone() *Full

	Pipeline(mongo.Pipeline) *Full
	Filters(rFilter.Inter) *Full
	Sorts(rSort.Inter) *Full
	Limit(int64) *Full
}

type InterBase interface {
	String() string
	Log()

	Prefix() string
	Coll() *mgm.Collection

	Drop(context.Context, ...rIndex.Inter) error
}

type InterOper interface {
	Create(model mgm.Model) error
	Update(model mgm.Model) error
	Upsert(model mgm.Model) error
	UpsertDoc(doc interface{}) error
	Count() (int64, error)
	Find(models interface{}) error
	FindOne(model mgm.Model) error
	FindById(id interface{}, model mgm.Model) error
}

type InterAgregate interface {
	Agregates(docs interface{}) error
	AgregatesWithCtx(ctx context.Context, docs interface{}) error
	AgregatesCount() ([]bson.M, error)
	AgregatesCountWithCtx(ctx context.Context) ([]bson.M, error)
}
