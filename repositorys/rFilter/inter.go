package rFilter

type Inter interface {
	InterComp

	Clone() Inter
	String() string
	Log(msg string)

	Append(key string, value interface{})
}

type InterComp interface {
	In(key string, values ...interface{})
	Nin(key string, values ...interface{})

	Like(key string, value string)

	Gt(key string, value interface{})
	Gte(key string, value interface{})
	Lt(key string, value interface{})
	Lte(key string, value interface{})

	GtLt(key string, value_1, value_2 interface{})
	GtLte(key string, value_1, value_2 interface{})
	GteLt(key string, value_1, value_2 interface{})
	GteLte(key string, value_1, value_2 interface{})
}
