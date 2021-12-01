package db

import (
	"eth_service/internal/config"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/jmoiron/sqlx"
	"gitlab.com/distributed_lab/kit/kv"
	"log"
	"time"
)

//go:generate packr2 clean
//go:generate packr2
var Migrations = packr.New("migrations", "./migrations")

func InitDB() *sqlx.DB {
	dbCridentials, err := config.ParseDB(kv.MustGetStringMap(kv.MustFromEnv(), "db"))
	if err != nil{
		log.Fatal(err)
	}

	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		dbCridentials.Host,
		dbCridentials.Port,
		dbCridentials.DB,
		dbCridentials.User,
		dbCridentials.Password,
	)
	db, err := sqlx.Open("postgres", pgConnString) // init db
	if err != nil {
		log.Println("Open")
		log.Fatal(err)
	}

	retryPing(db)

	return db
}

func retryPing(db *sqlx.DB) {
	// retry logic
	retries := 5
	for retries > 0 {
		pingErr := db.Ping() // really connect to database
		if pingErr != nil {
			retries -= 1
			fmt.Println("Retries left ", retries)
			log.Println(pingErr)
			time.Sleep(time.Second)
		} else {
			log.Printf("Left Ping")
			break
		}
	}
}
