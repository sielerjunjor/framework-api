package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/sielerjunjor/framework-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FrameworksDAO struct {
	Server   string
	Database string
	Collection string
}

var collection *mongo.Collection

const (
	COLLECTION = "frameworks"
)

// Establish a connection to database
func (m *FrameworksDAO) Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.Server))
	if err != nil {
		log.Print("first")
		log.Fatal(err)
	}

	collection = client.Database(m.Database).Collection(m.Collection)
}


// Find list of frameworks
func (m *FrameworksDAO) FindAll() ([]models.Framework, error) {
	var frameworks []models.Framework
	cur, err := collection.Find(nil, bson.M{})
	err = cur.All(nil, &frameworks)
	return frameworks, err
}

// Find a Framework by its id
func (m *FrameworksDAO) FindById(id string) (models.Framework, error) {
	var framework models.Framework
	oid, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(nil, bson.M{"_id":oid }).Decode(&framework)
	return framework, err
}

// Insert a framework into database
func (m *FrameworksDAO) Insert(framework models.Framework) error {
	_, err := collection.InsertOne(nil, &framework)
	return err
}

//Delete
func (m *FrameworksDAO) Delete(id string)  (bool, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	res, err := collection.DeleteOne(nil, bson.M{"_id": oid})
	return res.DeletedCount == 1, err
}

// Update an existing framework
func (m *FrameworksDAO) Update(id string, framework models.Framework) (error) {
	fmt.Print(id)

	oid, _ := primitive.ObjectIDFromHex(id)
	framework.ID = &oid
	fmt.Print(framework)

	filter := bson.M{"_id": oid}
	fmt.Println("filter: %d\n", filter)
	update := bson.M{"$set": &framework}

	res := collection.FindOneAndUpdate(nil, filter, update)



	return res.Err()
	//return err
}
