# Merchant-Bank API (Golang)

## Deskripsi

Ini adalah proyek API sederhana yang dikembangkan menggunakan **Golang** untuk mensimulasikan interaksi antara merchant dan bank. API ini menyediakan fungsionalitas login, pembayaran, dan logout. Semua data pelanggan, merchant, dan riwayat disimpan dalam file JSON sebagai simulasi.

---

## Fitur Utama

1. **Login**: Pelanggan dapat login menggunakan username dan password.
2. **Payment**: Pelanggan yang sudah login dapat melakukan pembayaran.
3. **Logout**: Pelanggan dapat logout dari sistem.
4. **History**: Semua aktivitas seperti login, payment, dan logout dicatat dalam riwayat.

---

## Arsitektur

Struktur direktori proyek ini adalah sebagai berikut:

1. data
2. handlers
3. models
4. repository

---

## Teknologi yang Digunakan

- **Golang**: Bahasa pemrograman utama yang digunakan untuk membangun API.
- **Gorilla Mux**: Router HTTP untuk menangani endpoint API.
- **JSON**: Digunakan untuk menyimpan data pelanggan, merchant, dan riwayat aktivitas.

---

## Instalasi dan Menjalankan Proyek

1. **Clone repository**

   ```bash
   git clone https://github.com/saydova/merchant-api-bank.git
   cd merchant-bank-api
   ```

2. **Inisialisasi Go Modules dan Install Dependency**

   ```bash
   go mod init project
   go get github.com/gorilla/mux
   ```

3. **Jalankan Aplikasi**

   ```bash
   go run main.go
   ```

4. **Aplikasi akan berjalan pada port 8080**.
   Akses endpoint dengan `http://localhost:8080`.

---

## API Endpoints

1. **Login**

   - **URL**: `/login`
   - **Method**: `POST`
   - **Request Body**:
     ```json
     {
       "username": "john_doe",
       "password": "password123"
     }
     ```
   - **Response**:
     - Success: `200 OK` dengan pesan `Login Successful`
     - Failure: `401 Unauthorized` dengan pesan `Invalid Username or Password`

2. **Payment**

   - **URL**: `/payment`
   - **Method**: `POST`
   - **Request Body**:
     ```json
     {
       "customer_id": 1,
       "amount": 200
     }
     ```
   - **Response**:
     - Success: `200 OK` dengan pesan `Payment Successful`
     - Failure:
       - `401 Unauthorized` jika pelanggan belum login
       - `400 Bad Request` jika pelanggan tidak ditemukan

3. **Logout**
   - **URL**: `/logout`
   - **Method**: `POST`
   - **Request Body**:
     ```json
     {
       "customer_id": 1
     }
     ```
   - **Response**:
     - Success: `200 OK` dengan pesan `Logout Successful`

---

## Simulasi Data

Data pelanggan, merchant, dan riwayat aktivitas disimpan dalam file JSON yang terletak di folder `/data`:

- **customers.json**: Data pelanggan dengan atribut `id`, `username`, `password`, `name`, dan `balance`.
- **merchants.json**: Data merchant dengan atribut `id` dan `name`.
- **history.json**: Riwayat aktivitas login, payment, dan logout oleh pelanggan.

---

## Contoh Request

1. **Login**:
   ```bash
   curl -X POST http://localhost:8080/login -d '{"username":"john_doe","password":"password123"}' -H "Content-Type: application/json"
   ```
