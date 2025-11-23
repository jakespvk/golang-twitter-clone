package main

import (
	"database/sql"
)

type Tweet struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type Reply struct {
	ID       int    `json:"id"`
	TweetID  int    `json:"tweetId"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type TweetReplyRequest struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type Server struct {
	DB *sql.DB
}
