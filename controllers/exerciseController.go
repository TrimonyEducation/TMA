package controllers

import (
	"golang-crud/models"
	"golang-crud/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//CREATE ----------------------------------------------------------------

func CreateExercise(c *fiber.Ctx) error {
	//PARSE BODY
	exercise := new(models.Exercise)
	x := new(ExerciseWithID)
	if err := c.BodyParser(exercise); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    err.Error(),
		})
	}

	// INSERT INTO DB
	result := utils.DB.Create(&exercise)
	if x.VideoID != "" {
		result := utils.DB.Exec("INSERT INTO video_exercise VALUES(" + exercise.ID.String() + "," + x.VideoID + ");")
		if result.Error != nil {
			return c.Status(400).JSON(fiber.Map{
				"status": "fail",
				"msg":    result.Error.Error(),
			})
		}
	}
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    result.Error.Error(),
		})
	}

	//RETURN CREATED
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   exercise,
	})
}

//READ ----------------------------------------------------------------

func GetExercise(c *fiber.Ctx) error {
	//GET ID
	id := c.Params("id")

	//INIT VARS
	var exercise models.Exercise
	//QUERY FOR EXERCISE AND RELATED VIDEOS, PROBLEMS AND TAKES
	result := utils.DB.Model(&exercise).Preload("Takes").Preload("Problems").Preload("Videos").Find(&exercise, "id=?", id)

	//CHECK FOR ERROR
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    result.Error.Error(),
		})
	}

	//CHECK FOR EXISTENCE
	if exercise.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "fail",
			"msg":    "record not found",
		})
	}

	//RETURN FOUND EXERCISE
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   exercise,
	})
}

func GetAllExercises(c *fiber.Ctx) error {
	//INIT VARS
	var exercises []models.Exercise

	//QUERY FOR EXERCISES
	result := utils.DB.Find(&exercises)
	//CHECK FOR ERROR
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    result.Error.Error(),
		})
	}

	//RETURN FOUND EXERCISES
	return c.JSON(fiber.Map{
		"status":  "success",
		"data":    exercises,
		"results": len(exercises),
	})
}

//UPDATE ----------------------------------------------------------------

// EXTENDS EXERCISE STRUCT TO RECIEVE ID FROM BODY(PARSER)
type ExerciseWithID struct {
	models.Exercise
	VideoID string
}

func UpdateExercise(c *fiber.Ctx) error {
	//IS GOING TO BE UPDATED
	exercise := new(models.Exercise)

	//FOR BODYPARSER TO RECIEVE VideoID FROM BODY
	x := new(ExerciseWithID)

	s := new(models.Exercise)
	//CHECK FOR ERROR
	if err := c.BodyParser(x); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    err.Error(),
		})
	}
	if err := c.BodyParser(exercise); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    err.Error(),
		})
	}
	//GET ID
	id := c.Params("id")

	//CHECK FOR ID
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    "ID incorrect",
		})
	}

	//IF VideoID IS PRESENT INSERT EXERCISE ID AND PRESENT ID INTO JOIN TABLE
	if x.VideoID != "" {
		result := utils.DB.Exec("INSERT INTO video_exercise VALUES(" + id + "," + x.VideoID + ");")
		if result.Error != nil {
			return c.Status(400).JSON(fiber.Map{
				"status": "fail",
				"msg":    result.Error.Error(),
			})
		}
	}

	utils.DB.Find(&s, "id =?", id)
	if s.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "fail",
			"msg":    "record not found",
		})
	}

	//FIND EXERCISE BY ID AND UPDATE IT
	result := utils.DB.Where("id = ?", id).Updates(&exercise)

	//CHECK FOR ERROR
	if result.Error != nil {
		return c.Status(203).JSON(fiber.Map{
			"status": "fail",
			"msg":    result.Error.Error(),
		})
	}

	//RETURN UPDATED EXERCISE
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   exercise,
	})
}

//DELETE ----------------------------------------------------------------

func DeleteExercise(c *fiber.Ctx) error {
	id := c.Params("id")
	result := utils.DB.Delete(&models.Exercise{}, "id =?", id)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    result.Error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
	})
}
