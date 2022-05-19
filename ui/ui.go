package ui

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UIState struct {
	app    tview.Application
	layout tview.Flex
}

func NewUIState() *UIState {
	return &UIState{
		app:    *tview.NewApplication(),
		layout: *tview.NewFlex(),
	}
}

func (u UIState) Run() error {
	folder := tview.NewBox().SetBorder(true).SetTitle("Folder")
	textView := tview.NewTextView()
	notes := tview.NewBox().SetBorder(true).SetTitle("Notes")

	flex := tview.NewFlex().
		AddItem(folder, 0, 1, false).
		AddItem(notes, 0, 1, false).
		AddItem(textView, 0, 3, false)
	flex.SetBackgroundColor(tcell.Color101)
	u.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			os.Exit(0)
		}

		// if event.Key() == tcell.KeyCtrlC || event.Key() == tcell.KeyCtrlZ {
		// 	return nil
		// }

		if event.Key() == tcell.KeyCtrlA {
			return nil
		}

		return nil
	})
	if err := u.app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		return err
	}
	return nil
}
