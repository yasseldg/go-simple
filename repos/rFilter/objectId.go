package rFilter

import (
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sId"
)

// ObjectId
func (f *Filters) ObjectId(field string, value sId.ObjectId) Inter {
	sLog.Debug("ObjectId: field: %s  value: %v", field, value)
	f.InterOper.Append(field, value)
	return f
}

func (f *Filters) ObjectId_in(field string, values ...sId.ObjectId) Inter {
	f.InterOper.In(field, values)
	return f
}

func (f *Filters) ObjectId_gt(field string, value interface{}) Inter {
	f.InterOper.Gt(field, value)
	return f
}
