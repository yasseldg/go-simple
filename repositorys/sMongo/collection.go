package sMongo

import (
	"context"

	"github.com/yasseldg/go-simple/logs/sLog"

	"github.com/yasseldg/mgm/v4"
)

type Collection struct {
	collection  *mgm.Collection
	environment Environment

	prefix string
	conn   string

	pipeline Pipeline
	filter   Filter
	sort     Sort
	limit    int64
}
type CollectionsMap map[string]*Collection

func (cm CollectionsMap) get(coll_name string) *Collection {
	if c, ok := cm[coll_name]; ok {
		return c
	}
	return nil
}

func (c Collection) Prefix() string {
	return c.prefix
}

func (c Collection) Log() {
	sLog.Info("Mongo Collection: %s  ..  %s \n", c.conn, c.prefix)
}

func (c *Collection) Drop(indexes ...Index) error {

	err := c.collection.Drop(context.TODO())
	if err != nil {
		sLog.Error("DropColl: %s err: %s", c.prefix, err)
		return err
	}

	sLog.Info("Delete coll: %s", c.prefix)

	c.createIndexes(indexes)

	return nil
}

func (c *Collection) Pipeline(p Pipeline) *Collection {
	c.pipeline = p
	return c
}

func (c *Collection) Filters(f Filter) *Collection {
	c.filter = f
	return c
}

func (c *Collection) Sorts(s Sort) *Collection {
	c.sort = s
	return c
}

func (c *Collection) Limit(l int64) *Collection {
	c.limit = l
	return c
}
