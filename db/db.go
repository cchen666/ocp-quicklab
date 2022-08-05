package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() (*sql.DB, error) {
	var conErr error
	c_string := "user:pass@tcp(10.72.94.119:3305)/db"
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
	_, err := Connect()
	if err != nil {
		return err.Error()
	}
	return "succeeded"
}

func Show(name string) ([]Album, error) {
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