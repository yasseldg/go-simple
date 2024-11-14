package tSymbol

import (
	"fmt"

	"github.com/yasseldg/go-simple/repos/rMongo"
)

func GetByExchangeNames(coll rMongo.InterRepo, exchange string, names ...string) (InterIterLimited, error) {

	var symbols []model
	err := coll.Filters(NewFilters().Exchange(exchange).Name_In(names...)).
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
