package rFilter

type Filters struct {
	InterOper
}

func New(inter InterOper) *Filters {
	return &Filters{InterOper: inter}
}

func (f *Filters) Clone() Inter {
	return &Filters{InterOper: f.InterOper.Clone_()}
}

func (f *Filters) Oper() InterOper {
	return f.InterOper
}
