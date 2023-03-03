package main

import "C"

import (
	"github.com/BurntSushi/toml"
	"github.com/deorth-kku/diva-proper-borderless/win32tools"
	"os"
	"path"
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

//export Init
func Init() {
	dir, _ := os.Getwd()
	conf_path := path.Join(dir, "config.toml")
	go run(conf_path, false)
}

//export OnResize
func OnResize() {
	dir, _ := os.Getwd()
	conf_path := path.Join(dir, "config.toml")
	go run(conf_path, true)
}

func run(conf_path string, is_resize bool) {
	pid := os.Getpid()
	var conf Config
	toml.DecodeFile(conf_path, &conf)

	if is_resize && !conf.OnResize {
		return
	}

	var hwnd syscall.Handle
	var i uint8
	for i < 100 {
		hwnd, _ = win32tools.FindWindowByPid(pid)
		if int(hwnd) == 0 {
			time.Sleep(1000 * 1000 * 100)
			i++
		} else {
			break
		}

	}
	if !win32tools.IfWindowBorderless(hwnd) {
		win32tools.SetBorderless(hwnd, conf.PositionX, conf.PositionY, conf.Width, conf.Height)
	}
}

func main() {
}
