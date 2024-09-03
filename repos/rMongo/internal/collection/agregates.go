package collection

import (
	"context"

	"github.com/yasseldg/go-simple/logs/sLog"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/bson"
)

func (c *Full) Agregates(docs interface{}) error {
	return c.AgregatesWithCtx(mgm.Ctx(), docs)
}

func (c *Full) AgregatesWithCtx(ctx context.Context, docs interface{}) error {

	cursor, err := c.Coll().Aggregate(ctx, c.pipeline.Fields())
	if err != nil {
		sLog.Error("sMongo: %s.AgregatesWithCtx: %s", c.Prefix(), err.Error())
	} else {
		err = cursor.All(ctx, docs)
		if err != nil {
			sLog.Error("sMongo: %s.AgregatesWithCtx: cursor.All(): %s", c.Prefix(), err.Error())
		}
	}
	return err
}

func (c *Full) AgregatesCount() ([]bson.M, error) {
	return c.AgregatesCountWithCtx(mgm.Ctx())
}

func (c *Full) AgregatesCountWithCtx(ctx context.Context) ([]bson.M, error) {

	var result []bson.M

	cursor, err := c.Coll().Aggregate(mgm.Ctx(), c.pipeline.Fields())
	if err != nil {
		sLog.Error("sMongo: %s.AgregatesCountWithCtx: %s", c.Prefix(), err.Error())
	} else {
		err = cursor.All(mgm.Ctx(), &result)
		if err != nil {
			sLog.Error("sMongo: %s.AgregatesCountWithCtx: cursor.All(): %s", c.Prefix(), err.Error())
		}
	}
	return result, err
}
