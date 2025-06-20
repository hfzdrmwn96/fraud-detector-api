package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Struktur transaksi sesuai format
type Transaction struct {
	UserID    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
	Location  string    `json:"location"`
	DeviceID  string    `json:"device_id"`
	Recipient string    `json:"recipient"`
}

// Struktur respon dari API
type Response struct {
	FraudScore float64 `json:"fraud_score"`
	RiskLevel  string  `json:"risk_level"`
}

func main() {
	file, err := os.ReadFile("transactions.json")
	if err != nil {
		log.Fatalf("Gagal membaca file JSON: %v", err)
	}

	var transactions []Transaction
	if err := json.Unmarshal(file, &transactions); err != nil {
		log.Fatalf("Gagal parse JSON: %v", err)
	}

	client := &http.Client{}

	csvFile, err := os.Create("results.csv")
	if err != nil {
		log.Fatalf("Gagal buat file CSV: %v", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Header kolom
	writer.Write([]string{"Index", "UserID", "Amount", "Timestamp", "Location", "DeviceID", "Recipient", "FraudScore", "RiskLevel"})

	for i, txn := range transactions {
		body, _ := json.Marshal(txn)
		req, err := http.NewRequest("POST", "http://localhost:8080/analyze", bytes.NewBuffer(body))
		if err != nil {
			log.Printf("Gagal buat request untuk transaksi ke-%d: %v", i, err)
			continue
		}

		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Gagal kirim request ke-%d: %v", i, err)
			continue
		}
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)

		var res Response
		if err := json.Unmarshal(respBody, &res); err != nil {
			log.Printf("Gagal parse respon ke-%d: %v", i, err)
			continue
		}

		writer.Write([]string{
			fmt.Sprintf("%d", i+1),
			txn.UserID,
			fmt.Sprintf("%.0f", txn.Amount),
			txn.Timestamp.Format(time.RFC3339),
			txn.Location,
			txn.DeviceID,
			txn.Recipient,
			fmt.Sprintf("%.2f", res.FraudScore),
			res.RiskLevel,
		})

		fmt.Printf("[%03d] User: %-6s | Amount: Rp%-10.0f | Risk: %-6s | Score: %.2f\n",
			i+1, txn.UserID, txn.Amount, res.RiskLevel, res.FraudScore)

		time.Sleep(100 * time.Millisecond) // jeda agar tidak terlalu cepat
	}
}
