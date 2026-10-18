package todo

import (
	"testing"
)

func TestToDo(t *testing.T) {

	t.Run("test 'Add'", func(t *testing.T) {

		tdList := TodoList{}

		// TODO: handle Error
		id, _ := tdList.Add("Do homework")

		got, err := tdList.Get(id)

		if err != nil {
			t.Errorf("got an error: %v", err)
		}

		want := Task{"Do homework", 1, id}

		if got != want {
			t.Errorf("want: %v | got: %v", want, got)
		}

	})

	t.Run("delete task by id", func(t *testing.T) {

		tdList := TodoList{}

		// TODO: handle Error
		id, _ := tdList.Add("Do homework")

		err := tdList.DeleteByID(id)

		if err != nil {
			t.Errorf("got an error: %v", err)
		}

		_, got := tdList.Get(id)

		want := errTaskNotFind

		if got != want {
			t.Errorf("want: %v | got: %v", want, got)
		}

	})

	t.Run("delete task by number", func(t *testing.T) {
		tdList := TodoList{}

		tdList.Add("Do homework")
		tdList.Add("Do homework 2")
		id, _ := tdList.Add("Do homework 3")

		err := tdList.DeleteByNumber(2)

		if err != nil {
			t.Errorf("got an error: %v", err)
		}

		thirdTask, _ := tdList.Get(id)

		want := 2

		if thirdTask.Number != 2 {
			t.Errorf("want: %v | got: %v", want, thirdTask.Number)
		}

	})

	t.Run("Update task by id", func(t *testing.T) {
		tdList := TodoList{}

		// TODO: handle Error
		id, _ := tdList.Add("Do homework")

		errUpdate := tdList.UpdateById(id, "Do 13 homework")

		if errUpdate != nil {
			t.Errorf("got an error: %v", errUpdate)
		}

		got, errGot := tdList.Get(id)
		if errGot != nil {
			t.Errorf("got an error: %v", errUpdate)
		}

		want := Task{"Do 13 homework", 1, id}

		if got != want {
			t.Errorf("want: %v | got: %v", want, got)
		}

	})

	t.Run("Update task by number", func(t *testing.T) {
		tdList := TodoList{}
		tdList.Add("Do homework")
		id, _ := tdList.Add("Do homework 2")
		tdList.Add("Do homework 3")

		errUpdate := tdList.UpdateByNumber(2, "Do mega homework")

		if errUpdate != nil {
			t.Errorf("got an error: %v", errUpdate)
		}

		got, errGot := tdList.Get(id)

		if errGot != nil {
			t.Errorf("got an error: %v", errUpdate)
		}

		want := Task{"Do mega homework", 2, id}

		if got != want {
			t.Errorf("want: %v | got: %v", want, got)
		}

	})

}
