package models

type User struct {
    ID           int    `json:"id"`
    Name         string `json:"name"`
    Email        string `json:"email"`
    Password     string `json:"password"`      // если ты принимаешь пароль от клиента
    PasswordHash string `json:"-"`             // если ты хранишь хеш в БД
}