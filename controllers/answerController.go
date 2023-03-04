package controllers

import (
	"golang-crud/models"
	"golang-crud/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateAnswer(c *fiber.Ctx) error {
	answer := new(models.Answers)
	if err := c.BodyParser(answer); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    err.Error(),
		})
	}
	result := utils.DB.Create(&answer)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    result.Error.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   answer,
	})
}

func GetAnswer(c *fiber.Ctx) error {
	id := c.Params("id")

	//INIT VARS
	var answer models.Answers

	//QUERY FOR answer
	result := utils.DB.Find(&answer, "id=?", id)

	//CHECK FOR ERROR
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    result.Error.Error(),
		})
	}

	//CHECK FOR EXISTENCE
	if answer.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "fail",
			"msg":    "record not found",
		})
	}

	//RETURN FOUND answer
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   answer,
	})
}

func GetAllAnswers(c *fiber.Ctx) error {
	var answers []models.Answers
	//QUERY FOR answerS
	result := utils.DB.Find(&answers)
	//CHECK FOR ERROR
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    result.Error.Error(),
		})
	}

	//RETURN FOUND answerS
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   answers,
		"result": len(answers),
	})
}

func UpdateAnswer(c *fiber.Ctx) error {
	id := c.Params("id")
	answer := new(models.Answers)
	s := new(models.Answers)
	if err := c.BodyParser(answer); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    err.Error(),
		})
	}
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    "ID is missing",
		})
	}

	utils.DB.Find(&s, "id =?", id)
	if s.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "fail",
			"msg":    "record not found",
		})
	}
	result := utils.DB.Where("id=?", id).Updates(&answer)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    result.Error.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   answer,
	})
}

func DeleteAnswer(c *fiber.Ctx) error {
	id := c.Params("id")
	result := utils.DB.Where("id=?", id).Delete(&models.Answers{})
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    result.Error.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
	})
}
