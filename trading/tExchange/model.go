package tExchange

import (
	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rIndex"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/repos/rSort"
)

// Model

type Model struct {
	rMongo.InterModelDateState `bson:",inline"`

	M_name string `bson:"n" json:"n"`
}

func NewModel(name string) *Model {
	return &Model{
		InterModelDateState: new(rMongo.ModelDateState),
		M_name:              name,
	}
}

func (b *Model) Name() string {
	return b.M_name
}

func (b *Model) Inter() *Base {
	return New(b.M_name)
}

// filters

type Filters struct{ rFilter.Filters }

func NewFilter() *Filters {
	return &Filters{Filters: *rMongo.NewFilter()}
}

func (f *Filters) Name(name string) *Filters {
	f.Like("n", name)
	return f
}

// sorts

type Sorts struct{ rSort.Sorts }

func NewSort() *Sorts {
	return &Sorts{Sorts: *rMongo.NewSort()}
}

func (s *Sorts) NameAsc() *Sorts { s.Asc("n"); return s }

// indexes

func Indexes() rIndex.Indexes {
	return rIndex.Indexes{
		rIndex.New(NewSort().NameAsc(), true),
	}
}
