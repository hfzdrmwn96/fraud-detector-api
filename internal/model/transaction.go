package model

import "time"

// Transaction merepresentasikan data transaksi masuk
type Transaction struct {
	UserID    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
	Location  string    `json:"location"`
	DeviceID  string    `json:"device_id"`
	Recipient string    `json:"recipient"`
}
