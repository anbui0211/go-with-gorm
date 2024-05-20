package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "user:pass@tcp(127.0.0.1:14306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	log.Println("Connecting to database successfully ...")

	query := `INSERT INTO users (user_id, user_name, age, gender, user_type, status, location_id)
              VALUES (?, ?, ?, ?, ?, ?, 'LC03')`

	for i := 1; i <= 1000; i++ {
		userID := fmt.Sprintf("user_%04d", i)
		userName := fmt.Sprintf("User %d", i)
		age := 20 + (i % 30)    // Tuổi ngẫu nhiên từ 20 đến 49
		gender := (i % 2) + 1   // Giới tính ngẫu nhiên (1 hoặc 2)
		userType := (i % 3) + 1 // Loại người dùng ngẫu nhiên (1, 2 hoặc 3)
		status := "inactive"
		if i%2 == 0 {
			status = "active"
		}

		// Thực hiện chèn dữ liệu
		_ = db.Exec(query, userID, userName, age, gender, userType, status)
	}

	fmt.Println("Successfully inserted 1000 rows into the users table.")
}
