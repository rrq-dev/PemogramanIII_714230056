package handler

import (
	"inibackend/repository"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	return c.SendString("Welcome to the homepage!")
}

func GetAllMahasiswa(c *fiber.Ctx) error {
	data, err := repository.GetAllMahasiswa(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(data)
}

func GetMahasiswaByNPM(c *fiber.Ctx) error {
	npmStr := c.Params("npm")

	npm, err := strconv.Atoi(npmStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid NPM",
		})
	}
	mhs, err := repository.GetMahasiswaByNPM(c.Context(), npm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	if mhs.NPM == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Mahasiswa not found",
		})
	}
    // Log before returning response to client
    log.Printf("GetMahasiswaByNPM: %v\n", mhs)
	return c.JSON(mhs)
}
