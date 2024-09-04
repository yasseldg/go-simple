package rMongo

import (
	"fmt"

	"github.com/yasseldg/mgm/v4"
)

// CreateMany
func CreateMany[T []InterModel](inters T, coll InterColl) error {
	if len(inters) == 0 {
		return nil
	}

	models := make([]mgm.Model, 0, len(inters))
	for _, inter := range inters {
		models = append(models, inter)
	}

	err := coll.Coll().CreateMany(models)
	if err != nil {
		return fmt.Errorf("%s.CreateMany(objs): %s  ..  objs: %#v", coll.Prefix(), err, models)
	}
	return nil
}
