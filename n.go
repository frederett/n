package main

import "github.com/ayakahokari/n/ui"

func main() {
	ui := ui.NewUIState()
	if err := ui.Run(); err != nil {
		panic(err)
	}
}
