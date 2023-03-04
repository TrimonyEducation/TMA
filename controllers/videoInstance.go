package controllers

import (
	"golang-crud/models"
	"golang-crud/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//CREATE ----------------------------------------------------------------

func CreateVideoInstance(c *fiber.Ctx) error {
	//PARSE BODY
	videoInstance := new(models.VideoInstance)
	if err := c.BodyParser(videoInstance); err!= nil {
        return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
			"msg": err.Error(),
		})
    }

	// INSERT INTO DB
	result:=utils.DB.Create(&videoInstance)
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
        "data": videoInstance,
	})
}

//READ ----------------------------------------------------------------

func GetVideoInstance(c *fiber.Ctx) error {
    //GET ID
    id := c.Params("id")

	//INIT VARS
	var videoInstance models.VideoInstance
	//QUERY FOR EXERCISE AND RELATED VIDEOS, PROBLEMS AND TAKES
   result:=utils.DB.First(&videoInstance, "id=?", id)

   //CHECK FOR ERROR
    if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
            "msg": result.Error.Error(),
            "error": result.Error,
		})
	}

	//CHECK FOR EXISTENCE
	if videoInstance.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
            "status":  "fail",
            "msg": "exercise not found",
        })
	}

	//RETURN FOUND EXERCISE
    return c.JSON(fiber.Map{
        "status":  "success",
        "data": videoInstance,    
    })
}

func GetAllVideoInstances(c *fiber.Ctx) error {
	//INIT VARS
	var videoInstances []models.VideoInstance

	//QUERY FOR EXERCISES
    result := utils.DB.Find(&videoInstances)
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
        "data": videoInstances,
		"results": len(videoInstances), 
    })
}

//UPDATE ----------------------------------------------------------------

func UpdateVideoInstance(c *fiber.Ctx) error {
	id := c.Params("id")
    videoInstance := new(models.VideoInstance)
	s := new(models.VideoInstance) 
    if err := c.BodyParser(videoInstance); err!=nil{
        return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": err.Error(),
            "error": err.Error,
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
	result := utils.DB.Where("id=?", id).Updates(&videoInstance)
	if result.Error!= nil {
		return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error.Error(),
        })
	}
    return c.Status(200).JSON(fiber.Map{
		"status": "success",
        "data":   videoInstance,
	})
}

//DELETE ----------------------------------------------------------------

func DeleteVideoInstance(c *fiber.Ctx) error {
	//PARAMS
    id := c.Params("id")

    //DELETE
    result := utils.DB.Where("id=?", id).Delete(&models.VideoInstance{})
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