package models

import (
	"neversitup-test-template/internal/pkg/MongoDB"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	ID           *primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt    time.Time           `bson:"created_at" json:"created_at"`
	FunctionName string              `bson:"function_name" json:"function_name"`
	Param        interface{}         `bson:"param" json:"param"`
	IsError      bool                `bson:"is_error" json:"is_error"`
	Error        error               `bson:"error" json:"error"`
}

func (r Log) Col() string {
	return "logs"
}

func (r Log) Insert() {
	_, _ = MongoDB.Db.InsertOne(r.Col(), r)
}
