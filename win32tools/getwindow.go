package win32tools

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	procGetWindowLongW           = user32.MustFindProc("GetWindowLongW")
	procGetWindowInfo            = user32.MustFindProc("GetWindowInfo")
	procGetWindowTextW           = user32.MustFindProc("GetWindowTextW")
	procGetWindowThreadProcessId = user32.MustFindProc("GetWindowThreadProcessId")
	procIsWindowVisible          = user32.MustFindProc("IsWindowVisible")
	kernel32                     = syscall.MustLoadDLL("kernel32.dll")
	procGetConsoleWindow         = kernel32.MustFindProc("GetConsoleWindow")
)

type DWORD = uint32
type RECT struct {
	Left, Top, Right, Bottom int32
}
type UINT = uint
type ATOM = uint16
type WORD = uint16
type WINDOWINFO struct {
	CbSize          DWORD
	RcWindow        RECT
	RcClient        RECT
	DwStyle         DWORD
	DwExStyle       DWORD
	DwWindowStatus  DWORD
	CxWindowBorders UINT
	CyWindowBorders UINT
	AtomWindowType  ATOM
	WCreatorVersion WORD
}

func GetWindowLong(hwnd syscall.Handle, index int32) int {
	ret, _, _ := syscall.SyscallN(procGetWindowLongW.Addr(), uintptr(hwnd), uintptr(index))
	return int(ret)
}

func GetWindowInfo(hwnd syscall.Handle, info *WINDOWINFO) int {
	ret, _, _ := syscall.SyscallN(procGetWindowInfo.Addr(), uintptr(hwnd), uintptr(unsafe.Pointer(info)))
	return int(ret)
}

func GetWindowText(hwnd syscall.Handle, maxCount int32) (title string, err error) {
	str := make([]uint16, maxCount)
	r0, _, e1 := syscall.SyscallN(procGetWindowTextW.Addr(), uintptr(hwnd), uintptr(unsafe.Pointer(&str[0])), uintptr(maxCount))
	len := int32(r0)
	if len == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	title = syscall.UTF16ToString(str)
	return
}
func GetWindowThreadProcessId(hwnd syscall.Handle) (pid int, err error) {
	r0, _, e1 := syscall.SyscallN(procGetWindowThreadProcessId.Addr(), uintptr(hwnd), uintptr(unsafe.Pointer(&pid)))
	len := int32(r0)
	if len == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func GetConsoleWindow() (hwnd syscall.Handle, err error) {
	h, _, e1 := syscall.SyscallN(procGetConsoleWindow.Addr())
	if e1 != 0 {
		err = error(e1)
	}
	hwnd = syscall.Handle(h)
	return
}

func IsWindowVisible(hwnd syscall.Handle) bool {
	ret, _, _ := syscall.SyscallN(procIsWindowVisible.Addr(), uintptr(hwnd))

	return ret != 0
}

func IsWindowBorderless(hwnd syscall.Handle) bool {
	style := GetWindowLong(hwnd, GWL_STYLE)
	ex_style := GetWindowLong(hwnd, GWL_EXSTYLE)
	return (style&WS_VISIBLE == WS_VISIBLE) && (style&WS_CLIPCHILDREN == WS_CLIPCHILDREN) && (ex_style == 0)
}

func FindWindow(title string) (syscall.Handle, error) {
	var hwnd syscall.Handle
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		t, err := GetWindowText(h, 200)
		if err != nil {
			// ignore the error
			return 1 // continue enumeration
		}
		if t == title {
			// note the window
			hwnd = h
			return 0 // stop enumeration
		}
		return 1 // continue enumeration
	})
	EnumWindows(cb, 0)
	if hwnd == 0 {
		return 0, fmt.Errorf("no window with title '%s' found", title)
	}
	return hwnd, nil
}

func FindWindowByPid(pid int) (hwnd_list []syscall.Handle, err error) {
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		if !IsWindowVisible(h) {
			return 1
		}

		b, err := GetWindowThreadProcessId(h)
		if err != nil {
			// ignore the error
			return 1 // continue enumeration
		}
		if b == pid {
			// note the window
			hwnd_list = append(hwnd_list, h)
			return 1
		}
		return 1 // continue enumeration
	})
	EnumWindows(cb, 0)
	if len(hwnd_list) == 0 {
		err = fmt.Errorf("no window with pid '%d' found", pid)
	}
	return
}
