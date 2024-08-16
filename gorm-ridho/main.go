package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Ganti dengan informasi koneksi database kamu
	dsn := "root:@tcp(127.0.0.1:3306)/tugas-mysql"

	// Membuka koneksi ke database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal terhubung ke database:", err)
		return
	}

	//tabel users
	type User struct {
		ID       int    `gorm:"primaryKey;autoincrement"`
		Name     string `gorm:"type:varchar(100)"`
		Email    string `gorm:"type:varchar(100)"`
		Password string `gorm:"type:varchar(255)"`
		Gender   string `gorm:"type:enum('L', 'P');default:'L'"`
		Photo    string `gorm:"type:varchar(255)"`
		Address  string `gorm:"type:varchar(255)"`
		Role     int    `gorm:"type:int"`
	}

	users := []User{
		{Name: "Fulan", Email: "Fulan@example.com", Password: "hashedpassword1", Gender: "L", Photo: "user1.jpg", Address: "Jl. Cisitu Indah VI no 6", Role: 1},
		{Name: "fulanah", Email: "fulanah@example.com", Password: "hashedpassword2", Gender: "P", Photo: "user2.jpg", Address: "Jl. Cisitu Indah VI no 6", Role: 2},
		{Name: "Ardi", Email: "Ardi@example.com", Password: "hashedpassword3", Gender: "L", Photo: "user3.jpg", Address: "Jl. Cisitu Indah VI no 6", Role: 2},
		{Name: "samsudin", Email: "samsudin@example.com", Password: "hashedpassword4", Gender: "L", Photo: "user4.jpg", Address: "Jl. Cisitu Indah VI no 6", Role: 2},
		{Name: "eko", Email: "eko@example.com", Password: "hashedpassword5", Gender: "L", Photo: "user5.jpg", Address: "Jl. Cisitu Indah VI no 6", Role: 2},
		{Name: "sugeng", Email: "sugeng@example.com", Password: "hashedpassword6", Gender: "L", Photo: "user6.jpg", Address: "Jl. Cisitu Indah VI no 6", Role: 3},
		{Name: "alif", Email: "alif@example.com", Password: "hashedpassword7", Gender: "L", Photo: "user7.jpg", Address: "Jl. Cisitu Indah VI no 6", Role: 3},
		{Name: "siti", Email: "siti@example.com", Password: "hashedpassword8", Gender: "P", Photo: "user8.jpg", Address: "Jl. Cisitu Indah VI no 6", Role: 3},
		{Name: "juminten", Email: "juminten@example.com", Password: "hashedpassword9", Gender: "L", Photo: "user9.jpg", Address: "Jl. Cisitu Indah VI no 6", Role: 3},
		{Name: "paijo", Email: "paijo@example.com", Password: "hashedpassword10", Gender: "L", Photo: "user10.jpg", Address: "Jl. Cisitu Indah VI no 6", Role: 3},
		{Name: "saifudin", Email: "saifudin@example.com", Password: "hashedpassword10", Gender: "L", Photo: "user10.jpg", Address: "Jl. Cisitu Indah VI no 6", Role: 3},
		{Name: "daffa", Email: "daffa@example.com", Password: "hashedpassword10", Gender: "L", Photo: "user10.jpg", Address: "Jl. Cisitu Indah VI no 6", Role: 3},
		{Name: "akbar", Email: "akbar@example.com", Password: "hashedpassword10", Gender: "L", Photo: "user10.jpg", Address: "Jl. Cisitu Indah VI no 6", Role: 3},
		{Name: "rafli", Email: "rafli@example.com", Password: "hashedpassword10", Gender: "L", Photo: "user10.jpg", Address: "Jl. Cisitu Indah VI no 6", Role: 3},
		{Name: "rini", Email: "rini@example.com", Password: "hashedpassword10", Gender: "P", Photo: "user10.jpg", Address: "Jl. Cisitu Indah VI no 6", Role: 3},
	}

	if err := db.Create(&users).Error; err != nil {
		fmt.Println("Gagal menyisipkan data:", err)
		return
	}

	//tabel role
	type role struct {
		ID   int    `gorm:"primaryKey;autoincrement"`
		Name string `gorm:"type:varchar(100)"`
	}

	roles := []role{
		{Name: "Admin"},
		{Name: "Technicial"},
		{Name: "Client"},
	}

	if err := db.Create(&roles).Error; err != nil {
		fmt.Println("Gagal menyisipkan data:", err)
		return
	}

	//tabel ac
	type ac struct {
		ID    int     `gorm:"primaryKey;autoincrement"`
		Name  string  `gorm:"type:varchar(100)"`
		Brand string  `gorm:"type:varchar(100)"`
		PK    float64 `gorm:"type:decimal(3.1)"`
		Price float64 `gorm:"type:decimal(5.5)"`
	}

	acs := []ac{
		{Name: "LG-1", Brand: "LG", PK: 0.5, Price: 50000},
		{Name: "Sharp-2", Brand: "Sharp", PK: 1.0, Price: 60000},
		{Name: "Panasonic-3", Brand: "Panasonic", PK: 2.0, Price: 70000},
		{Name: "Samsung-4", Brand: "Samsung", PK: 0.5, Price: 80000},
		{Name: "Daikin-5", Brand: "Daikin", PK: 1.0, Price: 90000},
		{Name: "Gree-6", Brand: "Gree", PK: 2.0, Price: 100000},
		{Name: "Polytron-7", Brand: "Polytron", PK: 0.5, Price: 110000},
		{Name: "Electrolux-8", Brand: "Electrolux", PK: 1.0, Price: 120000},
		{Name: "Aqua-9", Brand: "Aqua", PK: 2.0, Price: 130000},
		{Name: "Midea-10", Brand: "Midea", PK: 0.5, Price: 140000},
		{Name: "LG-11", Brand: "LG", PK: 1.0, Price: 200000},
		{Name: "Sharp-12", Brand: "Sharp", PK: 2.0, Price: 210000},
		{Name: "Panasonic-13", Brand: "Panasonic", PK: 0.5, Price: 220000},
		{Name: "Samsung-14", Brand: "Samsung", PK: 1.0, Price: 230000},
		{Name: "Daikin-15", Brand: "Daikin", PK: 2.0, Price: 240000},
		{Name: "Gree-16", Brand: "Gree", PK: 0.5, Price: 250000},
		{Name: "Polytron-17", Brand: "Polytron", PK: 1.0, Price: 260000},
		{Name: "Electrolux-18", Brand: "Electrolux", PK: 2.0, Price: 270000},
		{Name: "Aqua-19", Brand: "Aqua", PK: 0.5, Price: 280000},
		{Name: "Midea-20", Brand: "Midea", PK: 1.0, Price: 290000},
	}

	if err := db.Create(&acs).Error; err != nil {
		fmt.Println("Gagal menyisipkan data:", err)
		return
	}

	//tabel service
	type service struct {
		ID            int    `gorm:"primaryKey;autoincrement"`
		Technician_id int    `gorm:"type:int"`
		Client_id     int    `gorm:"type:int"`
		Ac_id         int    `gorm:"type:int"`
		Date          string `gorm:"type:varchar(100)"`
		Status        string `gorm:"type:varchar(100)"`
	}

	services := []service{
		{Technician_id: 2, Client_id: 6, Ac_id: 1, Date: "1 Jun 2020", Status: "finish"},
		{Technician_id: 3, Client_id: 7, Ac_id: 1, Date: "1 May 2020", Status: "finish"},
		{Technician_id: 4, Client_id: 8, Ac_id: 2, Date: "1 Jun 2020", Status: "finish"},
		{Technician_id: 5, Client_id: 9, Ac_id: 3, Date: "1 March 202", Status: "finish"},
		{Technician_id: 2, Client_id: 6, Ac_id: 5, Date: "1 Dec 2021", Status: "finish"},
		{Technician_id: 3, Client_id: 7, Ac_id: 6, Date: "25 Dec 2021", Status: "finish"},
		{Technician_id: 4, Client_id: 10, Ac_id: 7, Date: "1 Jan 2022", Status: "finish"},
		{Technician_id: 5, Client_id: 11, Ac_id: 8, Date: "2 Feb 2022", Status: "finish"},
		{Technician_id: 2, Client_id: 6, Ac_id: 9, Date: "4 Apr 2022", Status: "finish"},
		{Technician_id: 3, Client_id: 7, Ac_id: 10, Date: "5 May 2023", Status: "on repair"},
		{Technician_id: 4, Client_id: 12, Ac_id: 11, Date: "6 Jun 2023", Status: "on repair"},
		{Technician_id: 5, Client_id: 13, Ac_id: 12, Date: "7 Jul 2023", Status: "on repair"},
		{Technician_id: 2, Client_id: 6, Ac_id: 13, Date: "8 Aug 2023", Status: "paid"},
		{Technician_id: 3, Client_id: 7, Ac_id: 14, Date: "9 Sep 2023", Status: "paid"},
		{Technician_id: 4, Client_id: 14, Ac_id: 15, Date: "10 Oct 2023", Status: "unpaid"},
	}

	if err := db.Create(&services).Error; err != nil {
		fmt.Println("Gagal menyisipkan data:", err)
		return
	}

	// Hitung jumlah perbaikan dengan status "paid"
	var count int64
	if err := db.Model(&service{}).Where("status = ?", "paid").Count(&count).Error; err != nil {
		fmt.Println("Gagal menghitung jumlah perbaikan:", err)
		return
	}

	fmt.Printf("Jumlah perbaikan yang sudah dibayar: %d\n", count)

	// Update status pembayaran menjadi "paid" untuk client_id = 7
	if err := db.Model(&service{}).Where("client_id = ?", 7).Update("status", "paid").Error; err != nil {
		fmt.Println("Gagal memperbarui status:", err)
		return
	}

	fmt.Println("Status pembayaran berhasil diperbarui menjadi 'paid' untuk client_id = 7")

	// Menghapus semua data dari tabel service
	if err := db.Exec("DELETE FROM services").Error; err != nil {
		fmt.Println("Gagal menghapus data:", err)
		return
	}

	fmt.Println("Semua data berhasil dihapus dari tabel service")

}
