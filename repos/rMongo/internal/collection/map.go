package collection

type Map map[string]*Base

func (m Map) Get(coll_name string) *Base {
	if c, ok := m[coll_name]; ok {
		return c
	}
	return nil
}

func (m Map) Set(base *Base) {
	if m == nil {
		m = make(Map)
	}

	if _, ok := m[base.coll.Name()]; ok {
		return
	}

	m[base.coll.Name()] = base

	base.Log()
}
