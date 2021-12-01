package db

import (
	"eth_service/internal/config"
	"fmt"
	"github.com/rubenv/sql-migrate"
	"log"
)

var migrations = &migrate.PackrMigrationSource{Box: Migrations}

func MigrateUp(cfg config.Config) {
	log.Println("before migrate")
	n, err := migrate.Exec(cfg.DB().RawDB(), "postgres", migrations, migrate.Up)
	if err != nil {
		log.Println("migrate probs")
		log.Fatal(err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}

//func MigrateDown()
