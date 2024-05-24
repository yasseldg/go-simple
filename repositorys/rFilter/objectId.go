package rFilter

import (
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sId"
)

// ObjectId
func (f *Filters) ObjectId(field string, value sId.ObjectId) *Filters {
	sLog.Debug("ObjectId: field: %s  value: %v", field, value)
	f.Inter.Append(field, value)
	return f
}

func (f *Filters) ObjectId_in(field string, values ...sId.ObjectId) *Filters {
	f.Inter.In(field, values)
	return f
}

func (f *Filters) ObjectId_gt(field string, value interface{}) *Filters {
	f.Inter.Gt(field, value)
	return f
}
