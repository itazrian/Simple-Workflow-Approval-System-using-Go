
Sistem Workflow Approval Sederhana menggunakan Go

Project ini adalah Sistem Workflow Approval Sederhana yang dibuat menggunakan Go (Golang).
Fungsinya untuk mengelola proses persetujuan (approval) yang terdiri dari beberapa tahapan (step) dan harus disetujui secara berurutan.

# 1. Cara Menjalankan Aplikasi

- Kebutuhan Sistem
* Go 1.24+ (menggunakan framework Echo)
* MySQL 8.0+
* Git

- Clone Repository
```bash
git clone https://github.com/itazrian/Simple-Workflow-Approval-System-using-Go.git
cd Simple-Workflow-Approval-System-using-Go
```

- Install Dependency
```bash
go mod tidy
```

- Menjalankan Aplikasi
```bash
go run main.go
```

Secara default aplikasi akan berjalan di:
```
http://localhost:8080
```

# 2. Konfigurasi Database

- Membuat Database
```sql
CREATE DATABASE workflow_db;
```

- Migrasi Database
    Aplikasi ini menggunakan GORM AutoMigrate, sehingga tabel akan dibuat otomatis saat aplikasi pertama kali dijalankan.
    Tabel yang akan dibuat:
    * `workflows`
    * `workflow_steps`
    * `request_approvals`
    * `request`

---

# 3. API Endpoint

-Workflow

- Membuat Workflow Baru

```
POST /api/workflows
```

Contoh request body:

```json
{
  "name": "Purchase Approval",
}
```

# 4. Keputusan Desain Sistem

4.1 Arsitektur
    Sistem menggunakan RESTful API agar:
        * Mudah dipahami
        * Mudah dikembangkan
        * Mudah diintegrasikan

    Struktur aplikasi dibagi menjadi beberapa layer:
        * Handler → menangani request HTTP
        * Service → berisi logic bisnis
        * Repository → akses database

4.2 Desain Workflow
    * Workflow terdiri dari beberapa step berurutan
    * Approval harus dilakukan satu per satu
    * Step berikutnya tidak bisa diproses sebelum step sebelumnya selesai

4.3 Desain Database
    * `workflow_id` menggunakan tipe `VARCHAR`  (bukan `TEXT`)
    * Setiap data menggunakan UUID agar ID unik dan aman
    * Index diterapkan pada foreign key untuk meningkatkan performa

# 5. Asumsi dan Kompromi

- Asumsi
    * Satu step hanya membutuhkan satu approval
    * Alur approval linear (tidak bercabang)
    * Autentikasi dan otorisasi ditangani oleh sistem lain

- Kompromi (Trade-off)
    * Menggunakan AutoMigrate untuk kemudahan, bukan migrasi manual
    * Belum ada Role-Based Access Control (RBAC)
    * Belum ada mekanisme retry atau rollback jika approval gagal

Semua keputusan ini dibuat agar sistem tetap sederhana dan mudah dikembangkan.

# 6. Rencana Pengembangan Selanjutnya

Beberapa fitur yang bisa dikembangkan:
    * Approval berbasis role (RBAC)
    * Approval paralel (lebih dari satu approver)
    * Versi workflow
    * Audit log & riwayat approval
    * Autentikasi (JWT / OAuth)
