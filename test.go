package main

import (
	"fmt"

	"github.com/deorth-kku/diva-proper-borderless/win32tools"
)

func main() {
	name := "初音ミク Project DIVA Mega39's+"
	hwnd, _ := win32tools.FindWindow(name)
	if win32tools.IfWindowBorderless(hwnd) {
		fmt.Printf("window %s is borderless\n", name)
	} else {
		fmt.Printf("window %s is not borderless\n", name)
	}

}
