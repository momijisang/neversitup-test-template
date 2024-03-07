package models

import (
	"neversitup-test-template/internal/pkg/MongoDB"
	"neversitup-test-template/internal/pkg/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// create function convert []string to []primitive.ObjectID
func ConvertStringListToObjectIDs(ids []string) []primitive.ObjectID {
	var objectIDs []primitive.ObjectID
	for _, id := range ids {
		oid, err := primitive.ObjectIDFromHex(id)
		if err == nil {
			objectIDs = append(objectIDs, oid)
		}
	}
	return objectIDs
}

func ConvertStringToObjectID(id string) (*primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	return &oid, err
}

func AnyInsert(col string, r interface{}) error {
	_, err := MongoDB.Db.InsertOne(col, r)
	return err
}

func AnyUpdate(col string, findKey string, r interface{}) error {
	_, err := MongoDB.Db.UpdateOne(col, bson.M{"find_key": findKey}, bson.M{"$set": r})
	return err
}

func IsKeyFound(col string, findKey string) bool {
	count, err := MongoDB.Db.Count(col, bson.M{"find_key": findKey})
	return count > 0 && err == nil
}

func IncTime(v time.Time) time.Time {
	return v.Add(time.Duration(config.Config.Server.IncTimeBKK) * time.Hour)
}

func IncNowTime() time.Time {
	v := time.Now().Add(time.Duration(config.Config.Server.IncTimeBKK) * time.Hour)
	return v
}
