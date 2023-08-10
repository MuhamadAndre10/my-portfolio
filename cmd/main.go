package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/MuhamadAndre10/portfolio/config/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	_, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		AllowCredentials: true,
		MaxAge:           86400, //24 hour
	}))

	// listen from the diferent goroutine
	go func() {
		if err := app.Listen(":8000"); err != nil {
			log.Fatal(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)                    // create a channel a signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // when an interupt ot terminate signal is sent, notify the channel

}
