package collection

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"

	"github.com/yasseldg/go-simple/repos/rMongo/internal/filter"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/sort"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create
func (c *Full) Create(model mgm.Model) error {

	err := c.Coll().CreateWithCtx(mgm.Ctx(), model)
	if err != nil {
		return fmt.Errorf("%s err: %s  ..  model: %#v", c.Prefix(), err, model)
	}
	return nil
}

// Update
func (c *Full) Update(model mgm.Model) error {

	err := c.Coll().UpdateWithCtx(mgm.Ctx(), model)
	if err != nil {
		return fmt.Errorf("%s err: %s  ..  obj: %#v", c.Prefix(), err, model)
	}
	return nil
}

// Upsert
func (c *Full) Upsert(model mgm.Model) error {

	filter, err := filter.Fields(c.filter)
	if err != nil {
		return fmt.Errorf("mongo: %s.Upsert(): %s", c.Prefix(), err)
	}

	err = c.Coll().UpsertWithCtx(mgm.Ctx(), filter, model, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("%s.UpsertWithCtx(): %s  ..  filter: %#v", c.Prefix(), err, c.filter)
	}
	return nil
}

// Upsert
func (c *Full) UpsertDoc(doc interface{}) error {

	filter, err := filter.Fields(c.filter)
	if err != nil {
		return fmt.Errorf("mongo: %s.UpsertDoc(): %s", c.Prefix(), err)
	}

	err = c.Coll().UpsertDocWithCtx(mgm.Ctx(), filter, doc, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("%s.UpsertDocWithCtx(): %s  ..  filter: %#v", c.Prefix(), err, c.filter)
	}
	return nil
}

// Count
func (c *Full) Count() (int64, error) {

	filter, err := filter.Fields(c.filter)
	if err != nil {
		return 0, fmt.Errorf("mongo: %s.Count(): %s", c.Prefix(), err)
	}

	count, err := c.Coll().SimpleCountWithCtx(mgm.Ctx(), filter, options.Count())
	if err != nil {
		sLog.Error("Mongo: %s.Count(filter, opts): %s  ..  opts: %#v", c.Prefix(), err, options.Count())
		return 0, err
	}
	return count, nil
}

// Find
func (c *Full) Find(models interface{}) error {

	sort, err := sort.Fields(c.sort)
	if err != nil {
		return fmt.Errorf("mongo: %s.Find(): %s", c.Prefix(), err)
	}

	opts := options.Find().SetSort(sort)
	if c.limit > 0 {
		opts.SetLimit(c.limit)
	}

	filter, err := filter.Fields(c.filter)
	if err != nil {
		return fmt.Errorf("mongo: %s.Find(): %s", c.Prefix(), err)
	}

	err = c.Coll().SimpleFindWithCtx(mgm.Ctx(), models, filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sLog.Debug("Mongo: %s.SimpleFind(objs, filter, opts): %s", c.Prefix(), err)
			return err
		}
		sLog.Error("Mongo: %s.SimpleFind(objs, filter, opts): %s  ..  filter: %#v  ..  opts: %#v", c.Prefix(), err, c.filter, opts)
		return err
	}
	return nil
}

// FindOne
func (c *Full) FindOne(model mgm.Model) error {

	sort, err := sort.Fields(c.sort)
	if err != nil {
		return fmt.Errorf("mongo: %s.FindOne(): %s", c.Prefix(), err)
	}

	opts := options.FindOne().SetSort(sort)

	filter, err := filter.Fields(c.filter)
	if err != nil {
		return fmt.Errorf("mongo: %s.Upsert(): %s", c.Prefix(), err)
	}

	err = c.Coll().FirstWithCtx(mgm.Ctx(), filter, model, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sLog.Debug("Mongo: %s.First(filters, obj, &opts): %s", c.Prefix(), err)
			return err
		}
		sLog.Error("Mongo: %s.First(filters, obj, &opts): %s  ..  filter: %#v  ..  opts: %#v", c.Prefix(), err, c.filter, opts)
		return err
	}
	return nil
}

// FindById
func (c *Full) FindById(id interface{}, model mgm.Model) error {

	err := c.Coll().FindByIDWithCtx(mgm.Ctx(), id, model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sLog.Debug("Mongo: %s.FindByID(id, obj): %s", c.Prefix(), err)
			return err
		}
		sLog.Error("Mongo: %s.FindByID(id, obj): %s  ..  id: %s", c.Prefix(), err, id)
		return err
	}
	return nil
}
