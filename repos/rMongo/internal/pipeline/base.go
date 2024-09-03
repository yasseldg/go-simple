package pipeline

import (
	"github.com/yasseldg/go-simple/logs/sLog"

	"go.mongodb.org/mongo-driver/bson"
)

type Base struct {
	fields bson.D
}

func New() *Base {
	return &Base{fields: bson.D{}}
}

func (p Base) Fields() bson.D {
	return p.fields
}

func (f Base) Log(name string) {
	sLog.Debug("%s: Filter Mongo: %v", name, f.fields)
}

func (p *Base) Append(field string, value interface{}) {
	p.fields = append(p.fields, bson.E{Key: field, Value: value})
}
