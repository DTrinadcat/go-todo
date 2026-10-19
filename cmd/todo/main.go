package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"

	"github.com/DTrinadcat/go-todo"
)

const msgEmptyTask = "Failed to add task: name cannot be empty."

func show(tl todo.TodoList) {

	fmt.Println("-----------------------------")
	for _, task := range tl {
		fmt.Printf("%v | id: %v\n", task.Name, task.ID)
	}
	fmt.Println("-----------------------------")
}

func prompt(scanner *bufio.Scanner, message string) string {

	fmt.Print(message)

	scanner.Scan()

	return scanner.Text()
}

func main() {

	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Help page here!")
		return
	}

	tl := todo.TodoList{}

	tl.Load()

	switch os.Args[1] {

	case "show":
		show(tl)

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Task name can not be empty")
		}

		_, err := tl.Add(os.Args[2])

		if err != nil {
			fmt.Println(msgEmptyTask)
		}

		tl.Save()

	case "del":
		if len(os.Args) < 3 {
			fmt.Println("Task number or id can not be empty")
		}

		userArg := os.Args[2]

		num, errNum := strconv.Atoi(userArg)

		if errNum == nil {

			err := tl.DeleteByNumber(num)

			if err != nil {
				fmt.Println("Failed to delete: id or number does not exist")
			}
			tl.Save()

			break
		}

		id, errID := uuid.Parse(userArg)

		if errID == nil {
			err := tl.DeleteByID(id)

			if err != nil {
				fmt.Println("Failed to delete: id or number does not exist")
			}

			tl.Save()

			break
		}

		fmt.Println("Failed to delete: id or number does not exist")

	case "up":
		if len(os.Args) < 4 {
			fmt.Println("Task number or id or text can not be empty")
		}

		userArg := os.Args[2]
		text := os.Args[3]

		num, errNum := strconv.Atoi(userArg)

		if errNum == nil {

			err := tl.UpdateByNumber(num, text)

			if err != nil {
				fmt.Println("Failed to update: id or number does not exist")
			}

			tl.Save()

			break
		}

		id, errID := uuid.Parse(userArg)

		if errID == nil {

			err := tl.UpdateById(id, text)

			if err != nil {
				fmt.Println("Failed to update: id or number does not exist")
			}

			tl.Save()

			break
		}

		fmt.Println("Failed to delete: id or number does not exist")
	default:
		fmt.Println("Unknown command")
	}

}
