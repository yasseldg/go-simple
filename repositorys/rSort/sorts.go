package rSort

type Sorts struct {
	Inter
}

func New(inter Inter) *Sorts {
	return &Sorts{Inter: inter}
}

func (f *Sorts) Clone() *Sorts {
	return &Sorts{Inter: f.Inter.Clone()}
}

func (s *Sorts) IdAsc() *Sorts {
	s.Inter.Asc("_id")
	return s
}

func (s *Sorts) TsAsc() *Sorts {
	s.Inter.Asc("ts")
	return s
}

func (s *Sorts) TsDesc() *Sorts {
	s.Inter.Desc("ts")
	return s
}
