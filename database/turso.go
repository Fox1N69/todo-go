package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/libsql/go-libsql"
	"github.com/sirupsen/logrus"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var AToken = "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MDk1NTE4NjAsImlkIjoiOGU5MTRkMDEtZGExYS0xMWVlLWI1ZDQtZjY5YmQyNmE1NGJhIn0.yH-QfskJZoc3Jzq7me6y6MKtPlTcv_3IP7zfFeJbdirUO4O5xyx-0MI1tjwXQ9jAdwTA_0q3nHJA5NJoYyevBw"

type Users struct {
	ID int
	username string
	password string
}

func InitDB() {
	url := "libsql://site-sb-fox1n69.turso.io?authToken=AToken"

	db, err := sql.Open("libsql", url)
	if err != nil {
		logrus.Fatal(os.Stderr, "failed connect to database", url)
		os.Exit(1)
	}
	logrus.Infoln("Databases connection...")

	defer db.Close()
}

func queryDB(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []Users

	for rows.Next() {
		var user Users

		if err := rows.Scan(user.ID, user.username, user.password); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}

		users = append(users, user)
		fmt.Println(user.ID, user.username, user.password)
	}
}

func syncDatabase(connector *libsql.Connector) {
	if err := connector.Sync(); err != nil {
		logrus.Fatal("Error syncing database:", err)
	}
}
