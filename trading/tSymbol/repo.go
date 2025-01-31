package tSymbol

import (
	"fmt"

	"github.com/yasseldg/go-simple/repos/rMongo"
)

type GetByIDFunc func(rMongo.ObjectID) (Inter, error)

type repo struct {
	rMongo.InterRepo
}

func NewRepo(inter rMongo.InterRepo) InterRepo {
	return &repo{
		InterRepo: inter,
	}
}

func (r *repo) GetByID(id rMongo.ObjectID) (Inter, error) {

	var obj model
	err := r.Clone().FindById(id, &obj)
	if err != nil {
		return nil, err
	}

	return &base{obj}, nil
}

func (r *repo) GetByExchangeNames(exchange string, names ...string) (InterIterLimited, error) {

	var symbols []model
	err := r.Clone().Filters(NewFilters().Exchange(exchange).Name_In(names...)).
		Sorts(NewSorts().NameAsc()).Find(&symbols)
	if err != nil {
		return nil, fmt.Errorf("coll.Find(): %s", err)
	}

	iter := NewIterLimited()

	for _, symbol := range symbols {
		iter.Add(&base{symbol})
	}

	return iter, nil
}
