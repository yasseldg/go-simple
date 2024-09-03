package client

type Map map[string]*Base

func (cs Map) Get(conn_name string) *Base {
	if c, ok := cs[conn_name]; ok {
		return c
	}
	return nil
}

func (m Map) Set(conn_name string, base *Base) {
	if m == nil {
		m = make(Map)
	}

	if _, ok := m[conn_name]; ok {
		return
	}

	m[conn_name] = base
}
