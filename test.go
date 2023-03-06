package main

import (
	"fmt"

	"github.com/deorth-kku/diva-proper-borderless/win32tools"
)

func main() {
	name := "初音ミク Project DIVA Mega39's+"
	hwnd, _ := win32tools.FindWindow(name)
	if win32tools.IsWindowBorderless(hwnd) {
		fmt.Printf("window %s is borderless\n", name)
	} else {
		fmt.Printf("window %s is not borderless\n", name)
	}
	var pids = [3]int{8988, 5532, 8000}
	for _, pid := range pids {
		hwnd_list, _ := win32tools.FindWindowByPid(pid)
		for _, hwnd := range hwnd_list {
			title, _ := win32tools.GetWindowText(hwnd, 200)
			fmt.Printf("pid: %d, hwnd: %d, title: %s\n", pid, int(hwnd), title)
		}
	}
}
