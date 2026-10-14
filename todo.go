package todo

import (
	"errors"
	"slices"

	"github.com/google/uuid"
)

type Task struct {
	Name string
	ID   uuid.UUID
}

type TodoList []Task

func (tl *TodoList) Add(text string) uuid.UUID {

	taskID := uuid.New()

	task := Task{text, taskID}
	*tl = append(*tl, task)

	return taskID
}

var errTaskNotFind = errors.New("task not find")

func (tl *TodoList) Get(id uuid.UUID) (Task, error) {

	for _, el := range *tl {

		if el.ID == id {
			return el, nil
		}

	}

	return Task{}, errTaskNotFind

}

func (tl *TodoList) Delete(id uuid.UUID) error {

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

func (tl *TodoList) Update(id uuid.UUID, text string) error {

	for i, el := range *tl {
		if el.ID == id {
			(*tl)[i] = Task{text, el.ID}
			return nil
		}
	}

	return errTaskNotFind
}
