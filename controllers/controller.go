package controllers

import (
	model "BLOG_API/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

var dbName string = "blogs"
var colName string = "blog"

func init() {
	err := godotenv.Load("./.env")
	checkNilErr(err)
	var uri = os.Getenv("MONGO_URI")
	clientOption := options.Client().ApplyURI(uri)
	// Connect to DB
	client, err := mongo.Connect(context.TODO(), clientOption)
	checkNilErr(err)
	fmt.Println("Mongo Connection Succesful")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}

func insertBlog(blog model.Blog) {
	inserted, err := collection.InsertOne(context.Background(), blog)
	checkNilErr(err)
	fmt.Println("Movie got inserted", inserted)
}

func getAllBlogs() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	checkNilErr(err)
	var blogs []primitive.M
	for cur.Next(context.Background()) {
		var blog bson.M
		err := cur.Decode(&blog)
		checkNilErr(err)
		blogs = append(blogs, blog)
	}
	defer cur.Close(context.Background())
	return blogs
}

func getABlog(blogID string) model.Blog {
	id, _ := primitive.ObjectIDFromHex(blogID)
	filter := bson.M{"_id": id}
	var blog model.Blog
	if err := collection.FindOne(context.Background(), filter).Decode(&blog); err != nil {
		log.Fatal(err)
	}
	return blog
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var blog model.Blog
	_ = json.NewDecoder(r.Body).Decode(&blog)
	insertBlog(blog)
	json.NewEncoder(w).Encode(blog)
}
func GetABlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	aBlog := getABlog(params["id"])
	json.NewEncoder(w).Encode(aBlog)
}
func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allBlogs := getAllBlogs()
	json.NewEncoder(w).Encode(allBlogs)
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
