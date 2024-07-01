package output

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

var err = errors.New("usage: go run . --output=<fileName.txt> something standard")

func WriteAscii(content, fileName string) error {
	if !isValidFileName(fileName) {
		return err
	}

	errr := os.WriteFile(fileName, []byte(content), 0o644)
	if errr != nil {
		return fmt.Errorf("error while creating a file: %v", errr)
	}
	return errr
}

func isValidFileName(fileName string) bool {
	match, err := regexp.MatchString(`^.+\.txt$`, fileName)
	if err != nil {
		fmt.Printf("Error, %v", err)
	}
	return match
}
