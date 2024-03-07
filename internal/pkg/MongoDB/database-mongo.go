package MongoDB

import (
	"context"
	"fmt"
	"neversitup-test-template/internal/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db MongoDB

//var DbLine MongoDB

type MongoDB struct {
	client *mongo.Client
	ctx    context.Context
	cancel context.CancelFunc
	dbName string
	err    error
}

func (m *MongoDB) close() {
	defer m.cancel()
	defer func() {
		if err := m.client.Disconnect(m.ctx); err != nil {
			panic(err)
		}
	}()
}

func (m *MongoDB) Connect(configMongo config.DatabaseConfiguration) MongoDB {
	uri := fmt.Sprintf(configMongo.Host, configMongo.Username, configMongo.Password)

	//ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	ctx := context.Background()
	var cancel context.CancelFunc
	clientOption := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOption)
	return MongoDB{client, ctx, cancel, configMongo.Dbname, err}
}

func (m *MongoDB) InsertOne(col string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := m.client.Database(m.dbName).Collection(col)
	result, err := collection.InsertOne(m.ctx, doc)
	return result, err
}

func (m *MongoDB) InsertMany(col string, docs []interface{}) (*mongo.InsertManyResult, error) {
	collection := m.client.Database(m.dbName).Collection(col)
	result, err := collection.InsertMany(m.ctx, docs)
	return result, err
}

func (m *MongoDB) UpdateOne(col string, filter, data interface{}) (result *mongo.UpdateResult, err error) {
	collection := m.client.Database(m.dbName).Collection(col)
	result, err = collection.UpdateOne(m.ctx, filter, data)
	return
}

func (m *MongoDB) UpdateOneAndLogBeforeChange(col string, filter, data interface{}) (result *mongo.UpdateResult, err error) {
	//find data before change
	var beforeChange map[string]interface{}
	err = m.FindOne(col, filter, &beforeChange)
	if err != nil {
		return nil, err
	}
	beforeChange["collection"] = col
	_, _ = m.InsertOne("log_before_change", beforeChange)
	result, err = m.UpdateOne(col, filter, data)
	return
}

func (m *MongoDB) UpdateMany(col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {
	collection := m.client.Database(m.dbName).Collection(col)
	result, err = collection.UpdateMany(m.ctx, filter, update)
	return
}

func (m *MongoDB) Find(col string, query, results interface{}, opts ...*options.FindOptions) (err error) {
	collection := m.client.Database(m.dbName).Collection(col)
	result, err := collection.Find(m.ctx, query, opts...)
	if err = result.All(m.ctx, results); err != nil {
	}
	return
}

func (m *MongoDB) FindOne(col string, query, results interface{}, opts ...*options.FindOneOptions) (err error) {
	collection := m.client.Database(m.dbName).Collection(col)
	result := collection.FindOne(m.ctx, query, opts...)
	if err = result.Decode(results); err != nil {
	}
	return
}

func (m *MongoDB) FindOneAndUpdate(col string, query, update, results interface{}, opts ...*options.FindOneAndUpdateOptions) (err error) {
	collection := m.client.Database(m.dbName).Collection(col)
	result := collection.FindOneAndUpdate(m.ctx, query, update, opts...)
	if err = result.Decode(results); err != nil {
	}
	return
}

func (m *MongoDB) Aggregate(col string, query, results interface{}) (err error) {
	collection := m.client.Database(m.dbName).Collection(col)
	result, err := collection.Aggregate(m.ctx, query)
	if err == nil {
		if err = result.All(m.ctx, results); err != nil {
		}
	}
	return
}

func (m *MongoDB) Count(col string, query interface{}) (result int64, err error) {
	collection := m.client.Database(m.dbName).Collection(col)
	result, err = collection.CountDocuments(m.ctx, query)
	return
}
