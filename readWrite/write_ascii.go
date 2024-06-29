package output

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func WriteAscii(content, flname string) error {
	if !strings.HasSuffix(flname, ".txt") {
		return fmt.Errorf("usage: go run . --output=<fileName.txt> something standard")
	}
	file, err := os.Create(flname)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.WriteString(file, content)
	if err != nil {
		return err
	}
	return nil
}
