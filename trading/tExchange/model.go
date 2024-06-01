package tExchange

import (
	"github.com/yasseldg/go-simple/repositorys/rFilter"
	"github.com/yasseldg/go-simple/repositorys/rMongo"
	"github.com/yasseldg/go-simple/repositorys/rSort"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/bson"
)

// Model

type Model struct {
	mgm.ModelDateState `bson:",inline"`

	M_name string `bson:"n" json:"n"`
}
type Models []*Model

func NewModel(name string) *Model {
	return &Model{
		ModelDateState: new(mgm.DefaultModelDateState),
		M_name:         name,
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
	return &Filters{Filters: rMongo.NewFilter()}
}

func (f *Filters) Name(name string) *Filters {
	f.Like("n", name)
	return f
}

// sorts

type Sorts struct{ rSort.Sorts }

func NewSort() *Sorts {
	return &Sorts{Sorts: rMongo.NewSort()}
}

func (s *Sorts) Fields() bson.D {
	return rMongo.GetFields(s.Sorts)
}

func (s *Sorts) NameAsc() *Sorts { s.Asc("n"); return s }

// indexes

func Indexes() rMongo.Indexes {
	return rMongo.Indexes{
		{Fields: NewSort().NameAsc(), Unique: true},
	}
}
