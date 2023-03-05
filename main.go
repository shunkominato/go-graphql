package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"story.com/story/app/ent"
)

func main_() {

	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		"rails_api-db-1", "5432", "postgres", "rails_sample", "postgres"))

		log.Print()
		log.Print(err)
		aa, er := client.Todo.Query().All(context.Background())
    log.Print(aa[0].Todo)
    log.Print(er)
	// // e := echo.New()
	// // e.GET("/", func(c echo.Context) error {
	// // 	db.Query("SELECT * FROM employee")

	// // 	return c.String(http.StatusOK, "Hello, Echo World!!")
	// // })
	// // e.Logger.Fatal(e.Start(":8080"))
	// defer client.Close()
	
	log.Print("ent sample done.")
}