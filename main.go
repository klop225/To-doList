package main

import (
	"To-doList/cmd"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("To-doList.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("не получилось открыть файл", err)
		return
	}

	defer file.Close()

	for {
		var text string
		fmt.Scan(&text)

		if text == "add" {
			var task string
			fmt.Print("Введите задачу: ")
			fmt.Scan(&task)
			cmd.Add(file, task)
			fmt.Println("Задача добавлена!")
		} else if text == "exit" {
			fmt.Println("всем пока!")
			break
		}
	}
}
