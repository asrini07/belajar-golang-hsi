# Mini Project CLI Sederhana: Validasi Umur

## Cara Menjalankan Program
1. masuk ke folder soal-1-cli-validasi
    ```
    cd soal-1-cli-validasi
    ```
2. jalankan program 
    ```
    go run main.go
    ```

## Cara Kerja Program
1. masukan nama, dan umur diatas 18 maka keluar "Selamat Datang, [Nama]"!
2. masukan nama, dan umur dibawah 18 maka keluar error "Error: umur tidak valid (minimal 18 tahun)"
3. masukan nama, dan umur dengan string maka keluar error "Error: input umur harus berupa angka"
4. tidak memasukan nama maka keluar error "Error: nama harus diisi "