// В файле internal/api/handlers.go добавьте:
func CreateUser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var user models.User
        
        err := json.NewDecoder(r.Body).Decode(&user)
        if err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }
        
        // Хешируем пароль перед сохранением
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
        if err != nil {
            http.Error(w, "Failed to hash password", http.StatusInternalServerError)
            return
        }
        
        // Сохраняем пользователя в базе данных
        var id int
        err = db.QueryRow(
            "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
            user.Name, user.Email, string(hashedPassword),
        ).Scan(&id)
        
        if err != nil {
            // Проверяем на дублирование email
            if strings.Contains(err.Error(), "duplicate key") {
                http.Error(w, "Email already exists", http.StatusConflict)
                return
            }
            http.Error(w, "Failed to create user", http.StatusInternalServerError)
            return
        }
        
        // Не возвращаем пароль в ответе
        user.ID = id
        user.Password = ""
        
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "message": "User created successfully",
            "user": user,
        })
    }
}