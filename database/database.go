package database

import (
	"dg-test/ent"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewDatabaseConnection() *ent.Client {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		"ec2-54-252-184-68.ap-southeast-2.compute.amazonaws.com", "quixxy", "12345678", "dgtest", 8090,
	)

	client, err := ent.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("failed opening database %v", err)
	}
	log.Println("Connected to database")
	return client
}
