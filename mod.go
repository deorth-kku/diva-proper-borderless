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

//export D3DInit
func D3DInit() {
	dir, _ := os.Getwd()
	conf_path = path.Join(dir, "config.toml")
	pid = os.Getpid()
	fmt.Printf("[Proper Borderless] pid: %d\n", pid)
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

	var i uint8
	var hwnd_list []syscall.Handle
	for i < 100 {
		var err error
		hwnd_list, err = win32tools.FindWindowByPid(pid)
		if err != nil {
			time.Sleep(1000 * 1000 * 100)
			i++
		} else {
			break
		}
	}
	fmt.Printf("[Proper Borderless] found %d hwnds:", len(hwnd_list))
	for _, h := range hwnd_list {
		fmt.Printf("%d,", h)
	}
	fmt.Print("\n")

	var hwnd syscall.Handle
	for _, h := range hwnd_list {
		title, err := win32tools.GetWindowText(h, 200)
		if err != nil {
			fmt.Printf("[Proper Borderless] get title for %d failed, cause: %s\n", int(h), err)
		}
		if strings.HasSuffix(title, "+") {
			fmt.Printf("[Proper Borderless] using hwnd %d, title: %s\n", int(h), title)
			hwnd = h
			break
		} else {
			fmt.Printf("[Proper Borderless] skipping hwnd %d, title: %s\n", int(h), title)
			continue
		}
	}

	if !win32tools.IsWindowBorderless(hwnd) {
		fmt.Printf("[Proper Borderless] setting hwnd %d with (%d,%d,%d,%d)\n", int(hwnd), conf.PositionX, conf.PositionY, conf.Width, conf.Height)
		win32tools.SetBorderless(hwnd, conf.PositionX, conf.PositionY, conf.Width, conf.Height)
	} else {
		fmt.Printf("[Proper Borderless] skipping hwnd %d because it's already borderless\n", int(hwnd))
	}
}

func main() {
}
