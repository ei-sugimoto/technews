package infra

import (
	"log"
	"os"

	"github.com/ei-sugimoto/technews/api/ent"
	_ "github.com/go-sql-driver/mysql"
)

func NewClient() *ent.Client {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	database := os.Getenv("DB_NAME")
	if host == "" || port == "" || user == "" || password == "" || database == "" {
		log.Fatalf("DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME must be set")
	}
	client, err := ent.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	return client
}
