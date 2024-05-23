package rAccu

import (
	"github.com/yasseldg/go-simple/data/dAccu"

	"github.com/yasseldg/mgm/v4"
)

type Inter interface {
	dAccu.Inter

	Log(name string)

	Clone() Inter

	Add(model mgm.Model)
}
