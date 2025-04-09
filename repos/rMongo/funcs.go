package rMongo

import (
	"context"
	"fmt"

	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rSort"

	"github.com/yasseldg/go-simple/repos/rMongo/internal/filter"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/sort"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateMany
func CreateMany[T []InterModel](inters T, coll InterRepo) error {
	return CreateManyWithCtx(mgm.Ctx(), inters, coll)
}

func CreateManyWithCtx[T []InterModel](ctx context.Context, inters T, coll InterRepo) error {
	if len(inters) == 0 {
		return nil
	}

	models := make([]mgm.Model, 0, len(inters))
	for _, inter := range inters {
		models = append(models, inter)
	}

	err := coll.Coll().CreateManyWithCtx(ctx, models)
	if err != nil {
		return fmt.Errorf("%s.CreateMany(objs): %s  ..  objs: %#v", coll.Prefix(), err, models)
	}
	return nil
}

func GetID(id any) primitive.ObjectID {
	if strId, ok := id.(string); ok {
		objId, _ := primitive.ObjectIDFromHex(strId)
		return objId
	}
	return id.(primitive.ObjectID)
}

func FilterFields(inter rFilter.Inter) (bson.D, error) {
	return filter.Fields(inter)
}

func SortFields(inter rSort.Inter) (bson.D, error) {
	return sort.Fields(inter)
}

func GetInterID(id any) rFilter.InterID {

	field := new(mgm.IDField)

	field.SetID(GetID(id))

	return field
}

func BsonMarshal(val any) (M, error) {

	bson_bytes, err := bson.Marshal(val)
	if err != nil {
		return nil, fmt.Errorf("bson.Marshal: %s", err)
	}

	var m bson.M
	if err := bson.Unmarshal(bson_bytes, &m); err != nil {
		return nil, fmt.Errorf("bson.Unmarshal: %s", err)
	}
	return m, nil
}

func BsonUnmarshal(m M, val any) error {

	bson_bytes, err := bson.Marshal(m)
	if err != nil {
		return fmt.Errorf("bson.Marshal: %s", err)
	}

	if err := bson.Unmarshal(bson_bytes, val); err != nil {
		return fmt.Errorf("bson.Unmarshal: %s", err)
	}
	return nil
}
