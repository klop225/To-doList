package functions

import (
	"bufio"
	"os"
)

func Id() int {
	file, _ := os.Open("To-dolist.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	id := 1
	for scanner.Scan() {
		id++
	}
	return id
}
