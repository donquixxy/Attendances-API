package database

import (
	"context"
	"dg-test/ent"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func retryDatabase() *ent.Client {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		"ec2-54-252-184-68.ap-southeast-2.compute.amazonaws.com", "quixxy", "12345678", "dgtest", 8090,
	)

	client, err := ent.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("failed opening database %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed to connect to database : %v", err)
	}

	log.Println("Connected to default AWS database")

	return client
}

func NewDatabaseConnection() *ent.Client {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		"postgre", "quixxy", "12345678", "dgtest", 5432,
	)

	client, err := ent.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("failed opening database %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Printf("Failed to migrate local : %v", err)
		return retryDatabase()
	}

	log.Println("Connected to local database")

	return client
}
