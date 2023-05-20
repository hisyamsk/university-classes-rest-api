package helper

import "github.com/joho/godotenv"

func LoadDotenv() {
	err := godotenv.Load()
	PanicIfError(err)
}
