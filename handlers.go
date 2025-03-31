package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func (conn *Server) getChats(w http.ResponseWriter, r *http.Request) {
	chats := getTweets(conn.DB)

	jsonChats, err := json.Marshal(chats)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonChats)
}

func (conn *Server) postTweet(w http.ResponseWriter, r *http.Request) {
	var tweet Tweet

	if err := json.NewDecoder(r.Body).Decode(&tweet); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	insertTweet(conn.DB, tweet)
	w.WriteHeader(http.StatusOK)
}

func (conn *Server) getTweet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// TODO: add proper error handling for no 'id' param specified
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invlaid ID format", http.StatusBadRequest)
	}

	// TODO: add proper error handling if nothing is returned
	chat := getTweetById(conn.DB, id)

	jsonChats, err := json.Marshal(chat)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonChats)
}

func (conn *Server) deleteTweet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// TODO: add proper error handling for no 'id' param specified
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invlaid ID format", http.StatusBadRequest)
	}

	// TODO: add proper error handling if nothing is returned
	deleteTweetById(conn.DB, id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (conn *Server) getTweetsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	chats := getTweetsByUser(conn.DB, vars["user"])

	jsonChats, err := json.Marshal(chats)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonChats)
}

func (conn *Server) filterTweetsByKeyword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	chats := filterTweets(conn.DB, vars["keyword"])

	jsonChats, err := json.Marshal(chats)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonChats)
}
