package dao

import (
	"log"

	"github.com/ledex/framework-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FrameworksDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "frameworks"
)

// Establish a connection to database
func (m *FrameworksDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}


// Find list of frameworks
func (m *FrameworksDAO) FindAll() ([]models.Framework, error) {
	var frameworks []models.Framework
	err := db.C(COLLECTION).Find(bson.M{}).All(&frameworks)
	return frameworks, err
}

// Find a Framework by its id
func (m *FrameworksDAO) FindById(id string) (models.Framework, error) {
	var framework models.Framework
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&framework)
	return framework, err
}

// Insert a movie into database
func (m *FrameworksDAO) Insert(framework models.Framework) error {
	err := db.C(COLLECTION).Insert(&framework)
	return err
}

// Delete an existing movie
//func (m *FrameworksDAO) Delete(framework models.Framework) error {
//	err := db.C(COLLECTION).Remove(&framework)
//	return err
//}

//Delete
func (m *FrameworksDAO) Delete(id string)  error {
	var framework models.Framework
	xrr := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&framework)
	err := db.C(COLLECTION).Remove(xrr)
	return err
}

// Update an existing movie
func (m *FrameworksDAO) Update(id string, framework models.Framework) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &framework)
	return err
}
