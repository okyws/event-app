# Event App

REST API sederhana untuk membuat, melihat, mengubah, dan menghapus data event. Aplikasi dibuat dengan Go, Gin, GORM, dan MySQL.

## Persyaratan

- Go versi yang sesuai dengan `go.mod` (Go 1.27.3)
- MySQL, misalnya MySQL dari Laragon

## Menjalankan aplikasi

1. Buat database MySQL:

   ```sql
   CREATE DATABASE event_app;
   ```

2. Buat file `.env` di direktori utama project dan sesuaikan koneksi database:

   ```env
   DATABASE_URL=root:@tcp(127.0.0.1:3306)/event_app?charset=utf8mb4&parseTime=True&loc=Local
   ```

   Contoh tersebut untuk MySQL lokal dengan user `root` tanpa password. Jika MySQL menggunakan password, isi bagian setelah `root:`:

   ```env
   DATABASE_URL=root:password_anda@tcp(127.0.0.1:3306)/event_app?charset=utf8mb4&parseTime=True&loc=Local
   ```

3. Pastikan MySQL berjalan, lalu dari direktori utama project jalankan:

   ```bash
   go mod download
   go run .
   ```

   Saat aplikasi mulai, GORM akan membuat atau memperbarui tabel secara otomatis. API berjalan di `http://localhost:8080`.

## Endpoint API

Semua endpoint menggunakan prefix `/api/v1`.

| Method | Endpoint | Keterangan |
| --- | --- | --- |
| `POST` | `/api/v1/register` | Membuat user |
| `POST` | `/api/v1/login` | Melakukan login |
| `GET` | `/api/v1/events` | Mendapatkan semua event |
| `POST` | `/api/v1/event` | Membuat event |
| `GET` | `/api/v1/event/:id` | Mendapatkan event berdasarkan ID |
| `PUT` | `/api/v1/event/:id` | Mengubah event berdasarkan ID |
| `DELETE` | `/api/v1/event/:id` | Menghapus event berdasarkan ID |

### Format data event

Field `name`, `description`, `location`, dan `user_id` wajib diisi. Field `datetime` opsional; jika tidak diberikan saat membuat event, waktu saat ini digunakan.

```json
{
  "name": "Konferensi Teknologi",
  "description": "Diskusi tentang teknologi terkini",
  "location": "Jakarta",
  "user_id": 1,
  "datetime": "2026-12-01T09:00:00+07:00"
}
```

## Struktur project

```text
.
├── config/       # Koneksi dan konfigurasi database
├── controller/   # Handler endpoint event
├── model/        # Model Event untuk GORM
├── main.go       # Inisialisasi aplikasi dan routing
└── .env          # Konfigurasi lokal
```
