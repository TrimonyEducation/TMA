package controllers

import (
	"golang-crud/models"
	"golang-crud/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//CREATE ----------------------------------------------------------------

func CreateProblem(c *fiber.Ctx) error {
	problem := new(models.Problem)
	if err := c.BodyParser(problem); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    err.Error(),
		})
	}
	result := utils.DB.Create(&problem)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    result.Error.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   problem,
	})
}

//READ ------------------------------------------------------------------------------

func GetProblem(c *fiber.Ctx) error {
	//PARAMS
	id := c.Params("id")

	//INIT VARS
	var problem models.Problem

	//QUERY FOR PROBLEM
	result := utils.DB.First(&problem, "id=?", id)

	//CHECK FOR ERROR
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    result.Error.Error(),
		})
	}

	//CHECK FOR EXISTENCE
	if problem.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "fail",
			"msg":    "record not found",
		})
	}

	//RETURN FOUND EXERCISE
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   problem,
	})
}

func GetAllProblems(c *fiber.Ctx) error {
	var problems []models.Problem
	//QUERY FOR PROBLEMS
	result := utils.DB.Find(&problems)

	//CHECK FOR ERROR
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    result.Error.Error(),
		})
	}

	//RETURN FOUND EXERCISES
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   problems,
		"result": len(problems),
	})
}

//UPDATE -----------------------------------------------------------------------------------

func UpdateProblem(c *fiber.Ctx) error {

	//PARAMS
	id := c.Params("id")
	problem := new(models.Problem)
	s := new(models.Problem)
	if err := c.BodyParser(problem); err != nil {
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
	result := utils.DB.Where("id=?", id).Updates(&problem)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "fail",
			"msg":    result.Error.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   problem,
	})
}

//DELETE -----------------------------------------------------------------------------------

func DeleteProblem(c *fiber.Ctx) error {
	//PARAMS
	id := c.Params("id")

	//DELETE
	result := utils.DB.Where("id=?", id).Delete(&models.Problem{})
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
