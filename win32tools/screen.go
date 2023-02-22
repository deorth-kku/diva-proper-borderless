package win32tools

import "syscall"

const (
	SM_CXSCREEN = 0
	SM_CYSCREEN = 1
)

var (
	procGetSystemMetrics = user32.MustFindProc("GetSystemMetrics")
)

func GetScreenSize() (int32, int32) {
	r0, _, _ := syscall.SyscallN(procGetSystemMetrics.Addr(), SM_CXSCREEN)
	r1, _, _ := syscall.SyscallN(procGetSystemMetrics.Addr(), SM_CYSCREEN)
	return int32(r0), int32(r1)
}
