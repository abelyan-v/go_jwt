package main

import (
	"fmt"
	"log"

	"time"
	"net/http"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"github.com/golang-jwt/jwt/v5"
	
	"ProjectBackend/models"

	"ProjectBackend/HttpFunctions"
)

func main() {
	// Работа с базами данных
		// Подключение к базе данных
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	models.MigrateUser(db)
	models.MigrateToken(db)

	// HTTP сервисы
	HttpFunctions.Routes()
	HttpFunctions.DB = db


	http.HandleFunc("/CreateAccessToken", func(w http.ResponseWriter, r *http.Request) {
		secretKey := []byte("your-256-bit-secret")

		claims := jwt.MapClaims{
			"user_id": 123,
			"exp":     time.Now().Add(24 * time.Hour).Unix(),
		}
	
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		tokenString, err := token.SignedString(secretKey)
	
		if err != nil {
			panic(err)
		}
	
		fmt.Println("Токен:", tokenString)

		w.Write([]byte("О нас: это простой сервер на Go"))
	})

	http.HandleFunc("/CheckToken", func(w http.ResponseWriter, r *http.Request) {
		secretKey := []byte("your-256-bit-secret")

		tokenString := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDY2NDI3MzEsInVzZXJfaWQiOjEyM30.mAxX-eAKd-5St2mZd3WDLfI1kyYUaj3hJj0Ju23ciyTxOdTqDYTULwCPzqDllyNBpaDhKb58l5Uur6RdLpDFdQ"
	
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
	
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
	
		if token.Valid {
			fmt.Println("✅ Токен действителен")
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				fmt.Println("Данные:", claims)
			}
		} else {
			fmt.Println("❌ Токен недействителен")
		}

		w.Write([]byte("О нас: это простой сервер на Go"))
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			GUID := r.URL.Query().Get("GUID")
			if GUID == "" {
				fmt.Fprint(w, "Вы должны ввести GUID пользователя")
			}
			fmt.Fprintf(w, "Твой GUID: %s", GUID)
		} else {
			http.Error(w, "Только GET разрешен", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/ReceiveTokenByGUID", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			GUID := r.URL.Query().Get("GUID")
			if GUID == "" {
				http.Error(w, "Вы должны ввести GUID пользователя", http.StatusMethodNotAllowed)
			} else {
				models.CreatePairTokens(db)
				fmt.Fprintf(w, "Твой GUID: %s", GUID)
			}
		} else {
			http.Error(w, "Необходимо заходить только при методе GET", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}