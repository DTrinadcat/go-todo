package todo

import (
	"testing"
)

func TestToDo(t *testing.T) {

	t.Run("test 'Add'", func(t *testing.T) {

		tdList := TodoList{}

		id := tdList.Add("Do homework")

		got, err := tdList.Get(id)

		if err != nil {
			t.Errorf("got an error: %v", err)
		}

		want := Task{"Do homework", id}

		if got != want {
			t.Errorf("want: %v | got: %v", want, got)
		}

	})

	t.Run("delete task", func(t *testing.T) {

		tdList := TodoList{}

		id := tdList.Add("Do homework")

		err := tdList.Delete(id)

		if err != nil {
			t.Errorf("got an error: %v", err)
		}

		_, got := tdList.Get(id)

		want := errTaskNotFind

		if got != want {
			t.Errorf("want: %v | got: %v", want, got)
		}

	})

	t.Run("Update task", func(t *testing.T) {
		tdList := TodoList{}

		id := tdList.Add("Do homework")

		errUpdate := tdList.Update(id, "Do 13 homework")

		if errUpdate != nil {
			t.Errorf("got an error: %v", errUpdate)
		}

		got, errGot := tdList.Get(id)
		if errGot != nil {
			t.Errorf("got an error: %v", errUpdate)
		}

		want := Task{"Do 13 homework", id}

		if got != want {
			t.Errorf("want: %v | got: %v", want, got)
		}

	})

}
