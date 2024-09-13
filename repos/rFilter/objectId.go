package rFilter

import (
	"github.com/yasseldg/go-simple/logs/sLog"
)

// ObjectId
func (f *Filters) ObjectId(field string, value interface{}) Inter {
	sLog.Debug("ObjectId: field: %s  value: %v", field, value)
	f.InterOper.Append(field, value)
	return f
}

func (f *Filters) ObjectId_in(field string, values ...interface{}) Inter {
	f.InterOper.In(field, values)
	return f
}

func (f *Filters) ObjectId_gt(field string, value interface{}) Inter {
	f.InterOper.Gt(field, value)
	return f
}
