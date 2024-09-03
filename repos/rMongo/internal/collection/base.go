package collection

import (
	"context"
	"fmt"

	"github.com/yasseldg/go-simple/repos/rMongo/internal"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/connection"

	"github.com/yasseldg/go-simple/repos/rIndex"

	"github.com/yasseldg/go-simple/logs/sLog"

	"github.com/yasseldg/mgm/v4"
)

type Base struct {
	coll *mgm.Collection
	env  internal.Environment

	prefix string
	conn   string
}

func New(coll *mgm.Collection, db_name string, conn *connection.Base) *Base {
	return &Base{
		coll: coll,
		env:  internal.GetEnvironment(conn.Environment),

		prefix: fmt.Sprintf("%s.%s", db_name, coll.Name()),
		conn:   fmt.Sprintf("%s .. %s", conn.Host, conn.Environment),
	}
}

func (c *Base) Name() string {
	return c.coll.Name()
}

func (c *Base) Prefix() string {
	return c.prefix
}

func (c *Base) Coll() *mgm.Collection {
	return c.coll
}

func (c *Base) String() string {
	return fmt.Sprintf("coll: %s  ..  %s", c.conn, c.prefix)
}

func (c *Base) Log() {
	sLog.Info("Mongo: %s", c.String())
}

func (c *Base) Drop(ctx context.Context, indexes ...rIndex.Inter) error {
	if c.env == internal.Env_Read {
		return nil
	}

	err := c.coll.Drop(ctx)
	if err != nil {
		sLog.Error("DropColl: %s err: %s", c.prefix, err)
		return err
	}

	sLog.Info("Delete coll: %s", c.prefix)

	err = c.CreateIndexes(ctx, indexes...)
	if err != nil {
		return err
	}

	return nil
}
