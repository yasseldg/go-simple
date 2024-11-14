package client

import (
	"context"
	"fmt"
	"time"

	"github.com/yasseldg/go-simple/repos/rMongo/internal/collection"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/connection"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/database"

	"github.com/yasseldg/go-simple/repos/rIndex"

	"github.com/yasseldg/go-simple/configs/sEnv"

	"go.mongodb.org/mongo-driver/mongo"
)

type Base struct {
	conn connection.Base

	client *mongo.Client

	databases database.Map
}

func New(conn connection.Base, client *mongo.Client) *Base {
	return &Base{
		conn:      conn,
		client:    client,
		databases: make(database.Map),
	}
}

func (c *Base) Databases() database.Map {
	return c.databases
}

func (c *Base) Env() string {
	return c.conn.Env()
}

func (c *Base) GetColl(ctx context.Context, env, db_name, coll_name string, indexes ...rIndex.Inter) (*collection.Base, error) {

	db_name = sEnv.Get(fmt.Sprint("DB_", env), db_name)

	db, err := c.GetDatabase(db_name)
	if err != nil {
		return nil, err
	}

	coll_name = sEnv.Get(fmt.Sprint("COLL_", env), coll_name)

	coll, err := db.GetCollection(ctx, coll_name, &c.conn, indexes...)
	if err != nil {
		return nil, err
	}

	return coll, nil
}

// admin functions

func (c *Base) Timeout() *time.Duration {
	return c.client.Timeout()
}

func (c *Base) NumberSessionsInProgress() int {
	return c.client.NumberSessionsInProgress()
}

func (c *Base) ListDatabaseNames(ctx context.Context, filter interface{}) ([]string, error) {
	return c.client.ListDatabaseNames(ctx, filter)
}

func (c *Base) GetDatabase(db_name string) (*database.Base, error) {

	db := c.databases.Get(db_name)
	if db != nil {
		return db, nil
	}

	return c.setDatabase(db_name)
}

// private methods

func (c *Base) setDatabase(db_name string) (*database.Base, error) {

	db := c.client.Database(db_name)
	if db == nil {
		return nil, fmt.Errorf(" client.Database( %s ) is nil", db_name)
	}

	c.databases.Set(database.New(db))

	return c.databases[db.Name()], nil
}
