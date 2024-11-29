package rSort

type Sorts struct {
	InterOper
}

func New(oper InterOper) *Sorts {
	return &Sorts{InterOper: oper}
}

func (f *Sorts) Clone() Inter { return &Sorts{InterOper: f.InterOper.Clone_()} }

func (f *Sorts) Oper() InterOper { return f.InterOper }

func (s *Sorts) IdAsc() Inter { s.InterOper.Asc("_id"); return s }

func (s *Sorts) TsAsc() Inter { s.InterOper.Asc("ts"); return s }

func (s *Sorts) TsDesc() Inter { s.InterOper.Desc("ts"); return s }

func (s *Sorts) CreateAtAsc() Inter { s.InterOper.Asc("c_at"); return s }

func (s *Sorts) CreateAtDesc() Inter { s.InterOper.Desc("c_at"); return s }

func (s *Sorts) UpdateAtAsc() Inter { s.InterOper.Asc("u_at"); return s }

func (s *Sorts) UpdateAtDesc() Inter { s.InterOper.Desc("u_at"); return s }

func (s *Sorts) StateAsc() Inter { s.InterOper.Asc("st"); return s }

func (s *Sorts) StateDesc() Inter { s.InterOper.Desc("st"); return s }
