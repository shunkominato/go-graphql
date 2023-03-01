package main

import (
	"context"
	"fmt"
	"log"
	"server/ent"

	_ "github.com/lib/pq"
)

type Todo struct {
  ID        uint      // カラム名は`id`
  Todo      string    // カラム名は`name`
}

func main() {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		"rails_api-db-1", "5432", "postgres", "rails_sample", "postgres"))

		log.Print(err)
		aa, er := client.Todo.Query().All(context.Background())
    log.Print(aa)
    log.Print(er)
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	db.Query("SELECT * FROM employee")

	// 	return c.String(http.StatusOK, "Hello, Echo World!!")
	// })
	// e.Logger.Fatal(e.Start(":8080"))
	defer client.Close()
	
	log.Print("ent sample done.")
}