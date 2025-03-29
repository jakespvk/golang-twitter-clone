package main

import (
	"database/sql"
)

type Tweet struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type Server struct {
	DB *sql.DB
}
