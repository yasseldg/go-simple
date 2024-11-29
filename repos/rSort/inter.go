package rSort

type Inter interface {
	Clone() Inter
	Oper() InterOper

	IdAsc() Inter
	TsAsc() Inter
	TsDesc() Inter

	CreateAtAsc() Inter
	CreateAtDesc() Inter
	UpdateAtAsc() Inter
	UpdateAtDesc() Inter
	StateAsc() Inter
	StateDesc() Inter
}

type InterOper interface {
	Clone_() InterOper

	String() string
	Log(name string)

	Append(key string, value interface{})

	Asc(key string)
	Desc(key string)
}
