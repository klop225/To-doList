package cmd

import (
	"fmt"
	"os"
)

func Add(file *os.File, text string) (int, string) {
	data, _ := os.ReadFile("To-doList.txt")

	id := 0
	fmt.Sscanf(string(data), "%d", id)

	os.WriteFile("To-doList.txt", []byte(fmt.Sprintf("%d. %s", id+1, text)), 0644)
	return id + 1, text
}

func Delete(id int) {

}
