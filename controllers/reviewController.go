package controllers

import (
	"golang-crud/models"
	"golang-crud/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//CREATE ----------------------------------------------------------------
func CreateReview(c *fiber.Ctx) error {
	review:=new(models.Review)
	if err := c.BodyParser(review); err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
            "msg": err.Error(),
		})
	}
	result:=utils.DB.Create(&review)
	if result.Error!=nil{
		return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error.Error(),
        })
	}
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data": review,
	})
}

//READ ----------------------------------------------------------------

func GetReview(c *fiber.Ctx) error {
	id := c.Params("id")

	//INIT VARS
	var review models.Review
	
	//QUERY FOR PLAYLIST 
   result:=utils.DB.Model(&review).Preload("Videos").Find(&review, "id=?", id)

   //CHECK FOR ERROR
    if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
            "msg": result.Error.Error(),
		})
	}

	//CHECK FOR EXISTENCE
	if review.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
            "status":  "fail",
            "msg": "record not found",
        })
	}

	//RETURN FOUND PLAYLIST
    return c.JSON(fiber.Map{
        "status":  "success",
        "data": review,    
    })
}

func GetAllReviews(c *fiber.Ctx) error{
	var review []models.Review
	//QUERY FOR PLAYLISTS
    result:=utils.DB.Find(&review)
    //CHECK FOR ERROR
    if result.Error!= nil {
        return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error.Error(),
        })
    }

    //RETURN FOUND PLAYLISTS
    return c.JSON(fiber.Map{
        "status":  "success",
        "data": review,
        "result": len(review),
    })
}

func UpdateReview(c *fiber.Ctx) error {
	id := c.Params("id")
    review := new(models.Review)
	s := new(models.Review) 
    if err := c.BodyParser(review); err!=nil{
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
	result := utils.DB.Where("id=?", id).Updates(&review)
	if result.Error!= nil {
		return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error.Error(),
        })
	}
    return c.Status(200).JSON(fiber.Map{
		"status": "success",
        "data":   review,
	})
}

func DeleteReview(c *fiber.Ctx) error {
	id := c.Params("id")
    result := utils.DB.Where("id=?", id).Delete(&models.Review{})
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