# KB FINANSIA HOME TEST

Project ini menggunakan bahasa program golang & database mysql

## Installation

1. Replace dan ganti nama config.yml.example menjadi config.yml
2. Setup konfigurasi database
3. Jalankan migrasi database menggunakan command `go run main.go migration up`
4. Jalankan service menggunakan command `go run main.go start`, akses menggunakan port `8081`
5. Unit test menggunakan command `go test main_test.go -v`

## Docker

1. Jalankan command `docker-compose up -d --build`
2. Jalankan command `docker exec kbfinansia go run main.go migration up`
3. Akses menggunakan port `8081`

## Keamanan Aplikasi (OWASP)

1. Autentikasi dan Manajemen Sesi:
    - Implementasi JWT untuk autentikasi.
    - Implementasi secret key JWT kuat dan unik.
    - Implementasi middleware autentikasi token JWT divalidasi sebelum mengizinkan akses ke rute yang dilindungi.

2. Validasi Input:
    - Implementasi menggunakan `c.Bind(reqData)` untuk membaca dan memproses input pengguna. Pastikan semua input divalidasi dengan benar untuk mencegah serangan seperti XSS dan injeksi.

3. Pengelolaan Kesalahan:
    - Setiap log error harus memiliki ID unik, yang membantu dalam pelacakan dan debugging.

4. Logging:
    - Logging diimplementasikan.

5. Pengamanan Data Sensitif:
    - Implementasi password menerapkan bycrypt.

## Prinsip ACID

1. Atomicity:
    - Implementasi menggunakan transaksi (`tx := h.Db.Gorm.Begin()`) di beberapa endpoint, yang membantu memastikan atomicity. Pastikan semua operasi dalam satu transaksi berhasil atau gagal sebagai satu kesatuan.

2. Consistency:
    - Validasi input membantu menjaga konsistensi data.
    - Semua aturan bisnis diterapkan secara konsisten di seluruh aplikasi.

3. Isolation:
    - GORM mengelola tingkat isolasi transaksi.

4. Durability:
    - Durability dijamin oleh sistem manajemen database yang digunakan MySQL.

## Catatan Tambahan

- Rate Limiting dan CORS: Implementasi menggunakan middleware untuk rate limiting dan CORS, untuk keamanan dan menghindari serangan DDoS.
- Autentikasi dan Otorisasi: Implementasi logika otorisasi, terutama di rute yang sensitif, untuk memastikan bahwa pengguna hanya dapat mengakses data atau fungsi yang sesuai dengan peran mereka.
- Pembaruan dan Patch: Implementasi dependensi, terutama framework dan pustaka keamanan, selalu diperbarui.
