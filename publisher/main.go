package main

import (
	"encoding/json"
	"go-redis-pub-sub-example/model"
	"net/http"

	config "go-redis-pub-sub-example/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/publish", func(c *fiber.Ctx) error {
		newUser := model.User{}

		if err := c.BodyParser(&newUser); err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}

		payload, err := json.Marshal(newUser)
		if err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}

		if err := config.RedisClient.Publish(config.Ctx, "send-user-data", payload).Err(); err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}

		return c.JSON(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Message published successfully",
		})
	})

	app.Listen(":3000")
}
