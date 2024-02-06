package sMongo

import (
	"fmt"
	"time"

	"github.com/yasseldg/go-simple/configs/sEnv"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getConfig(env string) *mgm.Config {
	return &mgm.Config{
		CtxTimeout: time.Duration(sEnv.GetInt(fmt.Sprint("CTX_", env), 10)) * time.Second,
	}
}

func GetID(id interface{}) primitive.ObjectID {
	if strId, ok := id.(string); ok {
		objId, _ := primitive.ObjectIDFromHex(strId)
		return objId
	}
	return id.(primitive.ObjectID)
}

func GetIDs(objs []mgm.DefaultModel) []primitive.ObjectID {
	objIds := []primitive.ObjectID{}
	for _, obj := range objs {
		objIds = append(objIds, obj.ID)
	}
	return objIds
}
