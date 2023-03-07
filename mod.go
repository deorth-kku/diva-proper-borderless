package main

import "C"

import (
	"fmt"
	"os"
	"path"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/deorth-kku/diva-proper-borderless/win32tools"
)

type Config struct {
	PositionX int32
	PositionY int32
	Width     int32
	Height    int32
	OnResize  bool
}

const interval = 1000 * 1000 * 100

var conf_path string
var hwnd syscall.Handle

//export Init
func Init() {
	dir, _ := os.Getwd()
	conf_path = path.Join(dir, "config.toml")
	fmt.Printf("[Proper Borderless] using config file: %s\n", conf_path)
	go run(false)
}

//export OnResize
func OnResize() {
	go run(true)
}

func run(is_resize bool) {
	var conf Config
	toml.DecodeFile(conf_path, &conf)

	if is_resize && !conf.OnResize {
		fmt.Print("[Proper Borderless] skipping OnResize because it's not enabled\n")
		return
	}

	// only search for hwnd when global hwnd is not set
	if int(hwnd) == 0 {
		pid := os.Getpid()
		h_con, err := win32tools.GetConsoleWindow()
		if err == nil {
			con_title, err := win32tools.GetWindowText(h_con, 200)
			if err == nil {
				fmt.Printf("[Proper Borderless] find console window %d, title: %s\n", h_con, con_title)
			} else {
				fmt.Printf("[Proper Borderless] unable to get title for console window %d: %s\n", h_con, err)
			}
		}
		var i int8
		fmt.Printf("[Proper Borderless] searching windows for pid %d\n", pid)
		for i < 100 {
			hwnd = findDivaHwnd(pid, h_con)
			if int(hwnd) == 0 {
				time.Sleep(interval) //sleep 0.1s
				i++
			} else {
				break
			}
		}
	}
	if !win32tools.IsWindowBorderless(hwnd) {
		fmt.Printf("[Proper Borderless] setting hwnd %d with (%d,%d,%d,%d)\n", int(hwnd), conf.PositionX, conf.PositionY, conf.Width, conf.Height)
		win32tools.SetBorderless(hwnd, conf.PositionX, conf.PositionY, conf.Width, conf.Height)
	} else {
		fmt.Printf("[Proper Borderless] skipping hwnd %d because it's already borderless\n", int(hwnd))
	}
}

func findDivaHwnd(pid int, console_hwnd syscall.Handle) (result syscall.Handle) {
	hwnd_list, err := win32tools.FindWindowByPid(pid)
	if err != nil {
		return
	}
	for _, h := range hwnd_list {
		if h == console_hwnd {
			continue
		}
		title, err := win32tools.GetWindowText(h, 200)
		if err != nil {
			fmt.Printf("[Proper Borderless] get title for %d failed, cause: %s\n", int(h), err)
			return
		}
		fmt.Printf("[Proper Borderless] using hwnd %d, title: %s\n", int(h), title)
		result = h
		return

	}
	return
}

func main() {
}
