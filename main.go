package main

import (
	"database/sql"
	"fmt"
	"net/http"

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
	r.HandleFunc("/", helloWorld).Methods("GET")
	r.HandleFunc("/chats", conn.getChats).Methods("GET")
	r.HandleFunc("/chat", conn.postTweet).Methods("POST")
	r.HandleFunc("/chat/{id:[0-9]+}", conn.getTweet).Methods("GET")
	r.HandleFunc("/chat/{id:[0-9]+}", conn.deleteTweet).Methods("DELETE")
	r.HandleFunc("/chats/{user}", conn.getTweetsByUser).Methods("GET")
	http.Handle("/", r)

	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
