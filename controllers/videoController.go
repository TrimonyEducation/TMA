package controllers

import (
	"golang-crud/models"
	"golang-crud/utils"

	"github.com/gofiber/fiber/v2"
)

// CREATE --------------------------------------------------------------------------------

func CreateVideo(c *fiber.Ctx) error {
        video:= new(models.Video)
        if err := c.BodyParser(video); err != nil {
            return c.Status(503).SendString(err.Error())
        }
		result := utils.DB.Create(&video)
		if result.Error != nil {
			return c.Status(400).JSON(fiber.Map{
				"status": "fail",
				"err": result.Error.Error(),
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"status": "success",
			"data": video,
		})
}

// READ -----------------------------------------------------------------------------------

func GetAllVideos(c *fiber.Ctx) error {
	var video []models.Video
	result := utils.DB.Find(&video)
	if  result.RowsAffected == 0 {
        return c.SendStatus(404)
    }
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"result": len(video),
		"data": video,
	})
}

func GetVideo(c *fiber.Ctx) error {
	//GET ID
    id := c.Params("id")
	
	//INIT VARS
	var video models.Video

	//QUERY FOR EXERCISE AND RELATED VIDEOS, PROBLEMS AND TAKES
   result:=utils.DB.Model(&video).Preload("Exercises").Preload("Exercises.Problems").Find(&video, "id=?", id)

   //CHECK FOR ERROR
    if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(result.Error)
	}

	//CHECK FOR EXISTENCE
	if video.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status":  "failed",
            "message": "video not found",
        })
	}

	//RETURN FOUND EXERCISE
    return c.JSON(fiber.Map{
        "status":  "success",
        "exercise": video,    
    })
}

// UPDATE -----------------------------------------------------------------------------------

type VideoWithID struct {
	models.Video
	ExerciseID string
}

func UpdateVideo(c *fiber.Ctx) error {
	//TO UPDATE A VIDEO BY ID
	video:=new(models.Video)

	//FOR BODYPARSER TO RECIEVE ID FROM BODY
	x := new(VideoWithID)

	//TO FIND A VIDEO BY ID
	s:=new(models.Video)

	//CHECK FOR ERROR
    if err := c.BodyParser(x); err!= nil {
        return err
    }
	if err := c.BodyParser(video); err!= nil {
        return err
    }

	//GET ID
    id := c.Params("id")

	//CHECK FOR ID
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":"fail",
			"msg": "ID incorrect",
		})
	}

	//IF "ExerciseID" IS PRESENT INSERT VIDEO ID AND "ExerciseID" INTO JOIN TABLE
	if x.ExerciseID != "" {
		result := utils.DB.Exec("INSERT INTO video_exercise VALUES("+id+","+x.ExerciseID+");")
		if result.Error!= nil {
            return c.Status(fiber.StatusBadRequest).JSON(result.Error)
        }
	}

	//FIND A VIDEO BY ID AND CHECK IF IT EXISTS
	utils.DB.Find(&s, "id =?", id)
	if s.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status":  "fail",
            "message": "video not found",
        })
	}

	//FIND VIDEO BY ID AND UPDATE IT
	result := utils.DB.Where("id = ?", id).Updates(&video)	

	//CHECK FOR ERROR
	if result.Error!=nil {
        return c.Status(203).JSON(fiber.Map{
			"status":"fail",
			"msg":result.Error.Error(),
		})
    }

	//RETURN UPDATED VIDEO
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
        "data": video,
	})

}

// DELETE -----------------------------------------------------------------------------------

func DeleteVideo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var video models.Video
	if id == 0 {
		return c.Status(400).JSON(fiber.Map{
            "status":"fail",
            "msg": "ID incorrect",
        })
	}

    if err!= nil {	
		return c.Status(400).JSON(fiber.Map{
            "status":"fail",
            "msg": "ID missing",
        })
	}
	result := utils.DB.Delete(&video, "id =?", id)
	if result.Error!= nil {
		return c.Status(400).JSON(fiber.Map{
            "status":"fail",
            "msg": result.Error.Error(),
        })
	}
	if result.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
            "status":"fail",
            "msg": "Couldn't delete video",
        })
	}

	return c.Status(200).JSON(fiber.Map{
        "status": "success",
		"msg": "Video deleted",
    })
}