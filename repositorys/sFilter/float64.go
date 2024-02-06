package sFilter

// float64
func (f *Filters) Float64(field string, value float64) *Filters {
	f.Inter.Append(field, value)
	return f
}

func (f *Filters) Float64_in(field string, values ...float64) *Filters {
	f.Inter.In(field, values)
	return f
}

func (f *Filters) Float64_nin(field string, values ...float64) *Filters {
	f.Inter.Nin(field, values)
	return f
}

func (f *Filters) Float64_gt(field string, value float64) *Filters {
	f.Inter.Gt(field, value)
	return f
}

func (f *Filters) Float64_gte(field string, value float64) *Filters {
	f.Inter.Gte(field, value)
	return f
}

func (f *Filters) Float64_lt(field string, value float64) *Filters {
	f.Inter.Lt(field, value)
	return f
}

func (f *Filters) Float64_lte(field string, value float64) *Filters {
	f.Inter.Lte(field, value)
	return f
}

func (f *Filters) Float64_gt_lt(field string, value_1, value_2 float64) *Filters {
	f.Inter.GtLt(field, value_1, value_2)
	return f
}

func (f *Filters) Float64_gt_lte(field string, value_1, value_2 float64) *Filters {
	f.Inter.GtLte(field, value_1, value_2)
	return f
}

func (f *Filters) Float64_gte_lt(field string, value_1, value_2 float64) *Filters {
	f.Inter.GteLt(field, value_1, value_2)
	return f
}

func (f *Filters) Float64_gte_lte(field string, value_1, value_2 float64) *Filters {
	f.Inter.GteLte(field, value_1, value_2)
	return f
}
