package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UIState struct {
	app            tview.Application
	layout         tview.Flex
	flexFocusIndex int
}

func NewUIState() *UIState {
	return &UIState{
		app:            *tview.NewApplication(),
		layout:         *tview.NewFlex(),
		flexFocusIndex: 0,
	}
}

func (u *UIState) Run() error {
	folder := tview.NewBox().SetBorder(true).SetTitle("Folder")
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetChangedFunc(func() {
			u.app.Draw()
		})
	textView.SetBorder(true)
	// notes := tview.NewBox().SetBorder(true).SetTitle("Notes")
	list := tview.NewList().ShowSecondaryText(false)
	list.SetBorder(true)
	list.SetTitle("Notes")

	flex := tview.NewFlex().
		AddItem(folder, 0, 1, true).
		AddItem(list, 0, 1, true).
		AddItem(textView, 0, 3, true)
	flex.SetBackgroundColor(tcell.Color101)
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			u.app.Stop()
		}

		// if event.Key() == tcell.KeyCtrlC || event.Key() == tcell.KeyCtrlZ {
		// 	return nil
		// }

		if event.Key() == tcell.KeyCtrlA {
			for index, i := range []string{"1", "2", "3"} {
				list.AddItem(i, "", rune(49+index), nil)
			}
		}

		if event.Key() == tcell.KeyCtrlO {
			list.Clear()
		}

		return nil
	})

	u.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			u.app.Stop()
		}

		if event.Key() == tcell.KeyLeft {
			if u.flexFocusIndex > 0 && u.flexFocusIndex <= 2 {
				u.flexFocusIndex -= 1
			}
		}

		if event.Key() == tcell.KeyRight {
			if (u.flexFocusIndex >= 0) && (u.flexFocusIndex < 2) {
				u.flexFocusIndex += 1
			}
		}

		switch u.flexFocusIndex {
		case 0:
			u.app.SetFocus(folder)
		case 1:
			u.app.SetFocus(list)
		case 2:
			u.app.SetFocus(textView)
		}

		return nil
	})
	if err := u.app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		return err
	}
	return nil
}
