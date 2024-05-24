package repos

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repositorys/rFilter"
	"github.com/yasseldg/go-simple/repositorys/rMongo"
	"github.com/yasseldg/go-simple/repositorys/rSort"
	"github.com/yasseldg/go-simple/trading/tSide"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/bson"
)

type Model struct {
	mgm.DefaultModelDateState `bson:",inline"`

	Uuid   string     `bson:"uuid" json:"uuid"`
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

type Sorts struct{ rSort.Sorts }

func NewSort() *Sorts {
	return &Sorts{Sorts: rMongo.NewSort()}
}

func (s *Sorts) Fields() bson.D {
	return rMongo.GetFields(NewSort().Indexes().Sorts)
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

func Indexes() rMongo.Indexes {
	return rMongo.Indexes{rMongo.Index{Fields: NewSort().Indexes().Fields(), Unique: true}}
}

//  model methods

// TODO

func (m *Model) String() string {
	return fmt.Sprintf("uuid: %s .. name: %s .. code: %s .. symbol: %s .. side: %s .. state: %s", m.Uuid, m.Name, m.Code, m.Symbol, m.Side, m.State)
}

func (m *Model) Log() {
	sLog.Info("Model: %s", m.String())
}
