package tSymbol

import (
	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rIndex"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/repos/rSort"
)

// Model

type model struct {
	rMongo.ModelDateState `bson:",inline"`

	M_exchange    string `bson:"e" json:"e"`
	M_name        string `bson:"n" json:"n"`
	M_precision   int    `bson:"p" json:"p"`
	M_launch_time int64  `bson:"l_t" json:"l_t"`
}

func (b *model) Exchange() string {
	return b.M_exchange
}

func (b *model) Name() string {
	return b.M_name
}

func (b *model) Precision() int {
	return b.M_precision
}

func (b *model) LaunchTime() int64 {
	return b.M_launch_time
}

func (s *model) SetPrecision(prec int) {
	s.M_precision = prec
}

func (s *model) SetLaunchTime(launch_time int64) {
	s.M_launch_time = launch_time
}

// filters

type Filters struct{ rFilter.Filters }

func NewFilters() *Filters {
	return &Filters{Filters: *rMongo.NewFilter()}
}

func (f *Filters) Exchange(exchange string) *Filters { f.Append("e", exchange); return f }

func (f *Filters) Name(name string) *Filters { f.Append("n", name); return f }

func (f *Filters) Name_Gt(name string) *Filters { f.Append("n", name); return f }

func (f *Filters) Name_In(names ...string) *Filters { f.String_in("n", names...); return f }

// sorts

type Sorts struct{ rSort.Sorts }

func NewSorts() *Sorts {
	return &Sorts{Sorts: *rMongo.NewSort()}
}

func (s *Sorts) ExchangeAsc() *Sorts { s.Asc("e"); return s }

func (s *Sorts) NameAsc() *Sorts { s.Asc("n"); return s }

// indexes

func Indexes() rIndex.Indexes {
	return rIndex.Indexes{
		rIndex.New(NewSorts().ExchangeAsc().NameAsc(), true),
	}
}
