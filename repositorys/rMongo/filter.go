package rMongo

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repositorys/rFilter"

	"go.mongodb.org/mongo-driver/bson"
)

// Implementing uFilter interface

type Filter struct {
	fields bson.D
}

func NewFilter() rFilter.Filters {
	return *rFilter.New(&Filter{fields: bson.D{}})
}

// implementing interface rFilter.Inter

func (f Filter) Clone() rFilter.Inter {
	return &Filter{fields: f.fields}
}

func (f Filter) String() string {
	return fmt.Sprintf("Filter Mongo: %v", f.fields)
}

func (f Filter) Log(msg string) {
	sLog.Debug("%s: %s", msg, f.String())
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

//  Own methods

func (f *Filter) Id(id interface{}) {
	f.Append("_id", GetID(id))
}

// private methods

func (f Filter) getFields() bson.D {
	return f.fields
}

func getFilter(filter rFilter.Filters) (*Filter, error) {
	f, ok := filter.Inter.(*Filter)
	if !ok {
		return nil, fmt.Errorf("filter is not rMongo.Filter")
	}

	return f, nil
}
