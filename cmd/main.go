package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/google/uuid"

	"github.com/DTrinadcat/go-todo"
)

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
			fmt.Println(tl)
		case "2": // Add
			fmt.Print("Enter task name: ")

			scanner.Scan()

			userStr := scanner.Text()

			// TODO Should be added an error throw or return to the loop if the user entered an empty string

			if userStr != "" {
				tl.Add(userStr)
			}
		case "3": // Update
			fmt.Print("Enter task id: ")

			scanner.Scan()

			userStrID := scanner.Text()

			id, _ := uuid.Parse(userStrID) // TODO The error should be handled

			// TODO It should be checked whether the specified ID exist in tl

			fmt.Print("Enter task name: ") // TODO DRY principle

			scanner.Scan()

			userStrName := scanner.Text()

			// TODO Should be added an error throw or return to the loop if the user entered an empty string

			tl.Update(id, userStrName)

		case "4": // Delete
			// TODO: DRY principle

			fmt.Print("Enter task id: ")

			scanner.Scan()

			userStrID := scanner.Text()

			id, _ := uuid.Parse(userStrID) // TODO The error should be handled

			tl.Delete(id)
		}

	}

}
