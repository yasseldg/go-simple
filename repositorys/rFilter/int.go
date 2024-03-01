package rFilter

// int
func (f *Filters) Int(field string, value int) *Filters {
	f.Inter.Append(field, value)
	return f
}

func (f *Filters) Int_in(field string, values ...int) *Filters {
	f.Inter.In(field, values)
	return f
}

func (f *Filters) Int_nin(field string, values ...int) *Filters {
	f.Inter.Nin(field, values)
	return f
}

func (f *Filters) Int_gt(field string, value int) *Filters {
	f.Inter.Gt(field, value)
	return f
}

func (f *Filters) Int_gte(field string, value int) *Filters {
	f.Inter.Gte(field, value)
	return f
}

func (f *Filters) Int_lt(field string, value int) *Filters {
	f.Inter.Lt(field, value)
	return f
}

func (f *Filters) Int_lte(field string, value int) *Filters {
	f.Inter.Lte(field, value)
	return f
}

func (f *Filters) Int_gt_lt(field string, value_1, value_2 int) *Filters {
	f.Inter.GtLt(field, value_1, value_2)
	return f
}

func (f *Filters) Int_gt_lte(field string, value_1, value_2 int) *Filters {
	f.Inter.GtLte(field, value_1, value_2)
	return f
}

func (f *Filters) Int_gte_lt(field string, value_1, value_2 int) *Filters {
	f.Inter.GteLt(field, value_1, value_2)
	return f
}

func (f *Filters) Int_gte_lte(field string, value_1, value_2 int) *Filters {
	f.Inter.GteLte(field, value_1, value_2)
	return f
}
