package controllers

import (
	"golang-crud/models"
	"golang-crud/utils"

	"github.com/gofiber/fiber/v2"
)

//CREATE ----------------------------------------------------------------
func CreatePlaylist(c *fiber.Ctx) error {
	playlist:=new(models.Playlist)
	if err := c.BodyParser(playlist); err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
            "msg": err.Error(),
            "error": err.Error,
		})
	}
	result:=utils.DB.Create(&playlist)
	if result.Error!=nil{
		return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error.Error(),
            "error": result.Error,
        })
	}
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data": playlist,
	})
}

//READ ----------------------------------------------------------------

func GetPlaylist(c *fiber.Ctx) error {
	id := c.Params("id")

	//INIT VARS
	var playlist models.Playlist
	
	//QUERY FOR PROBLEM
   result:=utils.DB.First(&playlist, "id=?", id)

   //CHECK FOR ERROR
    if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "fail",
            "msg": result.Error.Error(),
		})
	}

	//CHECK FOR EXISTENCE
	if playlist.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
            "status":  "fail",
            "msg": "exercise not found",
        })
	}

	//RETURN FOUND EXERCISE
    return c.JSON(fiber.Map{
        "status":  "success",
        "data": playlist,    
    })
}

func GetAllPlaylists(c *fiber.Ctx) error{
	var playlist []models.Playlist
	//QUERY FOR PROBLEMS
    result:=utils.DB.Find(&playlist)

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
        "data": playlist,
    })
}

func UpdatePlaylist(c *fiber.Ctx) error {
	id := c.Params("id")
    playlist := new(models.Playlist)
	s := new(models.Playlist) 
    if err := c.BodyParser(playlist); err!=nil{
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
	if s.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
            "status":  "fail",
            "msg": "exercise not found",
        })
	}
	result := utils.DB.Where("id=?", id).Updates(&playlist)
	if result.Error!= nil {
		return c.Status(400).JSON(fiber.Map{
            "status":  "fail",
            "msg": result.Error.Error(),
        })
	}
    return c.Status(200).JSON(fiber.Map{
		"status": "success",
        "data":   playlist,
	})
}

func DeletePlaylist(c *fiber.Ctx) error {
	id := c.Params("id")
    result := utils.DB.Where("id=?", id).Delete(&models.Playlist{})
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