package rMongo

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create
func (c Collection) Create(obj mgm.Model) error {

	err := c.collection.CreateWithCtx(mgm.Ctx(), obj)
	if err != nil {
		return fmt.Errorf("%s err: %s  ..  obj: %#v", c.prefix, err, obj)
	}
	return nil
}

// Update
func (c Collection) Update(obj mgm.Model) error {

	err := c.collection.UpdateWithCtx(mgm.Ctx(), obj)
	if err != nil {
		return fmt.Errorf("%s err: %s  ..  obj: %#v", c.prefix, err, obj)
	}
	return nil
}

// Upsert
func (c Collection) Upsert(obj mgm.Model, field string) error {

	err := c.collection.UpdateWithCtx(mgm.Ctx(), obj, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("%s err: %s  ..  obj: %#v", c.prefix, err, obj)
	}
	return nil
}

// Count
func (c Collection) Count() (int64, error) {

	f, err := getFilter(c.filter)
	if err != nil {
		return 0, fmt.Errorf("mongo: %s.Count(): %s", c.prefix, err)
	}

	count, err := c.collection.SimpleCountWithCtx(mgm.Ctx(), f.getFields(), options.Count())
	if err != nil {
		sLog.Error("Mongo: %s.SimpleCount(filter, opts): %s  ..  opts: %#v", c.prefix, err, options.Count())
		return 0, err
	}
	return count, nil
}

// Find
func (c Collection) Find(objs interface{}) error {

	s, err := getSort(c.sort)
	if err != nil {
		return fmt.Errorf("mongo: %s.FindOne(): %s", c.prefix, err)
	}

	opts := options.Find().SetSort(s.getFields())
	if c.limit > 0 {
		opts.SetLimit(c.limit)
	}

	f, err := getFilter(c.filter)
	if err != nil {
		return fmt.Errorf("mongo: %s.Find(): %s", c.prefix, err)
	}

	err = c.collection.SimpleFindWithCtx(mgm.Ctx(), objs, f.getFields(), opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sLog.Debug("Mongo: %s.SimpleFind(objs, filter, opts): %s", c.prefix, err)
			return err
		}
		sLog.Error("Mongo: %s.SimpleFind(objs, filter, opts): %s  ..  filter: %#v  ..  opts: %#v", c.prefix, err, c.filter, opts)
		return err
	}
	return nil
}

// FindOne
func (c Collection) FindOne(obj mgm.Model) error {

	f, err := getFilter(c.filter)
	if err != nil {
		return fmt.Errorf("mongo: %s.FindOne(): %s", c.prefix, err)
	}

	s, err := getSort(c.sort)
	if err != nil {
		return fmt.Errorf("mongo: %s.FindOne(): %s", c.prefix, err)
	}

	opts := options.FindOne().SetSort(s.getFields())
	err = c.collection.FirstWithCtx(mgm.Ctx(), f.getFields(), obj, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sLog.Debug("Mongo: %s.First(filters, obj, &opts): %s", c.prefix, err)
			return err
		}
		sLog.Error("Mongo: %s.First(filters, obj, &opts): %s  ..  filter: %#v  ..  opts: %#v", c.prefix, err, c.filter, opts)
		return err
	}
	return nil
}

// FindById
func (c Collection) FindById(id interface{}, obj mgm.Model) error {

	err := c.collection.FindByIDWithCtx(mgm.Ctx(), id, obj)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sLog.Debug("Mongo: %s.FindByID(id, obj): %s", c.prefix, err)
			return err
		}
		sLog.Error("Mongo: %s.FindByID(id, obj): %s  ..  id: %s", c.prefix, err, id)
		return err
	}
	return nil
}
