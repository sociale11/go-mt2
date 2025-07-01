package main

import (
	cmd "go-mt2/cmd"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found")
	}
	cmd.Execute()
}
