package auth

import (
	"golang-crud/models"
	"golang-crud/utils"
	"log"
	"strings"
	"github.com/gofiber/fiber/v2"
)

func Protect(c *fiber.Ctx) error{
	var user models.User
	idToken:= strings.Replace(c.Get("Authorization"), "Bearer", "", 1)
	if idToken==""{
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "msg": "Unauthorized",
        })
    }
	client, err := App.Auth(utils.Ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	log.Println(idToken)
	token, err := client.VerifyIDToken(utils.Ctx, idToken)
	if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"msg": "Unauthorized",
				"err": err.Error(),
			})
	}
	
	result:=utils.DB.Model(&user).Preload("Take").Preload("Playlists").Preload("Review").Preload("Classes").Preload("Teacher").Preload("VideoInstance").Find(&user, "id=?", token.UID)

	//CHECK FOR ERROR
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
			"msg": result.Error.Error(),
			
		})
	}

	//CHECK FOR EXISTENCE
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "fail",
			"msg": "user not found",
		})
	}

	 c.Locals("user", user)
	 return c.Next()
	 
}