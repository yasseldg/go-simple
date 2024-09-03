package pipeline

import "go.mongodb.org/mongo-driver/bson"

type Inter interface {
	Fields() bson.D
	Append(field string, value interface{})
}
