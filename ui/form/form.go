package form

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/rivo/tview"
)

type Task struct {
	ID         string
	Title      string
	Detail     string
	Importance int
	Done       bool
	CreatedAt  time.Time
}

func Run(task *Task) *Task {
	is_save := false
	app := tview.NewApplication()
	form := tview.NewForm().
		AddInputField("Title", task.Title, 20, nil, func(text string) {
			task.Title = text
		}).
		AddDropDown("Importance", []string{"1", "2", "3", "4", "5"}, task.Importance-1, func(option string, optionIndex int) {
			task.Importance, _ = strconv.Atoi(option)
		}).
		AddTextArea("detail", task.Detail, 40, 0, 0, func(text string) {
			task.Detail = text
		}).
		AddCheckbox("Done", task.Done, func(checked bool) {
			task.Done = checked
		}).
		AddButton("Save", func() {
			is_save = true
			app.Stop()
		}).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true)
	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
	if task.ID == "" {
		task.ID = uuid.New().String()
	}
	if is_save {
		return task
	} else {
		return nil
	}
}
