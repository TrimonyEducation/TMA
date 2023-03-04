package controllers

import (
	"golang-crud/models"
	"golang-crud/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	result:=utils.DB.Create(&models.User{Email: user.Email, Name: user.Name, SchoolGrade: user.SchoolGrade, SchoolLevel: user.SchoolLevel, IsTeacher: user.IsTeacher, CompletedOnboarding: user.CompletedOnboarding, EmailVerified: user.EmailVerified, ProfilePicture: user.ProfilePicture})
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
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "fail",
			"msg": "record not found",
		})
	}

	//RETURN FOUND EXERCISE
	return c.JSON(fiber.Map{
		"status":  "success",
		"data": user,    
})
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	//QUERY FOR PROBLEMS
    result:=utils.DB.Find(&users)

    //CHECK FOR ERROR
    if result.Error!= nil {
        return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error.Error(),
        })
    }

    //RETURN FOUND EXERCISES
    return c.JSON(fiber.Map{
        "status":  "success",
        "data": users,
		"result": len(users),
    })	
}

//UPDATE -----------------------------------------------------------------------------------

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
    user := new(models.User)
	s := new(models.User) 
    if err := c.BodyParser(user); err!=nil{
        return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": err.Error(),
        })
    }
    if id == "" {
		return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": "ID is missing",
        })
	}

	utils.DB.Find(&s, "id =?", id)
	if s.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
            "status":  "fail",
            "msg": "record not found",
        })
	}
	result := utils.DB.Where("id=?", id).Updates(&models.User{Email: user.Email, Name: user.Name, SchoolGrade: user.SchoolGrade, SchoolLevel: user.SchoolLevel, IsTeacher: user.IsTeacher, CompletedOnboarding: user.CompletedOnboarding, EmailVerified: user.EmailVerified})
	if result.Error!= nil {
		return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error.Error(),
        })
	}
    return c.Status(200).JSON(fiber.Map{
		"status": "success",
        "data": user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
    result := utils.DB.Delete(&models.User{}, "id =?", id)
    if result.Error!=nil {
        return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error.Error(),
        })
    }
    return c.Status(200).JSON(fiber.Map{
        "status": "success",
    })
}