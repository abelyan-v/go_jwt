package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"github.com/google/uuid"
)


type Users struct {
	gorm.Model
	GUID		string `gorm:"size:36;uniqueIndex"`
	Token		Tokens `gorm:"foreignKey:UserID"`
}

func MigrateUser(db *gorm.DB) {
	err := db.AutoMigrate(&Users{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func CreateUser(db *gorm.DB) {
	id := uuid.New() 
	str := id.String()

	user := Users{
		GUID:	str,
	}

	result := db.Create(&user)
	if result.Error != nil {
		log.Fatalf("Failed to create user: %v", result.Error)
	}
		fmt.Printf("Hello World!\nNew user created:\nID: %d\nGUID: %s\nRows affected: %d\n",
		user.ID,
		user.GUID,
		result.RowsAffected,
	)
}