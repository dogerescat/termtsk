package form

import "github.com/rivo/tview"

func PrintForm() {
	app := tview.NewApplication()
	form := tview.NewForm().
		AddInputField("Title", "", 20, nil, nil).
		AddDropDown("Importance", []string{"1", "2", "3", "4", "5"}, 0, nil).
		AddTextArea("detail", "", 40, 0, 0, nil).
		AddCheckbox("Done", false, nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true)
	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
