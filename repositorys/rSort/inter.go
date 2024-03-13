package rSort

type Inter interface {
	Clone() Inter
	String() string
	Log(msg string)

	Append(key string, value interface{})

	Asc(key string)
	Desc(key string)
}
