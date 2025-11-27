package cmd

import (
	"To-doList/functions"
	"fmt"
	"os"
)

func Add(file *os.File, text string) {
	id := functions.Id()
	file.WriteString(fmt.Sprintf("%d. %s\n", id, text))
}

func Delete(id int) {

}
