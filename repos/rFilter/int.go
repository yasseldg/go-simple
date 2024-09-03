package rFilter

// int
func (f *Filters) Int(field string, value int) *Filters {
	f.InterOper.Append(field, value)
	return f
}

func (f *Filters) Int_in(field string, values ...int) *Filters {
	f.InterOper.In(field, values)
	return f
}

func (f *Filters) Int_nin(field string, values ...int) *Filters {
	f.InterOper.Nin(field, values)
	return f
}

func (f *Filters) Int_gt(field string, value int) *Filters {
	f.InterOper.Gt(field, value)
	return f
}

func (f *Filters) Int_gte(field string, value int) *Filters {
	f.InterOper.Gte(field, value)
	return f
}

func (f *Filters) Int_lt(field string, value int) *Filters {
	f.InterOper.Lt(field, value)
	return f
}

func (f *Filters) Int_lte(field string, value int) *Filters {
	f.InterOper.Lte(field, value)
	return f
}

func (f *Filters) Int_gt_lt(field string, value_1, value_2 int) *Filters {
	f.InterOper.GtLt(field, value_1, value_2)
	return f
}

func (f *Filters) Int_gt_lte(field string, value_1, value_2 int) *Filters {
	f.InterOper.GtLte(field, value_1, value_2)
	return f
}

func (f *Filters) Int_gte_lt(field string, value_1, value_2 int) *Filters {
	f.InterOper.GteLt(field, value_1, value_2)
	return f
}

func (f *Filters) Int_gte_lte(field string, value_1, value_2 int) *Filters {
	f.InterOper.GteLte(field, value_1, value_2)
	return f
}
