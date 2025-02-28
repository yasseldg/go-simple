package tSymbol

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sTime"
)

type base struct {
	model
}

func New(name, exchange string) (*base, error) {

	common, err := newCommon(name, exchange)
	if err != nil {
		return nil, err
	}

	return &base{
		model: model{
			Common: *common,
		}}, nil
}

func (b *base) String() string {
	return fmt.Sprintf("%s_%s", b.Exchange(), b.Name())
}

func (b *base) Log() {
	sLog.Info("%s .. prec: %d ..  launch: %s  ..  config: %v",
		b.String(), b.Precision(),
		sTime.ForLog(b.LaunchTime(), 2),
		(b.M_config))
}

func (b *base) GetInterModel() InterModel {
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
			M_precision:   b.M_precision,
			M_launch_time: b.M_launch_time,
			M_min_order:   b.M_min_order,
			M_config:      b.M_config,
		},
	}
}
