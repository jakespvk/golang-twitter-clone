package main

import (
	// "context"
	"database/sql"
	"fmt"
	"net/http"

	// "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"
)

var adapter *httpadapter.HandlerAdapter

func main() {

	dbPath, err := setupDatabase()
	if err != nil {
		fmt.Println("Error setting up database:", err)
		return
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	fmt.Println("Connected to SQLite database!")

	conn := &Server{DB: db}

	router := mux.NewRouter()

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Allow all origins (use specific origins in production)
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	router.HandleFunc("/", helloWorld).Methods("GET")
	router.HandleFunc("/chats", conn.getChats).Methods("GET")
	router.HandleFunc("/chat", conn.postTweet).Methods("POST")
	router.HandleFunc("/chat/{id:[0-9]+}", conn.getTweet).Methods("GET")
	router.HandleFunc("/chat/{id:[0-9]+}", conn.deleteTweet).Methods("DELETE")
	router.HandleFunc("/chats/{user}", conn.getTweetsByUser).Methods("GET")
	router.HandleFunc("/chats/filter/{keyword}", conn.filterTweetsByKeyword).Methods("GET")
	http.Handle("/", router)

	adapter = httpadapter.New(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Fix path for AWS Lambda function URLs
		if rawPath := r.Header.Get("X-Forwarded-Path"); rawPath != "" {
			r.URL.Path = rawPath
		}

		router.Use(corsHandler)
		router.ServeHTTP(w, r)
	}))
	lambda.Start(adapter.ProxyWithContext)
}
