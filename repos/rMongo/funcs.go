package rMongo

import (
	"fmt"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetID(id interface{}) primitive.ObjectID {
	if strId, ok := id.(string); ok {
		objId, _ := primitive.ObjectIDFromHex(strId)
		return objId
	}
	return id.(primitive.ObjectID)
}
