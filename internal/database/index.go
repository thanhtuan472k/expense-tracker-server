package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// index ...
func index() {
	// index key
}

// func get index key
//
// Method private
//

// process ...
func process(col *mongo.Collection, indexes []mongo.IndexModel) {
	opts := options.CreateIndexes().SetMaxTime(time.Minute * 30)
	_, err := col.Indexes().CreateMany(context.Background(), indexes, opts)
	if err != nil {
		fmt.Printf("Index collection %s err: %v", col.Name(), err)
	}
}

// newIndex ...
func newIndex(key ...string) mongo.IndexModel {
	var doc bsonx.Doc
	for _, s := range key {
		e := bsonx.Elem{
			Key:   s,
			Value: bsonx.Int32(1),
		}
		if strings.HasPrefix(s, "-") {
			e = bsonx.Elem{
				Key:   strings.Replace(s, "-", "", 1),
				Value: bsonx.Int32(-1),
			}
		}
		doc = append(doc, e)
	}

	return mongo.IndexModel{Keys: doc}
}

// newUniqIndex ...
func newUniqIndex(key ...string) mongo.IndexModel {
	var doc bsonx.Doc
	for _, s := range key {
		e := bsonx.Elem{
			Key:   s,
			Value: bsonx.Int32(1),
		}
		if strings.HasPrefix(s, "-") {
			e = bsonx.Elem{
				Key:   strings.Replace(s, "-", "", 1),
				Value: bsonx.Int32(-1),
			}
		}
		doc = append(doc, e)
	}
	opt := options.Index().SetUnique(true)
	return mongo.IndexModel{Keys: doc, Options: opt}
}
