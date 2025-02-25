package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var reset string = "\033[0m"

func PrintGreen(text string) {
	green := "\033[32m"
	fmt.Println(green + text + reset)
}

func PrintRed(text string) {
	red := "\033[31m"
	fmt.Println(red + text + reset)
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
