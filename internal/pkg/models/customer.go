package models

import (
	"errors"
	"neversitup-test-template/internal/pkg/MongoDB"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time           `bson:"updated_at" json:"updated_at"`

	CustomerNo   string `bson:"customer_no" json:"customer_no"`
	CustomerName string `bson:"customer_name" json:"customer_name"`
	LineID       string `bson:"line_id" json:"line_id"`
}

func (r Customer) Col() string {
	return "customer"
}

func (r *Customer) Insert() error {
	r.CreatedAt = IncNowTime()
	r.UpdatedAt = IncNowTime()
	ins, err := MongoDB.Db.InsertOne(r.Col(), r)
	if err != nil {
		return err
	}
	if oid, ok := ins.InsertedID.(primitive.ObjectID); ok {
		r.ID = &oid
		return nil
	}
	return errors.New("cannot insert")
}

func (r *Customer) Update() error {
	r.UpdatedAt = IncNowTime()
	_, err := MongoDB.Db.UpdateOne(r.Col(), bson.M{"_id": r.ID}, bson.M{"$set": r})
	return err
}

func (r *Customer) FindByID(id string) error {
	oid, err := ConvertStringToObjectID(id)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id": oid,
	}
	err = MongoDB.Db.FindOne(r.Col(), filter, &r)
	return err
}

func (r Customer) FindAll() ([]Customer, error) {
	var result []Customer
	err := MongoDB.Db.Find(r.Col(), bson.M{}, &result)
	return result, err
}
