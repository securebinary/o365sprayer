package logging

import (
	"os"
)

func LogEnumeratedAccount(file *os.File, email string) {
	file.WriteString(email + "\n")
}
