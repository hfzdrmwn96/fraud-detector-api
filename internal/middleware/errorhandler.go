package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// ErrorHandler middleware untuk handle error dan panic
func ErrorHandler(c *fiber.Ctx) error {
	// Recover panic jika ada
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[PANIC] %v\n", r)
			// Bisa tambah stack trace juga jika perlu
		}
	}()

	err := c.Next() // lanjut ke handler berikutnya
	if err != nil {
		// Log error server-side
		log.Printf("[ERROR] %v\n", err)

		// Tentukan status code default
		code := fiber.StatusInternalServerError
		msg := "Internal Server Error"

		// Jika error adalah *fiber.Error, ambil status dan message nya
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			msg = e.Message
		}

		// Return response JSON error konsisten
		return c.Status(code).JSON(fiber.Map{
			"error": msg,
		})
	}

	return nil
}
