package handlers

import (
	"github.com/SahithyTumma/docker/database"
	"github.com/SahithyTumma/docker/models"
	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	var facts []models.Fact
	if err := database.DB.Db.Find(&facts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve facts: " + err.Error(),
		})
	}

	if len(facts) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No facts found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body: " + err.Error(),
		})
	}

	if err := database.DB.Db.Create(fact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create fact: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fact)
}

func ShowFact(c *fiber.Ctx) error {
	var fact models.Fact
	id := c.Params("id")

	if err := database.DB.Db.Where("id = ?", id).First(&fact).Error; err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No fact found with the given ID",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve fact: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fact)
}

func UpdateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	id := c.Params("id")

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body: " + err.Error(),
		})
	}

	result := database.DB.Db.Model(&models.Fact{}).Where("id = ?", id).Updates(fact)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update fact: " + result.Error.Error(),
		})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No fact found to update with the given ID",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Fact updated successfully",
	})
}

func DeleteFact(c *fiber.Ctx) error {
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).Delete(&models.Fact{})
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete fact: " + result.Error.Error(),
		})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No fact found to delete with the given ID",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Fact deleted successfully",
	})
}
