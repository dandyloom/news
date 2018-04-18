package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/globalsign/mgo/bson"
	"encoding/json"
)

type User struct {
	// ID			bson.ObjectId	`bson:"_id,omitempty"`
	// ID			string	`json:"id"`
	Username	string
}





func main() {


	fmt.Println("Running server...")

	r := mux.NewRouter()

	// Routing Handlers
	r.HandleFunc("/posts", getPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", getPost).Methods("GET")
	r.HandleFunc("/posts", createPost).Methods("POST")
	r.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
	http.Handle("/", r)


	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}

func getPosts(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(Posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)

	// for  _, item := range Posts {
	// 	if item.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	} 
	// }

	// json.NewEncoder(w).Encode(&Post{})
}

func createPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var post Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	post.ID = bson.NewObjectId()

	if err := dao.Insert(post); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, post)
	// w.Header().Set("Content-Type", "application/json")

	// var Post Post
	// _ = json.NewDecoder(r.Body).Decode(&Post)
	// Post.ID = strconv.Itoa(rand.Intn(10000000))
	// Posts = append(Posts, Post)
	// json.NewEncoder(w).Encode(Post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {

}

func deletePost(w http.ResponseWriter, r *http.Request) {

}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}










