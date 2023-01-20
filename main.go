package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
)

var client = redis.NewClient(&redis.Options{
	Addr: "redis.hop:6379",
	OnConnect: func(ctx context.Context, cn *redis.Conn) error {
		fmt.Println("Connected to Redis")
		return nil
	},
})

func main() {
	hostname, _ := os.Hostname()
	app := fiber.New(fiber.Config{})

	app.Get("/", func(c *fiber.Ctx) error {
		ctx := context.Background()
		var count = client.Incr(ctx, "counter")

		return c.SendString(fmt.Sprintf("%s: Page has %d visits!", hostname, count.Val()))
	})

	app.Listen(":3000")
}
