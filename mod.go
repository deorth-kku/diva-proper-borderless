package main

import "C"

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/deorth-kku/diva-proper-borderless/win32tools"
	"os"
	"path"
	"strings"
	"syscall"
	"time"
)

type Config struct {
	PositionX int32
	PositionY int32
	Width     int32
	Height    int32
	OnResize  bool
}

var conf_path string
var pid int
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
		var i int8
		fmt.Printf("[Proper Borderless] searching windows for pid %d\n", pid)
		for i < 100 {
			hwnd = findDivaHwnd(pid)
			if int(hwnd) == 0 {
				time.Sleep(1000 * 1000 * 100) //sleep 0.1s
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

func findDivaHwnd(pid int) (result syscall.Handle) {
	hwnd_list, err := win32tools.FindWindowByPid(pid)
	if err != nil {
		return
	}
	for _, h := range hwnd_list {
		title, err := win32tools.GetWindowText(h, 200)
		if err != nil {
			fmt.Printf("[Proper Borderless] get title for %d failed, cause: %s\n", int(h), err)
			return
		}
		// search for window that title ends with "+". I know this is stupid but I cannot think of anything else
		if strings.HasSuffix(title, "+") {
			fmt.Printf("[Proper Borderless] using hwnd %d, title: %s\n", int(h), title)
			result = h
			return
		} else {
			fmt.Printf("[Proper Borderless] skipping hwnd %d, title: %s\n", int(h), title)
		}
	}
	return
}

func main() {
}
