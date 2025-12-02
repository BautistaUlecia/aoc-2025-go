package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFileAsLines(path string) ([]string, error) {
	file, err := os.Open(path)
	var lines []string

	if err != nil {
		log.Fatalf("failed to open file %s", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
