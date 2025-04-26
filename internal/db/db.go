package db

import (
    "fmt"
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

// 👇 Определение модели Product
type Product struct {
    ID         uint   `json:"id" gorm:"primaryKey"`
    Name       string `json:"name"`
    Price      string `json:"price"`
    ImageURL   string `json:"image_url"`
    Brand      string `json:"brand"`
    CategoryID uint   `json:"category_id"`
}

// 👇 Инициализация подключения и миграции
func InitDB() {
    dsn := "user=postgres password=1234567 dbname=cursovaya port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Ошибка при подключении к базе данных: ", err)
    }
    fmt.Println("✅ Подключение к базе данных успешно установлено!")

    // Автомиграция модели Product
    err = DB.AutoMigrate(&Product{})
    if err != nil {
        log.Fatal("❌ Ошибка миграции модели Product: ", err)
    }
}
