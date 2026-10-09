package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func formulirHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Tampilan Form HTML saat akses pertama kali (GET)
	if r.Method != "POST" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, formHTML)
		return
	}

	// 2. Ambil Input dari Form (POST)
	nama := r.FormValue("nama")
	umurStr := r.FormValue("umur")
	beratStr := r.FormValue("berat")
	tinggiStr := r.FormValue("tinggi")

	// 3. Validasi Kolom Kosong
	if nama == "" || umurStr == "" || beratStr == "" || tinggiStr == "" {
		renderError(w, "Semua kolom input wajib diisi!")
		return
	}

	// 4. Konversi Tipe Data & Validasi Angka
	umur, errUmur := strconv.Atoi(umurStr)
	beratFloat, errBerat := strconv.ParseFloat(beratStr, 64)
	tinggiFloat, errTinggi := strconv.ParseFloat(tinggiStr, 64)

	if errUmur != nil || errBerat != nil || errTinggi != nil {
		renderError(w, "Umur, berat, dan tinggi badan harus berupa angka!")
		return
	}

	// 5. Hitung BMI
	tinggiMeter := tinggiFloat / 100
	bmi := beratFloat / (tinggiMeter * tinggiMeter)

	// 6. Penentuan Kategori & Saran (Logika Lebih Ringkas)
	var kategori, saran string
	switch {
	case bmi < 18.5:
		kategori = "Kurus"
		saran = "Anda disarankan untuk menambah berat badan dengan pola makan sehat."
	case bmi < 25.0:
		kategori = "Normal"
		saran = "Pertahankan pola makan sehat dan rutin berolahraga."
	case bmi < 30.0:
		kategori = "Gemuk"
		saran = "Disarankan untuk mengurangi berat badan dengan pola makan sehat dan olahraga."
	default:
		kategori = "Obesitas"
		saran = "Segera konsultasikan dengan dokter untuk mendapatkan saran kesehatan yang tepat."
	}

	// 7. Render Tampilan Hasil
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, hasilHTML, nama, umur, beratFloat, tinggiFloat, bmi, kategori, saran)
}

// Fungsi pembantu untuk menampilkan pesan error agar tidak berulang
func renderError(w http.ResponseWriter, pesan string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h3>Error: %s</h3><br><a href='/'>Kembali ke Form</a>", pesan)
}

func main() {
	http.HandleFunc("/", formulirHandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}

// --- TEMPLATE HTML DILUAR FUNGSI (Agar Fungsi Utama Tetap Bersih) ---

const formHTML = `
<!DOCTYPE html>
<html>
<head>
	<title>Kalkulator BMI</title>
	<style>
		body { font-family: Arial, sans-serif; background-color: #f4f6f9; padding: 30px; max-width: 400px; margin: 0 auto; }
		form { background: white; padding: 20px; border-radius: 8px; box-shadow: 0 2px 5px rgba(0,0,0,0.1); }
		input[type="text"] { width: 100%; padding: 8px; margin-top: 5px; margin-bottom: 15px; box-sizing: border-box; border: 1px solid #ccc; border-radius: 4px; }
		input[type="submit"] { background-color: #4CAF50; color: white; padding: 10px 15px; border: none; border-radius: 4px; cursor: pointer; width: 100%; }
		input[type="submit"]:hover { background-color: #45a049; }
	</style>
</head>
<body>
	<form action="/" method="POST">
		<h2>Aplikasi Kalkulator BMI</h2>
		<label>Nama:</label>
		<input type="text" name="nama">
		
		<label>Umur:</label>
		<input type="text" name="umur">
		
		<label>Berat Badan (kg):</label>
		<input type="text" name="berat">
		
		<label>Tinggi Badan (cm):</label>
		<input type="text" name="tinggi">
		
		<input type="submit" value="Hitung BMI">
	</form>
</body>
</html>
`

const hasilHTML = `
<!DOCTYPE html>
<html>
<head>
	<title>Hasil Perhitungan BMI</title>
	<style>
		body { font-family: Arial, sans-serif; background-color: #f4f6f9; padding: 30px; max-width: 500px; margin: 0 auto; }
		.card { background: white; padding: 25px; border-radius: 8px; box-shadow: 0 2px 5px rgba(0,0,0,0.1); }
		h2 { color: #2c3e50; margin-top: 0; }
		p { font-size: 16px; line-height: 1.5; color: #333; margin: 8px 0; }
		.btn { display: inline-block; padding: 10px 15px; background: #3498db; color: white; text-decoration: none; border-radius: 4px; margin-top: 15px; }
		.btn:hover { background: #2980b9; }
	</style>
</head>
<body>
	<div class="card">
		<h2>Hasil Perhitungan BMI</h2>
		<p><b>Nama:</b> %s</p>
		<p><b>Umur:</b> %d tahun</p>
		<p><b>Berat:</b> %.2f kg</p>
		<p><b>Tinggi:</b> %.2f cm</p>
		<p><b>BMI:</b> %.2f</p>
		<p><b>Kategori:</b> %s</p>
		<p><b>Saran:</b> %s</p>
		<a href="/" class="btn">Kembali ke Form</a>
	</div>
</body>
</html>
`
