package models

type User struct {
    ID       int    `db:"id"`
    Username string `db:"username"`
    Password string `db:"password"` // Пароль должен храниться зашифрованным
}

// Функция для проверки пользователя (на примере простого сравнения)
func AuthenticateUser(username, password string) (User, error) {
    // Здесь должно быть извлечение пользователя из БД и проверка пароля
    return User{}, nil
}
