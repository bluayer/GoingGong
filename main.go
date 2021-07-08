package main

import (
	"context"
	"os"

	"log"

	"entgo.io/ent/dialect"
	"github.com/bluayer/GoingGong/ent"
	"github.com/bluayer/GoingGong/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv("DB_URL")
	port := os.Getenv("PORT")
	client, err := ent.Open(dialect.MySQL, dbUrl)

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	e := router.Create(client)

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
