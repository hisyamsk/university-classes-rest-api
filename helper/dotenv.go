package helper

import (
	"github.com/joho/godotenv"
)

func LoadDotenv(path string) {
	err := godotenv.Load(path)
	PanicIfError(err)
}
