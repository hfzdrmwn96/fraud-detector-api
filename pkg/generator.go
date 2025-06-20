package pkg

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Transaction struct {
	UserID    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
	Location  string    `json:"location"`
	DeviceID  string    `json:"device_id"`
	Recipient string    `json:"recipient"`
}

var locations = []string{"Jakarta", "Bandung", "Surabaya", "Tokyo", "New York"}
var deviceIDs = []string{"dev-a1", "dev-b2", "dev-c3", "dev-x1", "dev-y2"}
var recipients = []string{"R100", "R200", "R300", "R400", "R500"}

func randomTime() time.Time {
	now := time.Now()
	delta := rand.Intn(60 * 24 * 30) // up to 30 days
	return now.Add(time.Duration(-delta) * time.Minute)
}

func GenerateTransactions(n int, outputFile string) error {
	rand.Seed(time.Now().UnixNano())
	var transactions []Transaction

	for i := 0; i < n; i++ {
		txn := Transaction{
			UserID:    fmt.Sprintf("U%03d", rand.Intn(100)),
			Amount:    float64(rand.Intn(20000000)), // up to 20 juta
			Timestamp: randomTime(),
			Location:  locations[rand.Intn(len(locations))],
			DeviceID:  deviceIDs[rand.Intn(len(deviceIDs))],
			Recipient: recipients[rand.Intn(len(recipients))],
		}
		transactions = append(transactions, txn)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(transactions)
}
