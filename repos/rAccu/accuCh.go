package rAccu

import (
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type AccuCh struct {
	Inter

	SaveCh chan struct{}
}

func NewAccuCH(coll rMongo.InterRepo, limit int, saveCh chan struct{}) *AccuCh {
	accu := AccuCh{
		Inter:  New(coll, limit),
		SaveCh: saveCh,
	}

	go accu.start()

	return &accu
}

// save when saveCh is received
func (a *AccuCh) start() {
	for range a.SaveCh {
		a.Save()
	}
}
