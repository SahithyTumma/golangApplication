package handlers

import (
	"fmt"

	"github.com/SahithyTumma/docker/database"
	"github.com/SahithyTumma/docker/models"
	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	if len(facts) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No facts found",
		})
	}

	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

func ShowFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	database.DB.Db.Where("id = ?", id).First(&fact)

	if fact.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No fact found",
		})
	}

	return c.Status(200).JSON(fact)
}

func UpdateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	id := c.Params("id")

	// Parsing the request body
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}

	result := database.DB.Db.Model(&fact).Where("id = ?", id).Updates(fact)

	fmt.Println(fact)

	if result.Error != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(result.Error.Error())
	}

	return ShowFact(c)
}

func DeleteFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).Delete(&fact)

	if result.Error != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(result.Error.Error())
	}

	return ListFacts(c)
}
