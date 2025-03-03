package rFilter

// ----- Strings Filters

func (f *Filters) String_in(field string, values ...string) *Filters {
	f.InterOper.In(field, values)
	return f
}

func (f *Filters) String_nin(field string, values ...string) *Filters {
	f.InterOper.Nin(field, values)
	return f
}

func (f *Filters) String_like(field string, value string) *Filters {
	f.InterOper.Like(field, value)
	return f
}
