//+build windows,!server

package tainted

import (
	glog "log"
	"os"
	"syscall"

	"nox/v1/common/log"
)

func init() {
	// attach stdout/stderr to the console of the parent process
	// this way we can keep building the main binary as GUI App,
	// but still allow seeing console if started via cmd
	const ATTACH_PARENT_PROCESS = ^uint32(0) // (DWORD)-1
	modkernel32 := syscall.NewLazyDLL("kernel32.dll")
	procAttachConsole := modkernel32.NewProc("AttachConsole")

	r1, _, errno := syscall.Syscall(procAttachConsole.Addr(), 1, uintptr(ATTACH_PARENT_PROCESS), 0, 0)
	if errno != 0 || r1 == 0 {
		return
	}
	hout, err := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	if err != nil {
		return
	}
	herr, err := syscall.GetStdHandle(syscall.STD_ERROR_HANDLE)
	if err != nil {
		return
	}
	os.Stdout = os.NewFile(uintptr(hout), "/dev/stdout")
	os.Stderr = os.NewFile(uintptr(herr), "/dev/stderr")
	glog.SetOutput(os.Stderr)
	log.SetOutput(os.Stderr)
}
