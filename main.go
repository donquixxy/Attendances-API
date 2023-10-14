package main

import (
	"context"
	"dg-test/database"
	"dg-test/ent"
	"log"
)

func main() {
	log.Println("Hello world")

	c := database.NewDatabaseConnection()

	if err := c.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema %v", err)
	}

	TestCreateUser(context.Background(), c)
}

func TestCreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	result, err := client.User.Create().SetID("98821").SetEmail("testmail@mail").
		SetName("indra").Save(ctx)

	if err != nil {
		log.Printf("error ccreate user %v", err)
		return nil, err
	}

	log.Println("Result :", result)

	return result, nil
}
