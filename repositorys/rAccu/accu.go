package rAccu

import (
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repositorys/rMongo"
	"github.com/yasseldg/mgm/v4"
)

//  Accu

type Accu struct {
	coll  rMongo.Collection
	limit int

	items []mgm.Model

	count int

	empty bool
	err   error
}

func New(coll rMongo.Collection, limit int) Accu {
	return Accu{
		coll:  coll,
		limit: limit,
	}
}

func (iter *Accu) Log(name string) {
	sLog.Info("Accu ( %s ): %d", name, iter.Count())
}

func (iter *Accu) Add(model mgm.Model) {
	if model == nil {
		return
	}

	iter.items = append(iter.items, model)

	if len(iter.items) >= iter.Limit() {
		iter.save()
	}
}

func (iter *Accu) Save() {
	iter.save()
}

func (i *Accu) SetError(e error) {
	i.err = e
}

func (i Accu) Error() error {
	return i.err
}

func (i *Accu) SetEmpty(e bool) {
	i.empty = e
}

func (i Accu) Empty() bool {
	return i.empty
}

func (i Accu) Limit() int {
	return i.limit
}

func (i Accu) Count() int {
	return i.count
}

//  private methods

func (iter *Accu) save() {
	if len(iter.items) == 0 {
		return
	}

	err := rMongo.CreateMany(iter.items, iter.coll)
	if err != nil {
		iter.SetError(err)
		return
	}

	iter.count += len(iter.items)

	iter.items = []mgm.Model{}
}
