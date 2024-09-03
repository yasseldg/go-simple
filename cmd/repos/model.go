package repos

import (
	"fmt"

	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rIndex"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/repos/rSort"
	"github.com/yasseldg/go-simple/trading/tSide"
)

type InterModel interface {
	rMongo.InterModelDateState

	InterBasic

	CCode() string
	CSymbol() string
	CSide() tSide.Side
}

type Model struct {
	rMongo.DefaultModelDateState `bson:",inline"`
	ModelBasic                   `bson:",inline"`

	Name   string     `bson:"name" json:"name"`
	Code   string     `bson:"code" json:"code"`
	Symbol string     `bson:"symbol" json:"symbol"`
	Side   tSide.Side `bson:"sd" json:"sd"`
}
type Models []Model

// filters

type Filters struct{ rFilter.Filters }

func NewFilter() *Filters {
	return &Filters{Filters: rMongo.NewFilter()}
}

func (f *Filters) Uuid(uuid string) *Filters { f.Append("uuid", uuid); return f }

func (f *Filters) Code(code string) *Filters { f.Append("code", code); return f }

func (f *Filters) Side(side tSide.Side) *Filters { f.Int("sd", int(side)); return f }

func (f *Filters) State(state string) *Filters { f.Append("st", state); return f }

func (f *Filters) Symbol(symbol string) *Filters { f.Append("symbol", symbol); return f }

// sorts

type Sorts struct {
	rSort.Sorts
}

func NewSort() *Sorts {
	return &Sorts{Sorts: rMongo.NewSort()}
}

func (s *Sorts) UuidAsc() *Sorts { s.Asc("uuid"); return s }

func (s *Sorts) CodeAsc() *Sorts { s.Asc("code"); return s }

func (s *Sorts) SymbolAsc() *Sorts { s.Asc("symbol"); return s }

func (s *Sorts) SideAsc() *Sorts { s.Asc("sd"); return s }

func (s *Sorts) Indexes() *Sorts {
	s.CodeAsc().SymbolAsc().SideAsc().UuidAsc()
	return s
}

// indexes

func Indexes() rIndex.Indexes {
	return rIndex.Indexes{
		rIndex.New(NewSort().Indexes(), true),
		rIndex.New(NewSort().UuidAsc(), true),
	}
}

//  model methods

// TODO

func (m *Model) String() string {
	return fmt.Sprintf("uuid: %s .. name: %s .. code: %s .. symbol: %s .. side: %s .. state: %s", m.Uuid, m.Name, m.Code, m.Symbol, m.Side, m.State)
}

func (m *Model) CCode() string {
	return m.Code
}

func (m *Model) CSymbol() string {
	return m.Symbol
}

func (m *Model) CSide() tSide.Side {
	return m.Side
}
