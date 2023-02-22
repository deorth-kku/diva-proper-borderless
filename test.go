package main

import (
	"fmt"

	"github.com/deorth-kku/diva-proper-borderless/win32tools"
)

func main() {
	pid := 5240
	hwnd, _ := win32tools.FindWindowByPid(pid)
	fmt.Printf("%d\n", hwnd)
	win32tools.SetBorderless(hwnd, 0, 0, 1920, 1080)

}
