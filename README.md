# Test Backend SV

## Clone Repository

Jalankan perintah berikut untuk clone repository

```bash
git clone https://github.com/rochimfn/sv-be.git
cd sv-be
```

## Menjalankan Service

Buka file `app/config.go` ubah untuk menyesuaikan konfigurasi koneksi mysql. Pastikan database pada mysql memiliki nama sama dengan nilai `DB_SCHEMA`.

Jalankan service dengan perintah berikut

```bash
go run main.go
```

Perintah diatas sekaligus akan menjalankan migrasi tabel secara otomatis.

Kunjungi [http://localhost:8080/ping](http://localhost:8080/ping) untuk mengecek kondisi service.

### Manual Migrasi

Jalankan query DDL pada berkas `sql/init.sql` untuk membuat database dan tabel awal.

