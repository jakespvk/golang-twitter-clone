package main

import (
	"database/sql"
	"fmt"
)

func createTableFeed(db *sql.DB) {
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

func createTableReplies(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS replies (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tweetId INTEGER,
		username TEXT NOT NULL,
		message TEXT NOT NULL,
		FOREIGN KEY(tweetId) REFERENCES feed(id)
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

func insertReply(db *sql.DB, reply Reply) {
	query := `INSERT INTO replies (tweetID, username, message) VALUES (?, ?, ?);`
	_, err := db.Exec(query, reply.TweetID, reply.Username, reply.Message)
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
		tweets = append(tweets, Tweet{ID: id, Username: username, Message: message})
	}
	return tweets
}

func getTweetByID(db *sql.DB, inputID int) Tweet {
	query := `SELECT * FROM feed WHERE id = ?;`
	row := db.QueryRow(query, inputID)

	var id int
	var username, message string

	err := row.Scan(&id, &username, &message)

	if err == sql.ErrNoRows {
		fmt.Println("No tweet found with that ID")
		return Tweet{00, "not found", "not found"}
	} else if err != nil {
		panic(err)
	}

	return Tweet{ID: id, Username: username, Message: message}
}

func getTweetRepliesByTweetID(db *sql.DB, inputID int) []Reply {
	var replies []Reply
	query := `SELECT * FROM replies WHERE tweetId = ?`
	rows, err := db.Query(query, inputID)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id, tweetID int
		var username, message string
		err = rows.Scan(&id, &tweetID, &username, &message)
		if err == sql.ErrNoRows {
			return []Reply{{00, 00, "no content", "no content"}}
		} else if err != nil {
			panic(err)
		}
		replies = append(replies, Reply{ID: id, TweetID: tweetID, Username: username, Message: message})
	}
	return replies
}

func deleteTweetByID(db *sql.DB, tweetID int) {
	query := `DELETE FROM feed WHERE id = ?;`
	_, err := db.Exec(query, tweetID)
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
		tweets = append(tweets, Tweet{ID: id, Username: username, Message: message})
	}
	return tweets
}

func filterTweets(db *sql.DB, keyword string) []Tweet {
	var tweets []Tweet
	query := `SELECT * FROM feed WHERE message LIKE ?;`
	rows, err := db.Query(query, "%"+keyword+"%")
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
		tweets = append(tweets, Tweet{ID: id, Username: username, Message: message})
	}
	return tweets
}
