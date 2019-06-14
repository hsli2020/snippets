package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type ShopLocation struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"shopid"`
	Name     string        `bson:"name"          json:"name"`
	Location GeoJson       `bson:"location"      json:"location"`
}

type GeoJson struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}

func main() {
	cluster := "localhost" // mongodb host

	session, err := mgo.Dial(cluster)	// connect to mongo
	if err != nil {
		log.Fatal("could not connect to db: ", err)
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// search criteria
	long := 139.701642 
	lat := 35.690647
	scope := 3000 // max distance in metres

	var results []ShopLocation // to hold the results

	c := session.DB("test").C("shops")	// query the database
	err = c.Find(bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{long, lat},
				},
				"$maxDistance": scope,
			},
		},
	}).All(&results)
	if err != nil {
		panic(err)
	}

	// convert it to JSON so it can be displayed
	formatter := json.MarshalIndent
	response, err := formatter(results, " ", "   ")

	fmt.Println(string(response))
}
