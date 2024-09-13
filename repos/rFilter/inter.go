package rFilter

import (
	"github.com/yasseldg/go-simple/trading/tSide"
)

type Inter interface {
	Clone() Inter
	Oper() InterOper

	// ----- Timestamp Filters
	TsField(ts_from, ts_to int64, field string) Inter
	Ts(ts_from, ts_to int64) Inter
	TsIn(tss ...int64) Inter

	// ----- States Filters
	States(...string) Inter
	NotStates(...string) Inter

	// ----- Trading Filters
	Sides(sides ...tSide.Side) Inter

	// ----- ObjectId Filters
	ObjectId(field string, value interface{}) Inter
	ObjectId_in(field string, values ...interface{}) Inter
	ObjectId_gt(field string, value interface{}) Inter
}

type InterOper interface {
	InterComp

	Clone_() InterOper

	String() string
	Log(msg string)

	Append(key string, value interface{})
}

type InterComp interface {
	In(key string, values ...interface{})
	Nin(key string, values ...interface{})

	Like(key string, value string)

	NotNull(key string)
	NotEqual(key string, value interface{})

	Gt(key string, value interface{})
	Gte(key string, value interface{})
	Lt(key string, value interface{})
	Lte(key string, value interface{})

	GtLt(key string, value_1, value_2 interface{})
	GtLte(key string, value_1, value_2 interface{})
	GteLt(key string, value_1, value_2 interface{})
	GteLte(key string, value_1, value_2 interface{})
}
