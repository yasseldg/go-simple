package sort

import (
	"fmt"

	"github.com/yasseldg/go-simple/repos/rSort"

	"go.mongodb.org/mongo-driver/bson"
)

func Fields(sort rSort.Inter) (bson.D, error) {
	s, ok := sort.Oper().(*Sort)
	if !ok {
		return nil, fmt.Errorf("sort is not rMongo.Sort")
	}
	return s.fields, nil
}
