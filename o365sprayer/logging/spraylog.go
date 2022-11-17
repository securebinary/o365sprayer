package logging

import (
	"os"
)

func LogSprayedAccount(file *os.File, email string, password string) {
	file.WriteString(email + ":" + password + "\n")
}
