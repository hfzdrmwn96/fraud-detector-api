# 🛡️ Fraud Detector API – BI-OJK Hackathon 2025

Sistem RESTful API berbasis Golang yang menganalisis transaksi perbankan dan mengembalikan skor risiko fraud berdasarkan rule heuristik. Dirancang untuk digunakan oleh institusi keuangan sebagai alat bantu deteksi dini potensi kecurangan transaksi.

---

## 🚀 Fitur Utama

- ✅ Analisis transaksi real-time via `POST /analyze`
- 🔒 Skor risiko fraud berdasarkan nominal, waktu, lokasi, dan device
- 📊 Output risk level: Low / Medium / High
- 🧪 Dataset dummy untuk simulasi & testing
- 🧠 Siap dikembangkan ke versi AI/ML

---

## 🏗️ Teknologi yang Digunakan

- Golang (Fiber v2)
- JSON REST API
- Modular structure (`cmd/`, `internal/`, `pkg/`)
- Simulasi data via script custom

---

## 📦 Struktur Proyek

fraud-detector-api/
├── cmd/ # Entry point dan CLI
│ ├── server/ # Main API server
│ └── tester/ # Script untuk testing batch
├── internal/
│ ├── handler/ # HTTP endpoint handler
│ ├── model/ # Struct data transaksi
│ └── service/ # Logika fraud scoring
├── pkg/ # Utilitas dan data generator
├── transactions.json # Data simulasi (tidak masuk git)
├── go.mod / go.sum
└── README.md

