package db

import (
    "fmt"
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/Isabellemew/design-backend/internal/models"
)

var DB *gorm.DB

// 👇 Инициализация подключения и миграции
func InitDB() {
    dsn := "user=postgres password=1234567 dbname=cursovaya port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Ошибка при подключении к базе данных: ", err)
    }
    fmt.Println("✅ Подключение к базе данных успешно установлено!")

    err = DB.AutoMigrate(&models.Order{}, &models.OrderItem{})
    if err != nil {
        log.Fatal("❌ Ошибка миграции модели Order: ", err)
    }
}