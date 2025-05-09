package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"crypto/rand"
	"math/big"
	"golang.org/x/crypto/bcrypt"
)

type Tokens struct {
	gorm.Model
	UserID	uint   `gorm:"uniqueIndex"`
	Token	string `gorm:"size:60"`
}

func MigrateToken(db *gorm.DB) {
	err := db.AutoMigrate(&Tokens{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func CreateRefreshToken(db *gorm.DB, RefreshTokenHashed string) {
	token := Tokens{
		UserID: 1,
		Token: RefreshTokenHashed,
	}

	// Сохраняем в базу
	result := db.Create(&token)
	if result.Error != nil {
		log.Fatalf("Failed to create token: %v", result.Error)
	}
		fmt.Printf("Hello World!\nNew token created:\nID: %d\nUser ID: %s\nToken: %s\nRows affected: %d\n",
		token.ID,
		token.UserID,
		token.Token,
		result.RowsAffected,
	)
}

func RandomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	
	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		result[i] = chars[num.Int64()]
	}
	
	return string(result)
}

func StringToBcrypt(s string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

//Настоящие действия
func CreatePairTokens(db *gorm.DB) {
	RefreshTokenString := RandomString(72)
	RefreshTokenHashed, err := StringToBcrypt(RefreshTokenString)
	if err != nil {
		log.Fatalf("Ошибка хеширования: %v", err)
	}
	CreateRefreshToken(db, RefreshTokenHashed)
}