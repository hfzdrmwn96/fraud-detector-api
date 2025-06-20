package service

import (
	"fraud-detector-api/internal/model"
)

func CalculateFraudScore(txn model.Transaction) float64 {
	score := 0.0

	if txn.Amount > 10000000 { // nominal > 10 juta
		score += 0.4
	}

	hour := txn.Timestamp.Hour()
	if hour < 7 || hour > 20 { // jam tidak biasa
		score += 0.2
	}

	// Dummy rule lokasi (bisa kembangkan nanti)
	if txn.Location != "Jakarta" {
		score += 0.3
	}

	// Rule device baru bisa ditambahkan nanti

	if score > 1.0 {
		score = 1.0
	}

	return score
}

func GetRiskLevel(score float64) string {
	switch {
	case score > 0.7:
		return "High"
	case score > 0.4:
		return "Medium"
	default:
		return "Low"
	}
}
