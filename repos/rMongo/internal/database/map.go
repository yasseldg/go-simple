package database

type Map map[string]*Base

func (m Map) Get(db_name string) *Base {
	if db, ok := m[db_name]; ok {
		return db
	}
	return nil
}

func (m Map) Set(base *Base) {
	if m == nil {
		m = make(Map)
	}

	if _, ok := m[base.database.Name()]; ok {
		return
	}

	m[base.database.Name()] = base
}
