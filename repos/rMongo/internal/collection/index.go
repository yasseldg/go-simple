package collection

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/yasseldg/go-simple/repos/rMongo/internal"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/sort"

	"github.com/yasseldg/go-simple/repos/rIndex"

	"github.com/yasseldg/go-simple/logs/sLog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *Base) CreateIndexes(ctx context.Context, indexes ...rIndex.Inter) error {
	if c.env == internal.Env_Read {
		return nil
	}

	err_msg := fmt.Sprintf("Mongo: %s.CreateIndex()", c.prefix)

	errs := make([]error, 0, len(indexes))
	for _, index := range indexes {

		fields, err := sort.Fields(index.Sort())
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: sort.Fields(): %s", err_msg, err))
			continue
		}

		errs = append(errs, c.createIndex(ctx, fields, index.Unique()))
	}

	return errors.Join(errs...)
}

// CreateIndex, create an index for a specific field in a collectionName
func (c *Base) createIndex(ctx context.Context, fields interface{}, unique bool) error {
	// 1. Lets define the keys for the index we want to create
	mod := mongo.IndexModel{
		Keys:    fields, // index in ascending order or -1 for descending order
		Options: options.Index().SetUnique(unique),
	}

	if ctx == nil {
		ctx = context.Background()
	}

	ctx, cancel := context.WithTimeout(ctx, (35 * time.Second))
	defer cancel()

	err_msg := fmt.Sprintf("Mongo: %s.CreateIndex()", c.prefix)
	// 4. Create a single index
	count := 0
	for count < 10 {
		index, err := c.coll.Indexes().CreateOne(ctx, mod)
		if err == nil {
			sLog.Info("Mongo: Index %s.%s created \n", c.prefix, index)
			return nil
		}

		sLog.Error("%s: %s", err_msg, err)

		time.Sleep(5 * time.Second)
		count++
	}

	return fmt.Errorf("%s failed", err_msg)
}
