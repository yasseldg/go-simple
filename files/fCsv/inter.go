package fCsv

import (
	"github.com/yasseldg/go-simple/files/fAccu"
	"github.com/yasseldg/go-simple/files/fIter"
)

type InterCsv interface {
	CsvHeader() []string
	CsvData() []string
}

type InterCsvName interface {
	CsvHeader(string) []string
	CsvData() []string
}

type InterAccu interface {
	fAccu.Inter

	AddHeader([]string)
	AddData([]string)
}

type InterIter interface {
	fIter.Inter
}

type InterIterFunc interface {
	fIter.Inter

	Run()
}
