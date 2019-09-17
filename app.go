package main

import (
	"encoding/json"
	_ "encoding/json"
	"github.com/sielerjunjor/framework-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//	"google.golang.org/appengine"

	"github.com/gorilla/mux"
	"log"
	"net/http"

	"fmt"
	. "github.com/sielerjunjor/framework-api/config"
	. "github.com/sielerjunjor/framework-api/dao"
)

var config = Config{}
var dao = FrameworksDAO{}


//GET list of Frameworks
func AllFrameworksEndpoint(w http.ResponseWriter, r *http.Request){
	frameworks, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, frameworks)
}
// GET a Framework by its ID
func FindFrameworkById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	framework, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, framework)
}

// POST a new Framework
func CreateFrameworkEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var framework models.Framework
	if err := json.NewDecoder(r.Body).Decode(&framework); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	id := primitive.NewObjectID()
	framework.ID =&id
	if err := dao.Insert(framework); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, framework)
}

// PUT update an existing Framework
func UpdateFrameworkEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	defer r.Body.Close()
	var framework models.Framework
	if err := json.NewDecoder(r.Body).Decode(&framework); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err := dao.Update(params["id"] ,framework);if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing movie
func DeleteFrameworkEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_, err := dao.Delete(params["id"]);if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Collection = config.Collection
	dao.Connect()
}

func main() {

	fmt.Println("Starting Rest Endpoint")
	fmt.Println("Endpoint at localhost:3000/frameworks")

	r := mux.NewRouter()
	r.HandleFunc("/frameworks", AllFrameworksEndpoint).Methods("GET")
	r.HandleFunc("/frameworks/{id}", FindFrameworkById).Methods("GET")
	r.HandleFunc("/frameworks", CreateFrameworkEndPoint).Methods("POST")
	r.HandleFunc("/frameworks/{id}", UpdateFrameworkEndPoint).Methods("PUT")
	r.HandleFunc("/frameworks/{id}", DeleteFrameworkEndPoint).Methods("DELETE")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}

}
