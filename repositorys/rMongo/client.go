package rMongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	connection Connection

	client *mongo.Client

	databases DatabasesMap
}
type ClientsMap map[string]*Client

func (c *Client) getColl(db_name, coll_name string, indexes ...Index) (Collection, error) {

	db, err := c.getDatabase(db_name)
	if err != nil {
		return Collection{}, err
	}

	coll, err := db.getCollection(coll_name, &c.connection)
	if err != nil {
		return Collection{}, err
	}

	coll.createIndexes(indexes)

	return coll, nil
}

func (c *Client) getDatabase(db_name string) (*Database, error) {
	db := c.databases.get(db_name)
	if db != nil {
		return db, nil
	}

	return c.setDatabase(db_name)
}

func (c *Client) setDatabase(db_name string) (*Database, error) {

	db := c.client.Database(db_name)
	if db == nil {
		err := fmt.Errorf(" client.Database( %s ) is nil", db_name)
		return nil, err
	}

	c.databases[db.Name()] = &Database{
		database:    db,
		collections: make(CollectionsMap),
	}
	return c.databases[db.Name()], nil
}

func (cs ClientsMap) get(env string) *Client {
	if c, ok := cs[env]; ok {
		return c
	}
	return nil
}
