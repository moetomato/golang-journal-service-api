package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/moetomato/golang-journal-service-api/api"
)

func main() {
	var (
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASS")
		dbName = os.Getenv("DB_NAME")
		// [temporal] work on localhost とりあえず
		dbConn = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPass, dbName)
	)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("failed to connect DataBase")
		return
	}
	r := api.NewRouter(db)

	log.Println("journal api server starting at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
