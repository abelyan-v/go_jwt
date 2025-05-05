package main

import (
	"fmt"
	"jwt/models"
	"log"
)

func main() {
	db, err := models.InitDB("localhost", "postgres", "yourpassword", "medos")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	user := models.User{Username: "test2"}
	if err := user.Create(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User created! ID: %d\n", user.ID)
}