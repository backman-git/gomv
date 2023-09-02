package memwriter

import (
	"syscall"
	"unsafe"
)

type MemWriter struct {
	head    uintptr
	tail    uintptr
	enableW bool
	content []byte
}

var PAGESIZE = syscall.Getpagesize()

func NewMemWriter(head, tail uintptr) *MemWriter {

	return &MemWriter{head: head, tail: tail, enableW: false}
}

func getMemContent(head, tail uintptr) []byte {

	return (*(*[0xffffffff]byte)(unsafe.Pointer(head)))[:tail-head]

}
func (mw MemWriter) Write(headPtr unsafe.Pointer, content []byte) error {

	if mw.enableW == false {
		mw.EnableWrite()
	}

	mContent := getMemContent(mw.head, mw.tail)
	copy(mContent, content)
	return nil
}

func GetPageHead(addr uintptr) uintptr {
	return addr & (^uintptr(PAGESIZE) - 1)
}

func GetPage(addr uintptr) []byte {
	pageHeadAddr := GetPageHead(addr)
	return (*(*[0xffffffff]byte)(unsafe.Pointer(pageHeadAddr)))[:PAGESIZE]
}

func (mw MemWriter) EnableWrite() error {

	mPage := GetPage(mw.head)
	return syscall.Mprotect(mPage, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)
}
