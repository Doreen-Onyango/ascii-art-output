package output

import (
	"io"
	"os"
	"strings"

	usage "output/utils"
)

func WriteAscii(content, flname string) error {
	if !strings.HasSuffix(flname, ".txt") {
		usage.PrintUsage()
		return nil
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
