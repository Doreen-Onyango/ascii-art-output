package output

import (
	"errors"
	"os"
	"regexp"
)

var ErrInvalidFileName = errors.New("usage: go run . --output=<fileName.txt> something standard")

func WriteAscii(content, fileName string) error {
    if !isValidOutputFileName(fileName) {
        return ErrInvalidFileName
    }

    return os.WriteFile(fileName, []byte(content), 0644)
}

func isValidOutputFileName(fileName string) bool {
    match, _ := regexp.MatchString(`^.+\.txt$`, fileName)
    return match
}
