package rFilter

// int64
func (f *Filters) Int64(field string, value int64) *Filters {
	f.Inter.Append(field, value)
	return f
}

func (f *Filters) Int64_in(field string, values ...int64) *Filters {
	f.Inter.In(field, values)
	return f
}

func (f *Filters) Int64_nin(field string, values ...int64) *Filters {
	f.Inter.Nin(field, values)
	return f
}

func (f *Filters) Int64_gt(field string, value int64) *Filters {
	f.Inter.Gt(field, value)
	return f
}

func (f *Filters) Int64_gte(field string, value int64) *Filters {
	f.Inter.Gte(field, value)
	return f
}

func (f *Filters) Int64_lt(field string, value int64) *Filters {
	f.Inter.Lt(field, value)
	return f
}

func (f *Filters) Int64_lte(field string, value int64) *Filters {
	f.Inter.Lte(field, value)
	return f
}

func (f *Filters) Int64_gt_lt(field string, value_1, value_2 int64) *Filters {
	f.Inter.GtLt(field, value_1, value_2)
	return f
}

func (f *Filters) Int64_gt_lte(field string, value_1, value_2 int64) *Filters {
	f.Inter.GtLte(field, value_1, value_2)
	return f
}

func (f *Filters) Int64_gte_lt(field string, value_1, value_2 int64) *Filters {
	f.Inter.GteLt(field, value_1, value_2)
	return f
}

func (f *Filters) Int64_gte_lte(field string, value_1, value_2 int64) *Filters {
	f.Inter.GteLte(field, value_1, value_2)
	return f
}
