package main

import (
	"context"
	"dg-test/database"
	"dg-test/server"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.Println("Hello world")

	c := database.NewDatabaseConnection()
	defer c.Close()
	if err := c.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema %v", err)
	}

	// New echo instance
	serv := server.NewServer(c)

	go func() {
		if err := serv.E.Start(":9087"); err != nil {
			serv.E.Logger.Fatal("Server is shut down")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := serv.E.Shutdown(ctx); err != nil {
		serv.E.Logger.Fatal(err)
	}

}
