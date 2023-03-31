package Logger

import (
	"log"
	"os"
)

func Logger() *log.Logger {
	return log.New(os.Stdout, "ERROR\t", log.Ldate|log.Lshortfile|log.Ltime)
}
