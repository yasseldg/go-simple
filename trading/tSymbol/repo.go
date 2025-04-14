package tSymbol

import (
	"context"
	"fmt"
	"time"

	"github.com/yasseldg/go-simple/repos/rMongo"
)

type GetByIDFunc func(rMongo.ObjectID) (Inter, error)

type repo struct {
	rMongo.InterRepo
	groups rMongo.InterRepo
}

func NewRepo(getRepoFunc rMongo.GetRepoFunc) (InterRepo, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	symbols, err := getRepoFunc(ctx, "Symbols", "WRITE", "management", "symbols", Indexes()...)
	if err != nil {
		return nil, fmt.Errorf("failed to get collection %s: %s", "symbols", err)
	}

	groups, err := getRepoFunc(ctx, "Symbols", "WRITE", "management", "symbols_groups", Indexes()...)
	if err != nil {
		return nil, fmt.Errorf("failed to get collection %s: %s", "symbols_groups", err)
	}

	return &repo{
		InterRepo: symbols,
		groups:    groups,
	}, nil
}

func (r *repo) GroupsRepo() rMongo.InterRepo {
	return r.groups
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

	if len(names) > 0 {
		s_names, err := r.getByExchangeGroupsNames(exchange, names...)
		if err != nil && err != rMongo.ErrNoDocuments {
			return NewIterLimited(), err
		}
		names = append(names, s_names...)
	}

	return r.getSymbolsByExchangeNames(exchange, names...)
}

func (r *repo) getSymbolsByExchangeNames(exchange string, names ...string) (InterIterLimited, error) {

	iter := NewIterLimited()

	filter := NewFilters().Exchange(exchange)

	if len(names) > 0 {
		filter.Name_In(names...)
	}

	var symbols []model
	err := r.Clone().Filters(filter).
		Sorts(NewSorts().NameAsc()).Find(&symbols)
	if err != nil {
		return iter, fmt.Errorf("coll.Find(): %s", err)
	}

	for _, symbol := range symbols {
		iter.Add(&base{symbol})
	}

	return iter, nil
}

func (r *repo) getByExchangeGroupsNames(exchange string, names ...string) ([]string, error) {

	filter := NewFilters().Exchange(exchange)

	if len(names) > 0 {
		filter.Name_In(names...)
	}

	var groups []Group
	err := r.groups.Clone().Filters(filter).
		Sorts(NewSorts().NameAsc()).Find(&groups)
	if err != nil {
		return nil, fmt.Errorf("coll.Find(): %s", err)
	}

	s_names := []string{}
	for _, group := range groups {
		s_names = append(s_names, group.Names()...)
	}

	return s_names, nil
}
