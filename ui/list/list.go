package list

import (
	"termtsk/ui/form"

	"github.com/rivo/tview"
)

func Run(taskList []*form.Task) int {
	app := tview.NewApplication()
	list := tview.NewList()
	selectNum := -1
	for i := 0; i < len(taskList); i++ {
		func(j int) {
			var r rune = ' '
			if taskList[j].Done {
				r = '◉'
			}
			list.AddItem(taskList[j].Title, taskList[j].Detail, r, func() {
				selectNum = j
				app.Stop()
			})
		}(i)
	}
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
	return selectNum
}
