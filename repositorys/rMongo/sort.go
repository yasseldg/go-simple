package rMongo

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repositorys/rSort"

	"go.mongodb.org/mongo-driver/bson"
)

// Implementing uSort interface

type Sort struct {
	fields bson.D
}

func NewSort() rSort.Sorts {
	return *rSort.New(&Sort{fields: bson.D{}})
}

// implementing interface rSort.Inter

func (s Sort) Clone() rSort.Inter {
	return &Sort{fields: s.fields}
}

func (s Sort) String() string {
	return fmt.Sprintf("Sort Mongo: %v", s.fields)
}

func (s Sort) Log(msg string) {
	sLog.Debug("%s: Sort: %v", msg, s.fields)
}

// Append agrega un nuevo campo de ordenación con su dirección.
func (s *Sort) Append(field string, value interface{}) {
	s.fields = append(s.fields, bson.E{Key: field, Value: value})
}

// Ascending agrega un campo para ordenar en dirección ascendente.
func (s *Sort) Asc(field string) {
	s.Append(field, 1)
}

// Descending agrega un campo para ordenar en dirección descendente.
func (s *Sort) Desc(field string) {
	s.Append(field, -1)
}

// private methods

func (s Sort) getFields() bson.D {
	return s.fields
}

func getSort(sort rSort.Sorts) (*Sort, error) {
	s, ok := sort.Inter.(*Sort)
	if !ok {
		return nil, fmt.Errorf("filter is not rMongo.Filter")
	}

	return s, nil
}
