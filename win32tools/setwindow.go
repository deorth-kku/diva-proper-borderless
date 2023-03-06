package win32tools

import (
	"syscall"
)

var (
	user32             = syscall.MustLoadDLL("user32.dll")
	procEnumWindows    = user32.MustFindProc("EnumWindows")
	procSetWindowLongW = user32.MustFindProc("SetWindowLongW")
	procMoveWindow     = user32.MustFindProc("MoveWindow")
	procSetWindowPos   = user32.MustFindProc("SetWindowPos")
)

const (
	SWP_DRAWFRAME      = 0x0020
	SWP_FRAMECHANGED   = 0x0020
	SWP_HIDEWINDOW     = 0x0080
	SWP_NOACTIVATE     = 0x0010
	SWP_NOCOPYBITS     = 0x0100
	SWP_NOMOVE         = 0x0002
	SWP_NOSIZE         = 0x0001
	SWP_NOREDRAW       = 0x0008
	SWP_NOZORDER       = 0x0004
	SWP_SHOWWINDOW     = 0x0040
	SWP_NOOWNERZORDER  = 0x0200
	SWP_NOREPOSITION   = SWP_NOOWNERZORDER
	SWP_NOSENDCHANGING = 0x0400
	SWP_DEFERERASE     = 0x2000
	SWP_ASYNCWINDOWPOS = 0x4000
)

const (
	GWL_EXSTYLE     = -20
	GWL_STYLE       = -16
	GWL_WNDPROC     = -4
	GWLP_WNDPROC    = -4
	GWL_HINSTANCE   = -6
	GWLP_HINSTANCE  = -6
	GWL_HWNDPARENT  = -8
	GWLP_HWNDPARENT = -8
	GWL_ID          = -12
	GWLP_ID         = -12
	GWL_USERDATA    = -21
	GWLP_USERDATA   = -21
)

const (
	WS_OVERLAPPED       = 0x00000000
	WS_POPUP            = 0x80000000
	WS_CHILD            = 0x40000000
	WS_MINIMIZE         = 0x20000000
	WS_VISIBLE          = 0x10000000
	WS_DISABLED         = 0x08000000
	WS_CLIPSIBLINGS     = 0x04000000
	WS_CLIPCHILDREN     = 0x02000000
	WS_MAXIMIZE         = 0x01000000
	WS_CAPTION          = 0x00C00000
	WS_BORDER           = 0x00800000
	WS_DLGFRAME         = 0x00400000
	WS_VSCROLL          = 0x00200000
	WS_HSCROLL          = 0x00100000
	WS_SYSMENU          = 0x00080000
	WS_THICKFRAME       = 0x00040000
	WS_GROUP            = 0x00020000
	WS_TABSTOP          = 0x00010000
	WS_MINIMIZEBOX      = 0x00020000
	WS_MAXIMIZEBOX      = 0x00010000
	WS_TILED            = 0x00000000
	WS_ICONIC           = 0x20000000
	WS_SIZEBOX          = 0x00040000
	WS_OVERLAPPEDWINDOW = 0x00000000 | 0x00C00000 | 0x00080000 | 0x00040000 | 0x00020000 | 0x00010000
	WS_POPUPWINDOW      = 0x80000000 | 0x00800000 | 0x00080000
	WS_CHILDWINDOW      = 0x40000000
)

func EnumWindows(enumFunc uintptr, lparam uintptr) (err error) {
	r1, _, e1 := syscall.SyscallN(procEnumWindows.Addr(), uintptr(enumFunc), uintptr(lparam))
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetWindowLong(hwnd syscall.Handle, index, value int32) int32 {
	ret, _, _ := syscall.SyscallN(procSetWindowLongW.Addr(),
		uintptr(hwnd),
		uintptr(index),
		uintptr(value))

	return int32(ret)
}

func boolToInt(value bool) int32 {
	if value {
		return 1
	}
	return 0
}

func MoveWindow(hwnd syscall.Handle, x, y, width, height int32, repaint bool) bool {
	ret, _, _ := syscall.SyscallN(procMoveWindow.Addr(),
		uintptr(hwnd),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(boolToInt(repaint)))

	return ret != 0
}

func SetWindowPos(hwnd syscall.Handle, hwndInsertAfter syscall.Handle, x, y, width, height int32, flags uint32) bool {
	ret, _, _ := syscall.SyscallN(procSetWindowPos.Addr(),
		uintptr(hwnd),
		uintptr(hwndInsertAfter),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(flags))

	return ret != 0
}
func intToBool(value int32) bool {
	var a bool
	if value != 0 {
		a = true
	}
	return a
}

func SetBorderless(hwnd syscall.Handle, XPos, YPos, HRes, VRes int32) bool {
	if HRes == 0 && VRes == 0 {
		HRes, VRes = GetScreenSize()
	}
	a := intToBool(SetWindowLong(hwnd, GWL_STYLE,
		WS_VISIBLE|WS_CLIPCHILDREN))
	b := intToBool(SetWindowLong(hwnd, GWL_EXSTYLE, 0))
	c := MoveWindow(hwnd, XPos, YPos, HRes-1, VRes-1, true)
	var hwnd0 syscall.Handle
	d := SetWindowPos(hwnd, hwnd0, XPos, YPos, HRes, VRes, SWP_FRAMECHANGED|SWP_NOZORDER|SWP_NOOWNERZORDER)
	return a && b && c && d

}
