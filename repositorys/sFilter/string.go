package sFilter

// ----- Strings Filters

func (f *Filters) String_in(field string, values ...string) *Filters {
	f.Inter.In(field, values)
	return f
}

func (f *Filters) String_nin(field string, values ...string) *Filters {
	f.Inter.Nin(field, values)
	return f
}

func (f *Filters) String_like(field string, value string) *Filters {
	f.Inter.Like(field, value)
	return f
}
