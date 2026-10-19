package todo

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"slices"

	"github.com/google/uuid"
)

var errTaskNotFind = errors.New("task not find")
var errEmptyString = errors.New("String must not be empty")

type Task struct {
	Name string
	ID   uuid.UUID
}

type TodoList []Task

func (tl *TodoList) Load() error {

	dir, _ := os.UserHomeDir()

	path := filepath.Join(dir, ".todos.json")

	data, err := os.ReadFile(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err

	}

	return json.Unmarshal(data, tl)
}

func (tl *TodoList) Save() error {

	dir, _ := os.UserHomeDir()

	path := filepath.Join(dir, ".todos.json")

	data, err := json.Marshal(tl)

	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)

}

func (tl *TodoList) Add(text string) (uuid.UUID, error) {

	if text == "" {
		return uuid.Nil, errEmptyString
	}

	taskID := uuid.New()

	task := Task{text, taskID}
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

	if len(*tl) < number {
		return errTaskNotFind
	}

	*tl = slices.Delete(*tl, number-1, number)

	return nil

}

func (tl *TodoList) UpdateById(id uuid.UUID, text string) error {

	for i, el := range *tl {
		if el.ID == id {
			(*tl)[i] = Task{text, el.ID}
			return nil
		}
	}

	return errTaskNotFind
}

func (tl *TodoList) UpdateByNumber(number int, text string) error {
	if len(*tl) < number {
		return errTaskNotFind
	}

	(*tl)[number-1] = Task{text, (*tl)[number-1].ID}

	return nil
}
