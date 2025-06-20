package handler

import (
	"fraud-detector-api/internal/model"
	"fraud-detector-api/internal/service"
	"fraud-detector-api/pkg"

	"github.com/gofiber/fiber/v2"
)

func AnalyzeTransaction(c *fiber.Ctx) error {
	var txn model.Transaction
	if err := c.BodyParser(&txn); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON format")
	}
	if txn.Amount <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Amount must be greater than zero")
	}
	if txn.UserID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "UserID is required")
	}

	score := service.CalculateFraudScore(txn)
	risk := service.GetRiskLevel(score)

	pkg.FileLogger.Printf(
		"[UserID: %s] Amount: %.2f | Location: %s | Device: %s | Score: %.2f | Risk: %s",
		txn.UserID, txn.Amount, txn.Location, txn.DeviceID, score, risk,
	)

	return c.JSON(fiber.Map{
		"fraud_score": score,
		"risk_level":  risk,
	})
}
