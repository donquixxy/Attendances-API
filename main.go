package main

import (
	"context"
	"dg-test/database"
	"dg-test/repository"
	"dg-test/server"
	"dg-test/service"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/go-co-op/gocron"
	"gopkg.in/gomail.v2"
)

func main() {

	// Wait for postgres local to ready
	log.Println("Waiting for local postgres to be ready")
	time.Sleep(time.Second * 10)

	c := database.NewDatabaseConnection()
	defer c.Close()

	dial := gomail.NewDialer("smtp.gmail.com", 587, "agusariputra70@gmail.com", "plzw ffgv imri pgad")

	msgger := gomail.NewMessage()

	mailRepo := repository.NewMailRepository(msgger, dial)

	mailService := service.NewCronJOBService(mailRepo, repository.NewUserRepository(c))

	tz, _ := time.LoadLocation("Asia/Makassar")
	sched := gocron.NewScheduler(tz)

	// Cronjob scheduler
	go func() {
		sched.Every(1).Day().At("08:55").Do(func() {
			mailService.SendCheckinReminder(context.Background())
		})

		sched.Every(1).Day().At("13:27").Do(func() {
			mailService.SendCheckoutReminder(context.Background())
		})

		sched.StartAsync()
	}()

	// New echo instance
	serv := server.NewServer(c)

	// Echo Server
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
