package api

import (
	"fmt"
	"strconv"
	"time"

	"backend-config.Cache/config"

	"github.com/gofiber/fiber/v2"
)

func SetKey(c *fiber.Ctx) error {
	key := c.Query("key")
	val := c.Query("val")
	expiry := c.Query("expiry")

	if key != "" && val != "" && expiry != "" {
		msg := ""
		exp, err := strconv.Atoi(expiry)
		if err != nil {
			fmt.Println("failed to get expiry, setting up default expiry of 5 Sec")
			exp = 120
			msg = " Unable to get expiry; defaulting to 5 sec"
		}
		expiryDuration := time.Duration(exp) * time.Second
		ok := config.Lru.Put(key, val, expiryDuration)
		if ok {
			c.Status(200).JSON(
				&fiber.Map{
					"status":  true,
					"message": "success" + msg,
				},
			)

		} else {
			c.Status(500).JSON(
				&fiber.Map{
					"status":  false,
					"message": "failed to add key in cache",
				},
			)
		}
	} else {
		c.Status(400).JSON(
			&fiber.Map{
				"status":  false,
				"message": "Bad Request!",
			},
		)
	}
	return nil
}
