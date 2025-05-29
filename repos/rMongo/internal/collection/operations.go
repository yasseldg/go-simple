package collection

import (
	"context"
	"fmt"

	"github.com/yasseldg/go-simple/repos/rMongo/internal/filter"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/sort"

	"github.com/yasseldg/go-simple/logs/sLog"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create
func (c *Full) Create(model mgm.Model) error {
	return c.CreateWithCtx(mgm.Ctx(), model)
}

func (c *Full) CreateWithCtx(ctx context.Context, model mgm.Model) error {

	err := c.Coll().CreateWithCtx(ctx, model)
	if err != nil {
		return fmt.Errorf("%s err: %s  ..  model: %#v", c.Prefix(), err, model)
	}
	return nil
}

// Update
func (c *Full) Update(model mgm.Model) error {
	return c.UpdateWithCtx(mgm.Ctx(), model)
}

func (c *Full) UpdateWithCtx(ctx context.Context, model mgm.Model) error {

	err := c.Coll().UpdateWithCtx(ctx, model)
	if err != nil {
		return fmt.Errorf("%s err: %s  ..  obj: %#v", c.Prefix(), err, model)
	}
	return nil
}

// Upsert
func (c *Full) Upsert(model mgm.Model) error {
	return c.UpsertWithCtx(mgm.Ctx(), model)
}

func (c *Full) UpsertWithCtx(ctx context.Context, model mgm.Model) error {

	filter, _, err := c.getFilterSort()
	if err != nil {
		return fmt.Errorf("mongo: %s.Upsert(): %s", c.Prefix(), err)
	}

	err = c.Coll().UpsertWithCtx(ctx, filter, model, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("%s.UpsertWithCtx(): %s  ..  filter: %#v", c.Prefix(), err, c.filter)
	}
	return nil
}

// Upsert
func (c *Full) UpsertDoc(doc any) error {
	return c.UpsertDocWithCtx(mgm.Ctx(), doc)
}

func (c *Full) UpsertDocWithCtx(ctx context.Context, doc any) error {

	filter, _, err := c.getFilterSort()
	if err != nil {
		return fmt.Errorf("mongo: %s.UpsertDoc(): %s", c.Prefix(), err)
	}

	err = c.Coll().UpsertDocWithCtx(ctx, filter, doc, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("%s.UpsertDocWithCtx(): %s  ..  filter: %#v", c.Prefix(), err, c.filter)
	}
	return nil
}

// Delete
func (c *Full) Delete(model mgm.Model) error {
	return c.DeleteWithCtx(mgm.Ctx(), model)
}

func (c *Full) DeleteWithCtx(ctx context.Context, model mgm.Model) error {

	err := c.Coll().DeleteWithCtx(ctx, model)
	if err != nil {
		return fmt.Errorf("%s err: %s  ..  obj: %#v", c.Prefix(), err, model)
	}
	return nil
}

// DeleteMany
func (c *Full) DeleteMany(models []mgm.Model) error {
	return c.DeleteManyWithCtx(mgm.Ctx(), models)
}

func (c *Full) DeleteManyWithCtx(ctx context.Context, models []mgm.Model) error {

	filter, _, err := c.getFilterSort()
	if err != nil {
		return fmt.Errorf("mongo: %s.DeleteMany(): %s", c.Prefix(), err)
	}

	err = c.Coll().DelManyWithCtx(ctx, c.Coll(), models, filter)
	if err != nil {
		return fmt.Errorf("%s.DelManyWithCtx(): %s  ..  filter: %#v", c.Prefix(), err, c.filter)
	}
	return nil
}

// Count
func (c *Full) Count() (int64, error) {
	return c.CountWithCtx(mgm.Ctx())
}

func (c *Full) CountWithCtx(ctx context.Context) (int64, error) {

	filter, _, err := c.getFilterSort()
	if err != nil {
		return 0, fmt.Errorf("mongo: %s.Count(): %s", c.Prefix(), err)
	}

	count, err := c.Coll().SimpleCountWithCtx(ctx, filter, options.Count())
	if err != nil {
		sLog.Error("Mongo: %s.Count(filter, opts): %s  ..  opts: %#v", c.Prefix(), err, options.Count())
		return 0, err
	}
	return count, nil
}

// Find
func (c *Full) Find(models any) error {
	return c.FindWithCtx(mgm.Ctx(), models)
}

func (c *Full) FindWithCtx(ctx context.Context, models any) error {

	filter, sort, err := c.getFilterSort()
	if err != nil {
		return fmt.Errorf("mongo: %s.Find(): %s", c.Prefix(), err)
	}

	opts := options.Find().SetSort(sort)
	if c.limit > 0 {
		opts.SetLimit(c.limit)
	}

	err = c.Coll().SimpleFindWithCtx(ctx, models, filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return err
		}
		sLog.Error("Mongo: %s.SimpleFind(objs, filter, opts): %s  ..  filter: %#v  ..  opts: %#v", c.Prefix(), err, c.filter, opts)
		return err
	}
	return nil
}

// FindOne
func (c *Full) FindOne(model mgm.Model) error {
	return c.FindOneWithCtx(mgm.Ctx(), model)
}

func (c *Full) FindOneWithCtx(ctx context.Context, model mgm.Model) error {
	filter, sort, err := c.getFilterSort()
	if err != nil {
		return fmt.Errorf("mongo: %s.FindOne(): %s", c.Prefix(), err)
	}

	opts := options.FindOne().SetSort(sort)

	err = c.Coll().FirstWithCtx(ctx, filter, model, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return err
		}
		sLog.Error("Mongo: %s.First(filters, obj, &opts): %s  ..  filter: %#v  ..  opts: %#v", c.Prefix(), err, c.filter, opts)
		return err
	}
	return nil
}

// FindById
func (c *Full) FindById(id any, model mgm.Model) error {
	return c.FindByIdWithCtx(mgm.Ctx(), id, model)
}

func (c *Full) FindByIdWithCtx(ctx context.Context, id any, model mgm.Model) error {

	err := c.Coll().FindByIDWithCtx(ctx, id, model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return err
		}
		sLog.Error("Mongo: %s.FindByID(id, obj): %s  ..  id: %s", c.Prefix(), err, id)
		return err
	}
	return nil
}

func (c *Full) getFilterSort() (bson.D, bson.D, error) {
	if c.filter == nil {
		c.filter = filter.New()
	}

	filter_fields, err := filter.Fields(c.filter)
	if err != nil {
		return nil, nil, err
	}

	if c.sort == nil {
		c.sort = sort.New()
	}

	sort_fields, err := sort.Fields(c.sort)
	if err != nil {
		return nil, nil, err
	}

	return filter_fields, sort_fields, nil
}
