package main


import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/yourusername/project/db"
    "github.com/yourusername/project/models"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
    // Логика авторизации пользователя
}

func main() {
    router := mux.NewRouter()

    db, err := db.Connect()
    if err != nil {
        log.Fatal(err)
    }

    router.HandleFunc("/login", loginHandler).Methods("POST")

    log.Fatal(http.ListenAndServe(":8080", router))
}
