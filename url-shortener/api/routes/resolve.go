package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pratyush934/go-projects/api/database"
)

func ResolveURL(c *fiber.Ctx) error {

	url := c.Params("url")

	client := database.CreateClient(0)

	defer client.Close()

	result, err := client.Get(database.Ctx, url).Result()

	//if err != nil {
	//	if errors.Is(err, redis.Nil) {
	//
	//	}
	//}

	if err != nil && err.Error() == "redis: nil" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Value was not found in Redis",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	dbIncr := database.CreateClient(1)
	defer dbIncr.Close()

	_ = dbIncr.Incr(database.Ctx, "counter")

	return c.Redirect(result, 301)
}
