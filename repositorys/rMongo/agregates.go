package rMongo

import (
	"context"

	"github.com/yasseldg/go-simple/logs/sLog"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/bson"
)

func (c Collection) Agregates(docs interface{}) error {
	return c.AgregatesWithCtx(mgm.Ctx(), docs)
}

func (c Collection) AgregatesWithCtx(ctx context.Context, docs interface{}) error {

	cursor, err := c.collection.Aggregate(ctx, c.pipeline)
	if err != nil {
		sLog.Error("sMongo: %s.AgregatesWithCtx: %s", c.prefix, err.Error())
	} else {
		err = cursor.All(ctx, docs)
		if err != nil {
			sLog.Error("sMongo: %s.AgregatesWithCtx: cursor.All(): %s", c.prefix, err.Error())
		}
	}
	return err
}

func (c Collection) AgregatesCount() ([]bson.M, error) {
	return c.AgregatesCountWithCtx(mgm.Ctx())
}

func (c Collection) AgregatesCountWithCtx(ctx context.Context) ([]bson.M, error) {

	var result []bson.M

	cursor, err := c.collection.Aggregate(mgm.Ctx(), c.pipeline)
	if err != nil {
		sLog.Error("sMongo: %s.AgregatesCountWithCtx: %s", c.prefix, err.Error())
	} else {
		err = cursor.All(mgm.Ctx(), &result)
		if err != nil {
			sLog.Error("sMongo: %s.AgregatesCountWithCtx: cursor.All(): %s", c.prefix, err.Error())
		}
	}
	return result, err
}
