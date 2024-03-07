package models

import (
	"neversitup-test-template/internal/pkg/MongoDB"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApiLog struct {
	ID           *primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt    time.Time           `bson:"created_at" json:"created_at"`
	URL          string              `bson:"url" json:"url"`
	RequestData  string              `bson:"request_data" json:"request_data"`
	Status       int                 `bson:"status" json:"status"`
	ResponseData string              `bson:"response_data" json:"response_data"`
	IsSuccess    bool                `bson:"is_success" json:"is_success"`
}

func (r ApiLog) Col() string {
	return "api_logs"
}

func (r ApiLog) Insert() {
	_, _ = MongoDB.Db.InsertOne(r.Col(), r)
}
