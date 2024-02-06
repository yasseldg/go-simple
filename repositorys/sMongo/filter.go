package sMongo

import (
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repositorys/sFilter"

	"go.mongodb.org/mongo-driver/bson"
)

// Implementing uFilter interface

type Filter struct {
	*sFilter.Filters // Embedding filters by adding behaviors to the filter
	fields           bson.D
}

func Filters() *Filter {
	f := &Filter{fields: bson.D{}}
	f.Filters = sFilter.New(f)
	return f
}

func (f Filter) Fields() bson.D {
	return f.fields
}

func (f Filter) Log(msg string) {
	sLog.Debug("%s: Filter Mongo: %v", msg, f.fields)
}

func (f *Filter) Append(field string, value interface{}) {
	f.fields = append(f.fields, bson.E{Key: field, Value: value})
}

func (f *Filter) In(field string, values ...interface{}) {
	f.Append(field, bson.D{{Key: "$in", Value: values}})
}

func (f *Filter) Nin(field string, values ...interface{}) {
	f.Append(field, bson.D{{Key: "$nin", Value: values}})
}

func (f *Filter) Like(field string, value string) {
	f.Append(field, bson.D{{Key: "$regex", Value: value}, {Key: "$options", Value: "i"}})
}

func (f *Filter) Gt(field string, value interface{}) {
	f.Append(field, bson.D{{Key: "$gt", Value: value}})
}

func (f *Filter) Gte(field string, value interface{}) {
	f.Append(field, bson.D{{Key: "$gte", Value: value}})
}

func (f *Filter) Lt(field string, value interface{}) {
	f.Append(field, bson.D{{Key: "$lt", Value: value}})
}

func (f *Filter) Lte(field string, value interface{}) {
	f.Append(field, bson.D{{Key: "$lte", Value: value}})
}

func (f *Filter) GtLt(field string, value_1, value_2 interface{}) {
	f.Append(field, bson.D{{Key: "$gt", Value: value_1}, {Key: "$lt", Value: value_2}})
}

func (f *Filter) GtLte(field string, value_1, value_2 interface{}) {
	f.Append(field, bson.D{{Key: "$gt", Value: value_1}, {Key: "$lte", Value: value_2}})
}

func (f *Filter) GteLt(field string, value_1, value_2 interface{}) {
	f.Append(field, bson.D{{Key: "$gte", Value: value_1}, {Key: "$lt", Value: value_2}})
}

func (f *Filter) GteLte(field string, value_1, value_2 interface{}) {
	f.Append(field, bson.D{{Key: "$gte", Value: value_1}, {Key: "$lte", Value: value_2}})
}
