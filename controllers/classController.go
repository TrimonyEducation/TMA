package controllers

import (
	"golang-crud/models"
	"golang-crud/utils"

	"github.com/gofiber/fiber/v2"
)

//CREATE ----------------------------------------------------------------
func CreateClass(c *fiber.Ctx) error {
	class:=new(models.Class)
	if err := c.BodyParser(class); err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
            "msg": err.Error(),
            "error": err.Error,
		})
	}
	result:=utils.DB.Create(&class)
	if result.Error!=nil{
		return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error.Error(),
            "error": result.Error,
        })
	}
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data": class,
	})
}

//READ ----------------------------------------------------------------

func GetClass(c *fiber.Ctx) error {
	//PARAMS
    id := c.Params("id")

	//INIT VARS
	var class models.Class
	
	//QUERY FOR PROBLEM
   result:=utils.DB.Model(&class).Preload("Teacher").First(&class, "id=?", id)

   //CHECK FOR ERROR
    if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
            "msg": result.Error.Error(),
            "error": result.Error,
		})
	}

	//CHECK FOR EXISTENCE
	if class.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
            "status":  "fail",
            "msg": "exercise not found",
        })
	}

	//RETURN FOUND EXERCISE
    return c.JSON(fiber.Map{
        "status":  "success",
        "exercise": class,    
    })
}

func GetAllClasses(c *fiber.Ctx) error {
	var classes []models.Class
	//QUERY FOR PROBLEMS
    result:=utils.DB.Find(&classes)

    //CHECK FOR ERROR
    if result.Error!= nil {
        return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error.Error(),
            "error": result.Error,
        })
    }

    //RETURN FOUND EXERCISES
    return c.JSON(fiber.Map{
        "status":  "success",
        "data": classes,
    })
}

//UPDATE ----------------------------------------------------------------

func UpdateClass(c *fiber.Ctx) error {
	//PARAMS
    id := c.Params("id")
    class := new(models.Class)
	s := new(models.Class) 
    if err := c.BodyParser(class); err!=nil{
        return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": err.Error(),
            "error": err.Error,
        })
    }
    if id == "" {
		return c.Status(400).JSON(fiber.Map{
            "status": "fail",
            "msg": "ID is missing",
        })
	}

	utils.DB.Find(&s, "id =?", id)
	if s.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
            "status":  "fail",
            "msg": "exercise not found",
        })
	}
	result := utils.DB.Where("id=?", id).Updates(&class)
	if result.Error!= nil {
		return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error.Error(),
            "error": result.Error,
        })
	}
    return c.Status(200).JSON(fiber.Map{
		"status": "success",
        "data": class,
	})
}

//DELETE ----------------------------------------------------------------

func DeleteClass(c *fiber.Ctx) error {
    //PARAMS
    id := c.Params("id")

    //DELETE
    result := utils.DB.Where("id=?", id).Delete(&models.Class{})
    if result.Error!=nil{
        return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error.Error(),
            "error": result.Error,
        })
    }
    return c.Status(200).JSON(fiber.Map{
        "status": "success",
    })
}