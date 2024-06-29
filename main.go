package main

import (
	"flag"
	"fmt"
	"log"

	check "output/checksum"
	print "output/printAscii"
	output "output/readWrite"
)

var banners = map[string]string{
	"standard":   "standard.txt",
	"thinkertoy": "thinkertoy.txt",
	"shadow":     "shadow.txt",
}

func main() {
	checksum := flag.Bool("checksum", false, "Check integrity of specified file")
	flname := flag.String("output", "banner.txt", "Usage: go run . [OPTION] [STRING] [BANNER]")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Please provide at least one argument.")
		return
	}

	input := args[0]
	var banner string

	if len(args) > 1 {
		banner = args[1]
	} else {
		banner = "standard"
	}

	filename, ok := banners[banner]
	if !ok {
		fmt.Println("Invalid banner specified.")
		return
	}
	if *flname != "" {
	}

	if *checksum {
		err := check.ValidateFileChecksum(filename)
		if err != nil {
			log.Fatalf("Error checking integrity: %v", err)
		}
		fmt.Printf("Integrity check passed for file: %s\n", filename)
		return
	}

	err := check.ValidateFileChecksum(filename)
	if err != nil {
		log.Printf("Error downloading or validating file: %v", err)
		return
	}

	asciiArtGrid, err := output.ReadAscii(filename)
	if err != nil {
		log.Fatalf("Error reading ASCII map: %v", err)
	}

	err = print.PrintArt(input, asciiArtGrid)
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
