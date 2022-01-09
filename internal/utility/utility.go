package utility

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func LogErrorDetails(errorString string, err error) (newErr error) {
	errstr := fmt.Sprintf("[Error] %s - %s", errorString, err.Error())
	log.Printf(errstr)
	return errors.New(errorString)
}
