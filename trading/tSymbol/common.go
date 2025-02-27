package tSymbol

import (
	"fmt"
	"strings"

	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rIndex"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/repos/rSort"
)

type Common struct {
	rMongo.ModelDateState `bson:",inline"`

	M_exchange string `bson:"e" json:"e"`
	M_name     string `bson:"n" json:"n"`
}

func newCommon(name, exchange string) (*Common, error) {

	if len(name) == 0 {
		return nil, fmt.Errorf("name is empty")
	}

	if len(exchange) == 0 {
		return nil, fmt.Errorf("exchange is empty")
	}

	return &Common{
		M_name:     strings.ToUpper(name),
		M_exchange: strings.ToUpper(exchange),
	}, nil
}

func (b *Common) Exchange() string {
	return b.M_exchange
}

func (b *Common) Name() string {
	return b.M_name
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
