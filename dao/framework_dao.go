package dao

import (
	"context"
	"github.com/sielerjunjor/framework-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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
	log.Println("Connecting to ", m.Server)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(m.Server))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

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
	oid, _ := primitive.ObjectIDFromHex(id)
	framework.ID = &oid

	res := collection.FindOneAndUpdate(nil, bson.M{"_id": oid}, bson.M{"$set": &framework})
	return res.Err()
	//return err
}
