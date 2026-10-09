# 🧮 Go BMI Calculator Web Application

Aplikasi web sederhana untuk mengukur *Body Mass Index* (BMI) berbasis Go (*Golang*). Aplikasi ini menerima input data fisik pengguna, menghitung skor BMI, serta memberikan kategori kesehatan dan saran yang relevan.

## 🚀 Fitur Utama
- **Formulir Input Dinamis**: Pengisian nama, umur, berat badan (kg), dan tinggi badan (cm).
- **Validasi Input**: Menangani error input kosong maupun format data non-angka.
- **Kalkulator BMI**: Perhitungan skor BMI akurat beserta kategorisasinya (*Kurus, Normal, Gemuk, Obesitas*).
- **Clean Code & Architecture**: Menggunakan *idiomatic Go*, *helper functions*, serta pemisahan logika dari penanganan HTTP handler.

## 🛠️ Teknologi yang Digunakan
- **Bahasa**: Go (Golang)
- **Modul Standar**: `net/http`, `strconv`, `fmt`
- **Frontend**: HTML5 & CSS3 (Internal styling)

## 💻 Cara Menjalankan di Lokal

1. **Clone repositori ini**:
   ```bash
   git clone [https://github.com/username/bmi-calculator-go.git](https://github.com/username/bmi-calculator-go.git)
   cd bmi-calculator-go