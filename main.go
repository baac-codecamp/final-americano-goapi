package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/app/final-americano-goapi/helper"
	"github.com/app/final-americano-goapi/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type ResMsg struct {
	Response_status string  `json:"response_status"`
	Response_message string   `json:"response_message"`
	Response_data interface{}  `json:"response_data"`
}

func getNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// we created Book array
	var newses []models.News

	//Connection mongoDB with helper class
	collection := helper.ConnectDB()

	// bson.M{},  we passed empty filter. So we want to get all data.
	options := options.Find()
	sort := bson.D{}
	sort = append(sort, bson.E{"createdAt", -1})

	options.SetSort(sort)

	filter := bson.M{}
	cur, err := collection.Find(context.TODO(), filter, options)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var news models.News
		// & character returns the memory address of the following variable.
		err := cur.Decode(&news) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array	//line 20
		newses = append(newses, news)
	}

	var Response ResMsg
	Response.Response_status = "success"
	Response.Response_message = "Get News Success!"
	Response.Response_data = newses
	
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	
	json.NewEncoder(w).Encode(Response) // encode similar to serialize process.
}

func getNewByID(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var news models.News
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := helper.ConnectDB()

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&news)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	var Response ResMsg
	Response.Response_status = "success"
	Response.Response_message = "Get news by id success!"
	Response.Response_data = news

	json.NewEncoder(w).Encode(Response)
}

func createNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var news models.News

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&news)

	// connect db
	collection := helper.ConnectDB()

	// insert our book model.
	news.CreatedAt = time.Now().Add(time.Duration(15) * time.Minute)
	news.UpdatedAt = time.Now().Add(time.Duration(15) * time.Minute)
	result, err := collection.InsertOne(context.TODO(), news)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/go/getNews", getNews).Methods("GET")
	r.HandleFunc("/go/getNewsById/{id}", getNewByID).Methods("GET")
	//r.HandleFunc("/admin/addNews", createNews).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
	
}
