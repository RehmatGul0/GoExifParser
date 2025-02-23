package utils

import (
	"fmt"
	"os"
	"time"
)

type genericContainsCriteria[T any] func(in T) bool

func Contains[T any](s []T, fn genericContainsCriteria[T]) bool {
	for _, one := range s {
		if fn(one) {
			return true
		}
	}
	return false
}

func CreateFile(filePath string) (*os.File, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("error creating file %s", err.Error())
	}
	return file, nil
}

func GetFileName(format string) string {
	return fmt.Sprintf("%s%s", time.Now().String(), format)
}
