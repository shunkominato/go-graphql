package main

import (
	"context"
	"fmt"
	"log"

	"server/ent"

	"entgo.io/ent/entc/integration/ent"
	_ "github.com/lib/pq"
)

func main() {
	// opts := []ent.Option{
	// 	ent.("./model"),  // 出力先を./model配下に指定
	// 	ent.Package("model"),   // パッケージ名をmodelに指定
	// 	ent.Header(schema.Header), // ヘッダーを指定
	// }
	// if err := entc.Generate("./schema", opts...); err != nil {
	// 	log.Fatalf("running ent codegen: %v", err)
	// }
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		"rails_api-db-1", "5432", "postgres", "rails_sample", "postgres"))

		log.Print()
		log.Print(err)
		aa, er := client.Todo.Query().All(context.Background())
    log.Print(aa)
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