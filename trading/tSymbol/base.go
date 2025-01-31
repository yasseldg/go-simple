package tSymbol

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sTime"
)

type base struct {
	model
}

func New(exchange, name string) *base {
	return &base{
		model: model{
			M_exchange: exchange,
			M_name:     name,
		},
	}
}

func (b *base) String() string {
	return fmt.Sprintf("%s_%s", b.Exchange(), b.Name())
}

func (b *base) Log() {
	sLog.Info("%s .. prec: %d ..  launch: %s",
		b.String(), b.Precision(), sTime.ForLog(b.LaunchTime(), 2))
}

func (b *base) IsValid() bool {
	return b.Exchange() != "" && b.Name() != ""
}

func (b *base) Model() InterModel {
	return &b.model
}

func (b *base) Clone() Inter {
	return &base{
		model: model{
			ModelDateState: b.ModelDateState,
			M_exchange:     b.M_exchange,
			M_name:         b.M_name,
			M_precision:    b.M_precision,
		},
	}
}
