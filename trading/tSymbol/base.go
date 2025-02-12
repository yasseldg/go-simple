package tSymbol

import (
	"fmt"
	"strings"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sTime"
)

type base struct {
	model
}

func New(name, exchange string) (*base, error) {

	if len(name) == 0 {
		return nil, fmt.Errorf("name is empty")
	}

	if len(exchange) == 0 {
		return nil, fmt.Errorf("exchange is empty")
	}

	return &base{
		model: model{
			Common: Common{
				M_name:     strings.ToUpper(name),
				M_exchange: strings.ToUpper(exchange),
			},
		},
	}, nil
}

func (b *base) String() string {
	return fmt.Sprintf("%s_%s", b.Exchange(), b.Name())
}

func (b *base) Log() {
	sLog.Info("%s .. prec: %d ..  launch: %s",
		b.String(), b.Precision(), sTime.ForLog(b.LaunchTime(), 2))
}

func (b *base) Model() InterModel {
	return &b.model
}

func (b *base) Clone() Inter {
	return &base{
		model: model{
			Common: Common{
				ModelDateState: b.ModelDateState,
				M_exchange:     b.M_exchange,
				M_name:         b.M_name,
			},
			M_precision: b.M_precision,
		},
	}
}
