package controllers

import (
	"golang-crud/models"
	"golang-crud/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//CREATE ----------------------------------------------------------------
func CreateTake(c *fiber.Ctx) error {
	take:=new(models.Take)
	if err := c.BodyParser(take); err!=nil{
		return c.Status(401).JSON(fiber.Map{
			"status":  "fail",
			"msg": err.Error(),
		})
	}
	result:=utils.DB.Create(&take)
	if result.Error!=nil{
		return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error,
        })
	}
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   take,
	})
}

//READ ----------------------------------------------------------------

func GetTake(c *fiber.Ctx) error {
	//PARAMS
    id := c.Params("id")

	//INIT VARS
	var take models.Take
	
	//QUERY FOR TAKE
   result:=utils.DB.First(&take, "id=?", id)

   //CHECK FOR ERROR
    if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
            "msg": result.Error.Error(),
		})
	}

	//CHECK FOR EXISTENCE
	if take.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
            "status":  "fail",
            "msg": "record not found",
        })
	}

	//RETURN FOUND TAKE
    return c.JSON(fiber.Map{
        "status":  "success",
        "data": take,    
    })
}

func GetAllTakes(c *fiber.Ctx) error {
	var takes []models.Take
	//QUERY FOR TAKES
    result:=utils.DB.Find(&takes)

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
        "data": takes,
    })
}

//UPDATE ----------------------------------------------------------------

func UpdateTake(c *fiber.Ctx) error {
	//PARAMS
    id := c.Params("id")
    take := new(models.Take)
	s := new(models.Take) 
    if err := c.BodyParser(take); err!=nil{
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
	result := utils.DB.Where("id=?", id).Updates(&take)
	if result.Error!= nil {
		return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error.Error(),
        })
	}
    return c.Status(200).JSON(fiber.Map{
		"status": "success",
        "data": take,
	})
}

//DELETE ----------------------------------------------------------------

func DeleteTake(c *fiber.Ctx) error {
    //PARAMS
    id := c.Params("id")

    //DELETE
    result := utils.DB.Where("id=?", id).Delete(&models.Take{})
    if result.Error!=nil{
        return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error.Error(),
        })
    }
    return c.Status(200).JSON(fiber.Map{
        "status": "success",
    })
}