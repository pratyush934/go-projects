package routes

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pratyush934/go-projects/api/database"
	"github.com/pratyush934/go-projects/api/helpers"
	"os"
	"strconv"
	"time"
)

type request struct {
	URL         string    `json:"url"`
	CustomShort string    `json:"custom-short"`
	Expiry      time.Time `json:"expiry"`
}

type response struct {
	URL            string    `json:"url"`
	CustomShort    string    `json:"custom-short"`
	Expiry         time.Time `json:"expiry"`
	XRateRemaining int       `json:"rate_limit"`
	XRateReset     int       `json:"rate-reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	//implement rate limiting
	client := database.CreateClient(1)

	defer client.Close()

	result, err := client.Get(database.Ctx, c.IP()).Result()

	if err != nil && err.Error() == "redis: nil" {
		_ = client.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		result, _ = client.Get(database.Ctx, c.IP()).Result()
		valInt, _ := strconv.Atoi(result)

		if valInt <= 0 {
			limit, _ := client.TTL(database.Ctx, c.IP()).Result()

			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":      "Rate limit exceed",
				"rate_limit": limit / time.Nanosecond / time.Minute,
			})
		}
	}

	//check if the input if an actual URL
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL",
		})
	}

	//check for domain error

	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "There is an error in RemoveDomainError",
		})
	}

	//enforce https, ssl
	body.URL = helpers.EnforceHTTP(body.URL)

	var id string

	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	r := database.CreateClient(0)
	defer r.Close()

	result, _ = r.Get(database.Ctx, id).Result()

	if result != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Custom short URL already exists",
		})
	}

	if body.Expiry.IsZero() {
		body.Expiry = time.Now().Add(24 * time.Hour)
	}

	err = r.Set(database.Ctx, id, body.URL, body.Expiry.Sub(time.Now())).Err()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to connect to server",
		})
	}

	resp := response{
		URL:            body.URL,
		CustomShort:    "",
		Expiry:         body.Expiry,
		XRateRemaining: 10,
		XRateReset:     30,
	}
	client.Decr(database.Ctx, c.IP())

	val, _ := r.Get(database.Ctx, c.IP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	duration, err := r.TTL(database.Ctx, c.IP()).Result()

	resp.XRateReset = int(duration / time.Nanosecond / time.Minute)

	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id
	return c.Status(fiber.StatusOK).JSON(resp)
}
