package helpers

import "github.com/gofiber/fiber/v2"

func ResponseJson(c *fiber.Ctx, err string, msg string, code int, status string, data any) error {
	return c.Status(code).JSON(fiber.Map{
		"code":    code,
		"status":  status,
		"message": msg,
		"error":   err,
		"data":    data,
	})
}

func ResponseToken(c *fiber.Ctx, code int, status string, token string) error {
	return c.Status(code).JSON(fiber.Map{
		"code":   code,
		"status": status,
		"token":  token,
	})
}
