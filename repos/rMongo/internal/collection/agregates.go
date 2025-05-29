package collection

import (
	"context"
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/bson"
)

func (c *Full) Agregates(docs any) error {
	return c.AgregatesWithCtx(mgm.Ctx(), docs)
}

func (c *Full) AgregatesWithCtx(ctx context.Context, docs any) error {
	if c.pipeline == nil {
		return fmt.Errorf("no pipeline")
	}

	cursor, err := c.Coll().Aggregate(ctx, c.pipeline)
	if err != nil {
		sLog.Error("rMongo: %s.AgregatesWithCtx: %s", c.Prefix(), err.Error())
		return err
	}

	if err := cursor.All(ctx, docs); err != nil {
		sLog.Error("rMongo: %s.AgregatesWithCtx: cursor.All(): %s", c.Prefix(), err.Error())
		return err
	}

	return nil
}

func (c *Full) AgregatesCount() ([]bson.M, error) {
	return c.AgregatesCountWithCtx(mgm.Ctx())
}

func (c *Full) AgregatesCountWithCtx(ctx context.Context) ([]bson.M, error) {

	if c.pipeline == nil {
		return nil, fmt.Errorf("no pipeline")
	}

	cursor, err := c.Coll().Aggregate(mgm.Ctx(), c.pipeline)
	if err != nil {
		sLog.Error("rMongo: %s.AgregatesCountWithCtx: %s", c.Prefix(), err.Error())
		return nil, err
	}

	var result []bson.M

	if err := cursor.All(mgm.Ctx(), &result); err != nil {
		sLog.Error("rMongo: %s.AgregatesCountWithCtx: cursor.All(): %s", c.Prefix(), err.Error())
		return nil, err
	}

	return result, nil
}
