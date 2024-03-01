package rMongo

import (
	"github.com/yasseldg/go-simple/logs/sLog"

	"go.mongodb.org/mongo-driver/bson"
)

type Pipeline struct {
	fields bson.D
}

func Pipelines() *Pipeline {
	return &Pipeline{fields: bson.D{}}
}

func (p Pipeline) Fields() bson.D {
	return p.fields
}

func (f Pipeline) Log(msg string) {
	sLog.Debug("%s: Filter Mongo: %v", msg, f.fields)
}

func (p *Pipeline) Append(field string, value interface{}) {
	p.fields = append(p.fields, bson.E{Key: field, Value: value})
}
