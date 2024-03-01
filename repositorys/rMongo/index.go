package rMongo

import (
	"context"
	"time"

	"github.com/yasseldg/go-simple/logs/sLog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Index struct {
	Fields interface{}
	Unique bool
}
type Indexes []Index

func (c Collection) createIndexes(indexes Indexes) {
	if c.environment == EnvironmentRead {
		return
	}

	for _, index := range indexes {
		c.createIndex(index.Fields, index.Unique)
	}
}

// CreateIndex, create an index for a specific field in a collectionName
func (c *Collection) createIndex(fields interface{}, unique bool) {
	if c.environment == EnvironmentRead {
		return
	}

	// 1. Lets define the keys for the index we want to create
	mod := mongo.IndexModel{
		Keys:    fields, // index in ascending order or -1 for descending order
		Options: options.Index().SetUnique(unique),
	}

	ctx, cancel := context.WithTimeout(context.Background(), (35 * time.Second))
	defer cancel()

	// 4. Create a single index
	count := 0
	for count < 15 {
		index, err := c.collection.Indexes().CreateOne(ctx, mod)
		if err == nil {
			sLog.Info("sMongo: Index %s%s created \n", c.prefix, index)
			return
		}

		sLog.Error("sMongo: %sCreateIndex(): %s", c.prefix, err)

		time.Sleep(5 * time.Second)
		count++
	}
}
