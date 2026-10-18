package main

import (
	"bufio"
	"fmt"
	"os"

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

func promptID(scanner *bufio.Scanner) (uuid.UUID, error) {
	userInputID := prompt(scanner, "Enter task ID: ")
	return uuid.Parse(userInputID)
}

func main() {
	tl := todo.TodoList{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Hello, this is your TODO List")
	fmt.Println("Enter...")
	for {
		fmt.Println("1: Show tasks | 2: Add task | 3: Update task | 4: Delete task | exit: Exit program")
		fmt.Print("Action: ")

		if !scanner.Scan() {
			break
		}

		choice := scanner.Text()

		switch choice {
		case "1": // Show
			show(tl)
		case "2": // Add

			userStr := prompt(scanner, "Enter task name: ")

			_, err := tl.Add(userStr)

			if err != nil {
				fmt.Println(msgEmptyTask)
			}

		case "3": // Update
			id, errIDParsing := promptID(scanner)

			if errIDParsing != nil {
				fmt.Println("Invalid ID. Please try again")
				break
			}

			userStrName := prompt(scanner, "Enter task name: ")

			err := tl.Update(id, userStrName)

			if err != nil {
				fmt.Println("Failed to update task: id does not exist")
			}

		case "4": // Delete
			id, errIDParsing := promptID(scanner)

			if errIDParsing != nil {
				fmt.Println("Invalid ID. Please try again")
				break
			}

			err := tl.Delete(id)

			if err != nil {
				fmt.Println("Failed to delete task: id does not exist")
			}
		case "exit":
			fmt.Println("Bye, Bye!")
			return
		default:
			fmt.Println("Unknown command. Please try again")
		}

	}

}
