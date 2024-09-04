package rMongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type ObjectID struct {
	primitive.ObjectID
}

func GetID(id interface{}) ObjectID {
	if strId, ok := id.(string); ok {
		objId, _ := primitive.ObjectIDFromHex(strId)
		return ObjectID{objId}
	}
	return ObjectID{id.(primitive.ObjectID)}
}

func GetIDs(objs []DefaultModel) []ObjectID {
	objIds := []ObjectID{}
	for _, obj := range objs {
		objIds = append(objIds, ObjectID{obj.ID})
	}
	return objIds
}
