# Aplikasi Manajemen Peminjaman Komik 📚

Aplikasi konsol sederhana yang dikembangkan dengan **Go (Golang)** untuk mengelola data buku komik, data anggota (member), dan transaksi peminjaman/pengembalian di sebuah toko komik.

## 🌟 Fitur Utama

* **Manajemen Data Komik:** Tambah, Edit, Hapus, dan Lihat data komik.
* **Manajemen Data Member:** Tambah, Edit, dan Hapus data member.
* **Transaksi Peminjaman & Pengembalian:** Mencatat peminjaman dan mengupdate stok/data member.
* **Pengembalian:** Menghitung biaya denda (rusak, kotor, atau durasi pinjam melebihi 7 hari).
* **Pencarian & Pengurutan:**
    * Mengurutkan buku berdasarkan tahun rilis (menurun) dan stok (menurun).
    * Mencari buku berdasarkan ID menggunakan **Binary Search** (setelah diurutkan dengan **Insertion Sort**).
    * Mencetak daftar buku berdasarkan genre.

## 🛠️ Instalasi dan Cara Menjalankan

### Prasyarat

Pastikan Anda telah menginstal **Go (Golang)** di sistem Anda.

### Menjalankan Aplikasi

1.  **Kloning Repository**
    ```bash
    git clone [URL_REPOSITORY_ANDA]
    cd [NAMA_FOLDER_REPOSITORY]
    ```
2.  **Menjalankan Kode**
    Aplikasi ini berjalan sebagai program Go tunggal.
    ```bash
    go run bukukomik.go
    ```
3.  **Menggunakan Aplikasi**
    Ikuti instruksi menu yang muncul di konsol (Command Prompt/Terminal).

## 💻 Struktur Data dan Algoritma

| Struktur Data | Implementasi |
| :--- | :--- |
| **Buku Komik** | `struct buku_komik`, Array of struct (`tabbuku`) |
| **Member Toko** | `struct member_toko`, Array of struct (`tabmember`) |
| **Pencarian ID** | **Binary Search** (`binersearch`) |
| **Pengurutan ID** | **Insertion Sort** (`Shortid`) |
| **Pengurutan Stok** | **Selection Sort** (`shortstok`) |
| **Pengurutan Tahun** | **Insertion Sort** (`shorttahun`) |

## 🤝 Kontributor

* Muhammad Rizqy Ramdani
