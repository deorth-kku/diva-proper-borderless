package win32tools

import (
	"syscall"
	"unsafe"
)

var (
	procGetWindowLongW = user32.MustFindProc("GetWindowLongW")
	procGetWindowInfo  = user32.MustFindProc("GetWindowInfo")
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

func IfWindowBorderless(hwnd syscall.Handle) bool {
	style := GetWindowLong(hwnd, GWL_STYLE)
	ex_style := GetWindowLong(hwnd, GWL_EXSTYLE)
	return (style >= WS_VISIBLE|WS_CLIPCHILDREN) && (ex_style == 0)
}
