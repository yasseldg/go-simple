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
	InterTs

	Clone() *Full

	Pipeline(mongo.Pipeline) *Full
	Filters(rFilter.Inter) *Full
	Sorts(rSort.Inter) *Full
	Projections(rSort.Inter) *Full
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
	Create(mgm.Model) error
	Update(mgm.Model) error
	Upsert(mgm.Model) error
	UpsertDoc(any) error
	Count() (int64, error)
	Find(any) error
	FindOne(mgm.Model) error
	// FindById(id, model) error
	FindById(any, mgm.Model) error
}

type InterOperWithCtx interface {
	CreateWithCtx(context.Context, mgm.Model) error
	UpdateWithCtx(context.Context, mgm.Model) error
	UpsertWithCtx(context.Context, mgm.Model) error
	UpsertDocWithCtx(context.Context, any) error
	CountWithCtx(context.Context) (int64, error)
	// FindWithCtx(ctx, models) error
	FindWithCtx(context.Context, any) error
	FindOneWithCtx(context.Context, mgm.Model) error
	// FindByIdWithCtx(ctx, id, models) error
	FindByIdWithCtx(context.Context, any, mgm.Model) error
}

type InterAgregate interface {
	Agregates(any) error
	AgregatesWithCtx(context.Context, any) error
	AgregatesCount() ([]bson.M, error)
	AgregatesCountWithCtx(context.Context) ([]bson.M, error)
}

type InterTs interface {
	// First(ts_from, ts_to, model) error
	First(int64, int64, mgm.Model) error
	// FirstTs(ts_from, ts_to) int64
	FirstTs(int64, int64) int64

	// Last(ts_from, ts_to, model) error
	Last(int64, int64, mgm.Model) error
	// LastTs(ts_from, ts_to) int64
	LastTs(int64, int64) int64
}
