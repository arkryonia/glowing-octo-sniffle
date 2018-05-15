package db

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

var (
	// Session enregistre en mémoire une session mongo
	Session *mgo.Session

	// Mongo enregistre en mémoire les informations liées à la connexion
	Mongo *mgo.DialInfo
)

const (
	// MongoDBUrl est l'adresse par défaut de la base de données
	MongoDBUrl = "mongodb://localhost:20017/blog-db"
)

// Connect connecte à la base de données mongodb
func Connect() {
	uri := os.Getenv("MONGODB_URL")

	if len(uri) == 0 {
		uri = MongoDBUrl
	}

	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)

	if err != nil {
		fmt.Printf("Can't connect to mongo, go error%v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", uri)
	Session = s
	Mongo = mongo
}
