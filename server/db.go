package main

import (
	"database/sql"
	"log"
	"time"
)

func SetupDb(db *sql.DB) {
	const create string = `
		CREATE TABLE IF NOT EXISTS wisdoms (
			id INTEGER NOT NULL PRIMARY KEY,
			time DATETIME NOT NULL,
			description  TEXT UNIQUE,
			explanation TEXT
		);
	`

	_, err := db.Exec(create)
	if err != nil {
		log.Fatalf("failed in DB setup: %v", err)
	}

	for _, w := range BaseWisdomList {
		CreateWisdom(db, w)

	}

}

func GetWisdoms(db *sql.DB) []Wisdom {
	selectWisdoms := `SELECT description, explanation FROM wisdoms`

	rows, err := db.Query(selectWisdoms)
	if err != nil {
		log.Printf("failed get wisdoms: %v", err)
		return []Wisdom{}
	}

	wisdoms := []Wisdom{}

	for rows.Next() {
		w := Wisdom{}
		err := rows.Scan(&w.Description, &w.Explanation)
		if err != nil {
			log.Printf("failed parse wisdom: %v", err)
		}
		wisdoms = append(wisdoms, w)
		log.Println(wisdoms)
	}

	return wisdoms
}

func CreateWisdom(db *sql.DB, w Wisdom) {
	const insert string = `INSERT INTO wisdoms VALUES (NULL,?,?,?);`

	_, err := db.Exec(insert, time.Now().Unix(), w.Description, w.Explanation)
	if err != nil {
		log.Printf("failed to insert wisdom: %v", err)
	}
}
