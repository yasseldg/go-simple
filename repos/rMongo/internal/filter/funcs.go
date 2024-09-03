package filter

import (
	"fmt"

	"github.com/yasseldg/go-simple/repos/rFilter"

	"go.mongodb.org/mongo-driver/bson"
)

func Fields(filter rFilter.Inter) (bson.D, error) {
	f, ok := filter.Oper().(*Filter)
	if !ok {
		return nil, fmt.Errorf("filter is not rMongo.Filter")
	}
	return f.fields, nil
}