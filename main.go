package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type dungeon struct {
	text   []string
	matrix [][]string
	row    int
	col    int
	startX int
	startY int
}

func main() {
	d := newDungeon("case.txt")
	d.buildMatrix()
	d.findPossibleTreasureLocation()
}

func newDungeon(filename string) dungeon {
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

	return dungeon{
		text: text,
		row:  len(text),
		col:  len(text[0]),
	}
}

func (d *dungeon) buildMatrix() {
	fmt.Println()
	fmt.Println("dungeon layout : ")

	matrix := make([][]string, len(d.text))
	for i, line := range d.text {
		chars := []rune(line)
		strs := make([]string, len(chars))

		for j, v := range chars {
			str := string(v)
			strs[j] = str

			if strings.ToLower(str) == "x" {
				d.startX = i
				d.startY = j
			}
		}

		matrix[i] = strs
		fmt.Println(line)
	}

	d.matrix = matrix
}

func (d *dungeon) findPossibleTreasureLocation() {
	fmt.Println()
	fmt.Println("list possible treasure index location :")

	// rules: up > right > down
	for i := d.startX - 1; i > 0; i-- {
		for j := d.startY + 1; j < d.col; j++ {
			if d.matrix[i][j] == "#" {
				continue
			}

			for k := i + 1; k < d.row; k++ {
				if d.matrix[k][j] == "#" {
					break
				}

				fmt.Println(fmt.Sprintf("[%d,%d]", k, j))
				d.matrix[k][j] = "$"
			}

		}
	}

	fmt.Println()
	fmt.Println("marked possible treasure location in dungeon :")

	for _, v := range d.matrix {
		fmt.Println(strings.Join(v, ""))
	}

	fmt.Println()
}
