package sort

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rSort"

	"go.mongodb.org/mongo-driver/bson"
)

// Implementing rSort interface

type Sort struct {
	fields bson.D
}

func New() *rSort.Sorts {
	return rSort.New(&Sort{fields: bson.D{}})
}

// implementing interface rSort.Inter

func (s *Sort) Clone_() rSort.InterOper {
	return &Sort{fields: s.fields}
}

func (s *Sort) String() string {
	return fmt.Sprintf("Sort Mongo: %v", s.fields)
}

func (s *Sort) Log(name string) {
	sLog.Info("%s: Sort: %v", name, s.fields)
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
