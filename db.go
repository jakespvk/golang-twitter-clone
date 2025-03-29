package main

import (
	"database/sql"
	"fmt"
)

func createTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS feed (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		message TEXT NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func seedData(db *sql.DB, data []Tweet) {
	query := `INSERT INTO feed (name, message) VALUES (?, ?);`
	tx, err := db.Begin()
	for _, each := range data {
		tx.Exec(query, each.Username, each.Message)
	}
	tx.Commit()
	if err != nil {
		panic(err)
	}
}

func insertTweet(db *sql.DB, tweet Tweet) {
	query := `INSERT INTO feed (name, message) VALUES (?, ?);`
	_, err := db.Exec(query, tweet.Username, tweet.Message)
	if err != nil {
		panic(err)
	}
}

func getTweets(db *sql.DB) []Tweet {
	var tweets []Tweet
	query := `SELECT * FROM feed;`
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var username string
		var message string
		err = rows.Scan(&id, &username, &message)
		if err == sql.ErrNoRows {
			return []Tweet{{00, "no content", "no content"}}
		} else if err != nil {
			panic(err)
		}
		tweets = append(tweets, Tweet{Id: id, Username: username, Message: message})
	}
	return tweets
}

func getTweetById(db *sql.DB, inputId int) Tweet {
	query := `SELECT * FROM feed WHERE id = ?;`
	row := db.QueryRow(query, inputId)

	var id int
	var username, message string

	err := row.Scan(&id, &username, &message)

	if err == sql.ErrNoRows {
		fmt.Println("No tweet found with that ID")
		return Tweet{00, "not found", "not found"}
	} else if err != nil {
		panic(err)
	}

	return Tweet{Id: id, Username: username, Message: message}
}

func deleteTweetById(db *sql.DB, tweet_id int) {
	query := `DELETE FROM feed WHERE id = ?;`
	_, err := db.Exec(query, tweet_id)
	if err != nil {
		panic(err)
	}
}

func getTweetsByUser(db *sql.DB, user string) []Tweet {
	var tweets []Tweet
	query := `SELECT * FROM feed WHERE name = ?;`
	rows, err := db.Query(query, user)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var username string
		var message string
		err = rows.Scan(&id, &username, &message)
		if err == sql.ErrNoRows {
			return []Tweet{{00, "no content", "no content"}}
		} else if err != nil {
			panic(err)
		}
		tweets = append(tweets, Tweet{Id: id, Username: username, Message: message})
	}
	return tweets
}
