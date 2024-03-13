package rMongo

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create
func (c Collection) Create(obj mgm.Model) bool {

	err := c.collection.Create(obj)

	if err != nil {
		sLog.Error("sMongo: %sCreate(obj): %s  ..  obj: %#v", c.prefix, err, obj)
		return false
	}
	return true
}

// CreateMany
// func CreateMany[T mgm.Model](models []T, coll Collection) bool {
// 	if len(models) > 0 {
// 		err := mgm.CreateMany(models, coll.collection)
// 		if err != nil {
// 			sLog.Error("sMongo: %s.CreateMany(objs): %s  ..  objs: %#v", coll.prefix, err, models)
// 			return false
// 		}
// 	}
// 	return true
// }

// Update
func (c Collection) Update(obj mgm.Model) bool {

	err := c.collection.Update(obj)
	if err != nil {
		sLog.Error("sMongo: %s.Update(&obj): %s  ..  obj: %#v", c.prefix, err, obj)
		return false
	}
	return true
}

// Upsert
func (c Collection) Upsert(obj mgm.Model, field string) bool {

	err := c.collection.Update(obj, options.Update().SetUpsert(true))
	if err != nil {
		sLog.Error("sMongo: %s.Upsert(&obj): %s  ..  obj: %#v", c.prefix, err, obj)
		return false
	}
	return true
}

// Count
func (c Collection) Count() (int64, error) {

	f, err := getFilter(c.filter)
	if err != nil {
		return 0, fmt.Errorf("sMongo: %s.Count(): %s", c.prefix, err)
	}

	count, err := c.collection.SimpleCount(f.getFields(), options.Count())
	if err != nil {
		sLog.Error("sMongo: %s.SimpleCount(filter, opts): %s  ..  opts: %#v", c.prefix, err, options.Count())
		return 0, err
	}
	return count, nil
}

// Find
func (c Collection) Find(objs interface{}) error {

	s, err := getSort(c.sort)
	if err != nil {
		return fmt.Errorf("sMongo: %s.FindOne(): %s", c.prefix, err)
	}

	opts := options.Find().SetSort(s.getFields())
	if c.limit > 0 {
		opts.SetLimit(c.limit)
	}

	f, err := getFilter(c.filter)
	if err != nil {
		return fmt.Errorf("sMongo: %s.Find(): %s", c.prefix, err)
	}

	err = c.collection.SimpleFind(objs, f.getFields(), opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sLog.Debug("sMongo: %s.SimpleFind(objs, filter, opts): %s", c.prefix, err)
			return err
		}
		sLog.Error("sMongo: %s.SimpleFind(objs, filter, opts): %s  ..  filter: %#v  ..  opts: %#v", c.prefix, err, c.filter, opts)
		return err
	}
	return nil
}

// FindOne
func (c Collection) FindOne(obj mgm.Model) error {

	f, err := getFilter(c.filter)
	if err != nil {
		return fmt.Errorf("sMongo: %s.FindOne(): %s", c.prefix, err)
	}

	s, err := getSort(c.sort)
	if err != nil {
		return fmt.Errorf("sMongo: %s.FindOne(): %s", c.prefix, err)
	}

	opts := options.FindOne().SetSort(s.getFields())
	err = c.collection.First(f.getFields(), obj, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sLog.Debug("sMongo: %s.First(filters, obj, &opts): %s", c.prefix, err)
			return err
		}
		sLog.Error("sMongo: %s.First(filters, obj, &opts): %s  ..  filter: %#v  ..  opts: %#v", c.prefix, err, c.filter, opts)
		return err
	}
	return nil
}

// FindById
func (c Collection) FindById(id interface{}, obj mgm.Model) error {

	err := c.collection.FindByID(id, obj)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sLog.Debug("sMongo: %sFindByID(id, obj): %s", c.prefix, err)
			return err
		}
		sLog.Error("sMongo: %sFindByID(id, obj): %s  ..  id: %s", c.prefix, err, id)
		return err
	}
	return nil
}
