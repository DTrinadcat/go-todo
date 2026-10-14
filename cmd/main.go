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
			userStrID := prompt(scanner, "Enter task id: ")

			id, _ := uuid.Parse(userStrID) // TODO The error should be handled

			// TODO It should be checked whether the specified ID exist in tl

			userStrName := prompt(scanner, "Enter task name: ")

			err := tl.Update(id, userStrName)

			if err != nil {
				fmt.Println("Failed to update task: id does not exist")
			}

		case "4": // Delete
			// TODO: DRY principle

			userStrID := prompt(scanner, "Enter task id: ")

			id, _ := uuid.Parse(userStrID) // TODO The error should be handled

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
