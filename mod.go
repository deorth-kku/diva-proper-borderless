package main

import "C"

import (
	"github.com/deorth-kku/diva-proper-borderless/win32tools"
	"os"
	"syscall"
	"time"
)

func init() {
	pid := os.Getpid()
	var hwnd syscall.Handle
	var i uint8
	for int(hwnd) == 0 && i < 100 {
		time.Sleep(1000 * 1000 * 100)
		hwnd, _ = win32tools.FindWindowByPid(pid)
		i++
	}
	win32tools.SetBorderless(hwnd, 0, 0, 0, 0)
}

func main() {
}
