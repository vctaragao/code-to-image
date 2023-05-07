package helper

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const LOG_FILE = "var/app.log"

func LogError(message string, err error) {
	writeToLogFile(message, err)
	fmt.Println(message+":", getOriginalError(err))
}

func writeToLogFile(message string, err error) {
	logFile, fileErr := os.Create("var/app.log")
	if fileErr != nil {
		fmt.Println("unable to write to log file")
		return
	}

	fmt.Fprintf(io.Writer(logFile), message+": %+v", err)
}

func getOriginalError(err error) error {
	for {
		unwrapped := errors.Unwrap(err)
		if unwrapped == nil {
			return err
		}
		err = unwrapped
	}
}
