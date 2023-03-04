package controllers

import (
	"golang-crud/models"
	"golang-crud/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//CREATE ----------------------------------------------------------------
func CreateClass(c *fiber.Ctx) error {
	class:=new(models.Class)
    x := new(UpdateWithID)
	if err := c.BodyParser(class); err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
            "msg": err.Error(),
		})
	}
   
	result:=utils.DB.Create(&class)
    
    if x.UserID != "" {
		result := utils.DB.Exec("INSERT INTO user_classes VALUES("+x.UserID+","+class.ID.String()+");")
		if result.Error!= nil {
            return c.Status(fiber.StatusBadRequest).JSON(result.Error)
        }
	}
	if result.Error!=nil{
		return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error.Error(),
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
		})
	}

	//CHECK FOR EXISTENCE
	if class.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
            "status":  "fail",
            "msg": "record not found",
        })
	}

	//RETURN FOUND EXERCISE
    return c.Status(200).JSON(fiber.Map{
        "status":  "success",
        "exercise": class,    
    })
}

func GetAllClasses(c *fiber.Ctx) error {
	var classes []models.Class
	//QUERY FOR CLASSES
    result:=utils.DB.Find(&classes)

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
        "data": classes,
        "result": len(classes),
    })
}

//UPDATE ----------------------------------------------------------------

type UpdateWithID struct {
	models.Class
	UserID string
}

func UpdateClass(c *fiber.Ctx) error {
	//PARAMS
    id := c.Params("id")
    class := new(models.Class)
	s := new(models.Class) 
    x := new(UpdateWithID)
    if err := c.BodyParser(class); err!=nil{
        return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": err.Error(),
        })
    }
    if id == "" {
		return c.Status(400).JSON(fiber.Map{
            "status": "fail",
            "msg": "ID is missing",
        })
	}

    if x.UserID != "" {
		result := utils.DB.Exec("INSERT INTO user_classes VALUES("+x.UserID+","+id+");")
		if result.Error!= nil {
            return c.Status(fiber.StatusBadRequest).JSON(result.Error)
        }
	}

	utils.DB.Find(&s, "id =?", id)
	if s.ID == uuid.Nil {
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
        })
    }
    return c.Status(200).JSON(fiber.Map{
        "status": "success",
    })
}