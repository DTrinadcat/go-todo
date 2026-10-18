package todo

import (
	"errors"
	"slices"

	"github.com/google/uuid"
)

var errTaskNotFind = errors.New("task not find")
var errEmptyString = errors.New("String must not be empty")

type Task struct {
	Name   string
	Number int
	ID     uuid.UUID
}

type TodoList []Task

func getLastNumber(tl *TodoList) int {

	if len(*tl) == 0 {
		return 0
	}

	return (*tl)[len(*tl)-1].Number

}

func (tl *TodoList) Add(text string) (uuid.UUID, error) {

	if text == "" {
		return uuid.Nil, errEmptyString
	}

	taskID := uuid.New()

	task := Task{text, getLastNumber(tl) + 1, taskID}
	*tl = append(*tl, task)

	return taskID, nil
}

func (tl *TodoList) Get(id uuid.UUID) (Task, error) {

	for _, el := range *tl {

		if el.ID == id {
			return el, nil
		}

	}

	return Task{}, errTaskNotFind

}

func (tl *TodoList) DeleteByID(id uuid.UUID) error {

	isFind := false

	*tl = slices.DeleteFunc(*tl, func(task Task) bool {
		if task.ID == id {
			isFind = true
			return true
		}
		return false
	})

	if isFind {
		return nil
	}

	return errTaskNotFind

}

func (tl *TodoList) DeleteByNumber(number int) error {
	isFind := false

	var newTl []Task

	for i := 0; i < len(*tl); i++ {

		task := (*tl)[i]

		if !isFind {

			if task.Number == number {
				isFind = true
				continue
			}

			newTl = append(newTl, task)

		} else {

			updatedTask := Task{task.Name, task.Number - 1, task.ID}

			newTl = append(newTl, updatedTask)

		}

	}

	*tl = newTl

	if isFind {
		return nil
	}

	return errTaskNotFind

}

func (tl *TodoList) UpdateById(id uuid.UUID, text string) error {

	for i, el := range *tl {
		if el.ID == id {
			(*tl)[i] = Task{text, el.Number, el.ID}
			return nil
		}
	}

	return errTaskNotFind
}

func (tl *TodoList) UpdateByNumber(number int, text string) error {
	if getLastNumber(tl) < number {
		return errTaskNotFind
	}

	(*tl)[number-1] = Task{text, number, (*tl)[number-1].ID}

	return nil
}
