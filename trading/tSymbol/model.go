package tSymbol

import (
	"github.com/yasseldg/mgm/v4"
)

// Model

type Model struct {
	mgm.ModelDateState `bson:",inline"`

	M_exchange  string `bson:"e" json:"e"`
	M_name      string `bson:"n" json:"n"`
	M_precision int    `bson:"p" json:"p"`
}
type Models []*Model

func NewModel(exchange, name string, precision int) *Model {
	return &Model{
		ModelDateState: new(mgm.DefaultModelDateState),

		M_exchange:  exchange,
		M_name:      name,
		M_precision: precision,
	}
}

func (b *Model) Exchange() string {
	return b.M_exchange
}

func (b *Model) Name() string {
	return b.M_name
}

func (b *Model) Precision() int {
	return b.M_precision
}
