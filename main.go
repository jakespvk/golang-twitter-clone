package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Connected to SQLite database!")

	conn := &Server{DB: db}

	r := mux.NewRouter()
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Allow all origins (use specific origins in production)
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)
	r.HandleFunc("/", helloWorld).Methods("GET")
	r.HandleFunc("/chats", conn.getChats).Methods("GET")
	r.HandleFunc("/chat", conn.postTweet).Methods("POST")
	r.HandleFunc("/chat/{id:[0-9]+}", conn.getTweet).Methods("GET")
	r.HandleFunc("/chat/{id:[0-9]+}", conn.deleteTweet).Methods("DELETE")
	r.HandleFunc("/chats/{user}", conn.getTweetsByUser).Methods("GET")
	r.HandleFunc("/chats/filter/{keyword}", conn.filterTweetsByKeyword).Methods("GET")
	http.Handle("/", r)

	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", corsHandler(r))
}
