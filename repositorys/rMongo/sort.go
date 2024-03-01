package rMongo

import (
	"github.com/yasseldg/go-simple/logs/sLog"

	"go.mongodb.org/mongo-driver/bson"
)

// Implementing uSort interface

type Sort struct {
	fields bson.D
}

func Sorts() *Sort {
	return &Sort{fields: bson.D{}}
}

func (s Sort) Fields() bson.D {
	return s.fields
}

func (s Sort) Log(msg string) {
	sLog.Debug("%s: Sort: %v", msg, s.fields)
}

// Append agrega un nuevo campo de ordenación con su dirección.
func (s *Sort) Append(field string, value int) *Sort {
	s.fields = append(s.fields, bson.E{Key: field, Value: value})
	return s
}

// Ascending agrega un campo para ordenar en dirección ascendente.
func (s *Sort) Asc(field string) *Sort {
	return s.Append(field, 1)
}

// Descending agrega un campo para ordenar en dirección descendente.
func (s *Sort) Desc(field string) *Sort {
	return s.Append(field, -1)
}

func (s *Sort) TsAsc() *Sort {
	return s.Asc("ts")
}

func (s *Sort) TsDesc() *Sort {
	return s.Desc("ts")
}
