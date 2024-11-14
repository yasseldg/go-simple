package rFilter

// ObjectId
func (f *Filters) ObjectId(field string, value InterID) Inter {
	f.InterOper.Append(field, value.GetID())
	return f
}

func (f *Filters) ObjectId_in(field string, values ...InterID) Inter {

	var ids []interface{}
	for _, v := range values {
		ids = append(ids, v.GetID())
	}

	f.InterOper.In(field, ids)
	return f
}

func (f *Filters) ObjectId_gt(field string, value InterID) Inter {
	f.InterOper.Gt(field, value.GetID())
	return f
}

func (f *Filters) ObjectId_lt(field string, value InterID) Inter {
	f.InterOper.Lt(field, value.GetID())
	return f
}
