package rMongo

import (
	"fmt"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	database *mongo.Database

	collections CollectionsMap
}
type DatabasesMap map[string]*Database

func (db *Database) getCollection(coll_name string, conn *Connection) (Collection, error) {
	coll := db.collections.get(coll_name)
	if coll != nil {
		return *coll, nil
	}

	return db.setCollection(coll_name, conn)
}

func (db *Database) setCollection(coll_name string, conn *Connection) (Collection, error) {

	coll := mgm.NewCollection(db.database, coll_name)
	if coll == nil {
		err := fmt.Errorf(" mgm.NewCollection( %s , %s ) is nil", db.database.Name(), coll_name)
		return Collection{}, err
	}

	db.collections[coll.Name()] = &Collection{
		collection:  coll,
		environment: getEnvironment(conn.Environment),

		prefix: fmt.Sprintf("%s.%s", db.database.Name(), coll.Name()),
		conn:   fmt.Sprintf("%s .. %s", conn.Host, conn.Environment),

		pipeline: *Pipelines(),
		filter:   NewFilter(),
		sort:     *Sorts(),
		limit:    0,
	}

	db.collections[coll.Name()].Log()

	return *db.collections[coll.Name()], nil
}

func (dbs DatabasesMap) get(db_name string) *Database {
	if db, ok := dbs[db_name]; ok {
		return db
	}
	return nil
}
