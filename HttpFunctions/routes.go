package HttpFunctions

import (
	"net/http"
	"ProjectBackend/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Routes() {
	// Главная страница
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Главная страница"))
	})

	// Создание пользователя
	http.HandleFunc("/CreateUser", func(w http.ResponseWriter, r *http.Request) {
		models.CreateUser(DB)
		w.Write([]byte("Новый пользователь был создан"))
	})
}