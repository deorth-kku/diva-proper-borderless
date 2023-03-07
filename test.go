package main

import (
	"fmt"
	"time"

	"github.com/deorth-kku/diva-proper-borderless/win32tools"
)

func main() {
	name := "初音ミク Project DIVA Mega39's+"
	h_con, _ := win32tools.GetConsoleWindow()
	title, err := win32tools.GetWindowText(h_con, 200)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("my console is %d, title: %s\n", h_con, title)
	hwnd, _ := win32tools.FindWindow(name)
	if win32tools.IsWindowBorderless(hwnd) {
		fmt.Printf("window %s is borderless\n", name)
	} else {
		fmt.Printf("window %s is not borderless\n", name)
	}
	time.Sleep(1000 * 1000 * 1000 * 10)

}
