package rSort

type Sorts struct {
	InterOper
}

func New(oper InterOper) *Sorts {
	return &Sorts{InterOper: oper}
}

func (f *Sorts) Clone() Inter {
	return &Sorts{InterOper: f.InterOper.Clone_()}
}

func (f *Sorts) Oper() InterOper {
	return f.InterOper
}

func (s *Sorts) IdAsc() Inter {
	s.InterOper.Asc("_id")
	return s
}

func (s *Sorts) TsAsc() Inter {
	s.InterOper.Asc("ts")
	return s
}

func (s *Sorts) TsDesc() Inter {
	s.InterOper.Desc("ts")
	return s
}
