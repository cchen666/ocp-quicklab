package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func Connect() (*sql.DB, error) {
	var conErr error
	c_string := "user:pass@tcp(10.72.94.119:3306)/db"
	db, conErr = sql.Open("mysql", c_string)
	if conErr != nil {
		return db, conErr
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return db, pingErr
	}
	return db, nil
}

func Test() string {
	albums, err := SearchByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	return albums[0].Title + albums[0].Artist + strconv.Itoa(int(albums[0].ID))
}

func SearchByArtist(name string) ([]Album, error) {
	var albums []Album
	db, _ = Connect()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}
