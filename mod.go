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

func console_printf(t string, args ...any) {
	t = "[Proper Borderless] " + t + "\n"
	if len(args) == 0 {
		fmt.Print(t)
	} else {
		fmt.Printf(t, args...)
	}

}

var conf_path string
var hwnd syscall.Handle

//export Init
func Init() {
	dir, _ := os.Getwd()
	conf_path = path.Join(dir, "config.toml")
	console_printf("using config file: %s", conf_path)
	go run(false)
}

//export OnResize
func OnResize() {
	go run(true)
}

const interval = 1000 * 1000 * 100

func run(is_resize bool) {
	var conf Config
	toml.DecodeFile(conf_path, &conf)

	if is_resize && !conf.OnResize {
		console_printf("skipping OnResize because it's not enabled")
		return
	}

	// only search for hwnd when global hwnd is not set
	if int(hwnd) == 0 {
		pid := os.Getpid()
		h_con, err := win32tools.GetConsoleWindow()
		if err == nil {
			con_title, err := win32tools.GetWindowText(h_con, 200)
			if err == nil {
				console_printf("found console window %d, title: %s", h_con, con_title)
			} else {
				console_printf("unable to get title for console window %d: %s", h_con, err)
			}
		}
		var i int8
		console_printf("searching windows for pid %d", pid)
		for i < 100 {
			hwnd = findDivaHwnd(pid, h_con)
			if int(hwnd) == 0 {
				time.Sleep(interval) //sleep 0.1s
				i++
			} else {
				console_printf("found game window after %d retries", i)
				break
			}
		}
	}
	if !win32tools.IsWindowBorderless(hwnd) {
		console_printf("setting hwnd %d with (%d,%d,%d,%d)", int(hwnd), conf.PositionX, conf.PositionY, conf.Width, conf.Height)
		win32tools.SetBorderless(hwnd, conf.PositionX, conf.PositionY, conf.Width, conf.Height)
	} else {
		console_printf("skipping hwnd %d because it's already borderless", int(hwnd))
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
			console_printf("get title for %d failed, cause: %s", int(h), err)
			return
		}
		console_printf("using hwnd %d, title: %s", int(h), title)
		result = h
		return

	}
	return
}

func main() {
}
