package rFilter

import "time"

func (f *Filters) Time_gt(field string, value time.Time) *Filters {
	f.InterOper.Gt(field, value)
	return f
}

func (f *Filters) Time_gte(field string, value time.Time) *Filters {
	f.InterOper.Gte(field, value)
	return f
}

func (f *Filters) Time_lt(field string, value time.Time) *Filters {
	f.InterOper.Lt(field, value)
	return f
}

func (f *Filters) Time_lte(field string, value time.Time) *Filters {
	f.InterOper.Lte(field, value)
	return f
}

func (f *Filters) Time_gt_lt(field string, value_1, value_2 time.Time) *Filters {
	f.InterOper.GtLt(field, value_1, value_2)
	return f
}

func (f *Filters) Time_gt_lte(field string, value_1, value_2 time.Time) *Filters {
	f.InterOper.GtLte(field, value_1, value_2)
	return f
}

func (f *Filters) Time_gte_lt(field string, value_1, value_2 time.Time) *Filters {
	f.InterOper.GteLt(field, value_1, value_2)
	return f
}

func (f *Filters) Time_gte_lte(field string, value_1, value_2 time.Time) *Filters {
	f.InterOper.GteLte(field, value_1, value_2)
	return f
}
