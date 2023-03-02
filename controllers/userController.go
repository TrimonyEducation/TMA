package controllers

import (
	"golang-crud/models"
	"golang-crud/utils"

	"github.com/gofiber/fiber/v2"
)

// CREATE ----------------------------------------------------------------
func CreateUser(c *fiber.Ctx) error{
	//PARSE BODY
	user := new(models.User)
	if err := c.BodyParser(user); err!= nil {
        return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
			"msg": err.Error(),
		})
    }

	// INSERT INTO DB
	result:=utils.DB.Create(&user)
	if result.Error!=nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
            "msg": result.Error.Error(),
            "error": result.Error,
		})
	}

	//RETURN CREATED
	return c.JSON(fiber.Map{
		"status":  "success",
        "data": user,
	})
}

// READ -----------------------------------------------------------------------------------

func GetUser(c *fiber.Ctx) error {
	//GET ID
	id := c.Params("id")

	//INIT VARS
	var user models.User
	//QUERY FOR EXERCISE AND RELATED VIDEOS, PROBLEMS AND TAKES
	result:=utils.DB.Model(&user).Preload("Take").Preload("Playlists").Preload("Review").Preload("Classes").Preload("Teacher").Preload("VideoInstance").Find(&user, "id=?", id)

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
			"msg": "exercise not found",
		})
	}

	//RETURN FOUND EXERCISE
	return c.JSON(fiber.Map{
		"status":  "success",
		"data": user,    
})
}