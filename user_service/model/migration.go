package model

func migration() {
	DB.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&User{})
	DB.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&NewUser{})
	DB.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&Message{})
}
