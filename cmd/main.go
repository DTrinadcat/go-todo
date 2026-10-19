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
	tl := todo.TodoList{}
	tl.Load()
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

			tl.Save()

		case "3": // Update
			userAns := prompt(scanner, "Enter Number or ID: ")
			text := prompt(scanner, "Enter text: ")

			num, errNum := strconv.Atoi(userAns)

			if errNum == nil {

				err := tl.UpdateByNumber(num, text)

				if err != nil {
					fmt.Println("Failed to update: id or number does not exist")
				}

				tl.Save()

				break
			}

			id, errID := uuid.Parse(userAns)

			if errID == nil {

				err := tl.UpdateById(id, text)

				if err != nil {
					fmt.Println("Failed to update: id or number does not exist")
				}

				tl.Save()

				break
			}

			fmt.Println("Failed to delete: id or number does not exist")

		case "4": // Delete

			userAns := prompt(scanner, "Enter Number or ID: ")

			num, errNum := strconv.Atoi(userAns)

			if errNum == nil {

				err := tl.DeleteByNumber(num)

				if err != nil {
					fmt.Println("Failed to delete: id or number does not exist")
				}
				tl.Save()

				break
			}

			id, errID := uuid.Parse(userAns)

			if errID == nil {
				err := tl.DeleteByID(id)

				if err != nil {
					fmt.Println("Failed to delete: id or number does not exist")
				}

				tl.Save()

				break
			}

			fmt.Println("Failed to delete: id or number does not exist")

		case "exit":
			fmt.Println("Bye, Bye!")
			return
		default:
			fmt.Println("Unknown command. Please try again")
		}

	}

}
