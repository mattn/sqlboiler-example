package main

//go:generate sqlboiler sqlite3

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mattn/sqlboiler-example/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func doSomething() error {
	db, err := sql.Open("sqlite3", "models.sqlite")
	if err != nil {
		return err
	}

	// select Users
	users, err := models.Users().All(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Println(user.Name)
	}

	// begin transaction
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer tx.Rollback()

	// insert new User
	user := &models.User{}
	user.Name = "mattn"
	err = user.Insert(context.Background(), tx, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()

	user, err = models.Users().One(context.Background(), db)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func main() {
	err := doSomething()
	if err != nil {
		log.Fatal(err)
	}
}
