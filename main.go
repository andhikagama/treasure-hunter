package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	text := getTextFromFile("case.txt")

	matrix := make([][]string, len(text))

	fmt.Println("case layout : ")
	for i, line := range text {
		chars := []rune(line)
		strs := make([]string, len(chars))

		for j, v := range chars {
			str := string(v)
			strs[j] = str
		}

		matrix[i] = strs
		fmt.Println(line)
	}
}

func getTextFromFile(filename string) []string {
	var text []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		text = append(text, scanner.Text())
	}

	file.Close()
	return text
}
