package rIter

type InterTs interface {
	Inter

	TsFrom() int64
	TsTo() int64

	SetTsFrom(int64)
	SetTsTo(int64)
}
