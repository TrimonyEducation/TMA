package auth

import (
	"golang-crud/models"

	"github.com/gofiber/fiber/v2"
)

func AdminOnly(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	if user == nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":"fail",
			"error":"unauthorized",
		})
    }
	if user.User_Role != "admin"{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "status":"fail",
            "error":"unauthorized",
        })
	} 
	return c.Next()
}

func TeacherAndAdminOnly(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	if user == nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":"fail",
			"error":"unauthorized",
		})
    }
	if user.User_Role != "teacher" && user.User_Role != "admin"{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "status":"fail",
            "error":"unauthorized",
        })
	} 
	return c.Next()
}