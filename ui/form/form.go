package form

import (
	"strconv"

	"github.com/rivo/tview"
)

type Task struct {
	title      string
	detail     string
	importance int
	done       bool
}

func Run() *Task {
	var task Task
	is_save := false
	app := tview.NewApplication()
	form := tview.NewForm().
		AddInputField("Title", "", 20, nil, func(text string) {
			task.title = text
		}).
		AddDropDown("Importance", []string{"1", "2", "3", "4", "5"}, 0, func(option string, optionIndex int) {
			task.importance, _ = strconv.Atoi(option)
		}).
		AddTextArea("detail", "", 40, 0, 0, func(text string) {
			task.detail = text
		}).
		AddCheckbox("Done", false, func(checked bool) {
			task.done = checked
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
	if is_save {
		return &task
	} else {
		return nil
	}
}
