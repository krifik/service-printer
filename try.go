// package main

// import (
// 	"fmt"
// 	"log"
// 	"syscall"
// 	"time"
// 	"unsafe"
// )

// const (
// 	MEM_COMMIT  = 0x1000
// 	MEM_RESERVE = 0x2000
// 	MEM_RELEASE = 0x8000
// 	MEM_RESET   = 0x80000

// 	PAGE_NOACCESS          = 0x0001
// 	PAGE_EXECUTE_READWRITE = 0x40

// 	B  = 1
// 	KB = 1024 * B
// 	MB = 1024 * KB
// 	GB = 1024 * MB
// )

// var (
// 	dll          = syscall.("kernel32.dll")
// 	VirtualAlloc = dll.MustFindProc("VirtualAlloc")
// 	VirtualFree  = dll.MustFindProc("VirtualFree")
// )

// func SysAlloc(n uintptr) (uintptr, error) {
// 	addr, _, err := VirtualAlloc.Call(0, n, MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)
// 	if addr == 0 {
// 		return 0, err
// 	}
// 	return addr, nil
// }

// func SysUnused(p, n uintptr) error {
// 	r, _, err := VirtualAlloc.Call(p, n, MEM_RESET, PAGE_NOACCESS)
// 	if r == 0 {
// 		return err
// 	}
// 	return nil
// }

// func SysFree(p uintptr) error {
// 	r, _, err := VirtualFree.Call(p, 0, MEM_RELEASE)
// 	if r == 0 {
// 		return err
// 	}
// 	return nil
// }

// func firstByte(p uintptr) *byte {
// 	return (*byte)(unsafe.Pointer(p))
// }

// func touch(p uintptr) {
// 	f := firstByte(p)
// 	fmt.Printf("reading p=%v first=%d ...", p, *f)
// 	*f = 1
// 	fmt.Printf(" DONE first=%d\n", *f)
// }

// func pause(msg string) {
// 	fmt.Printf("%s ...\n", msg)
// 	time.Sleep(10 * time.Second)
// }

// func main() {
// 	const n = 200 * MB
// 	pause("before alloc")
// 	p, err := SysAlloc(n)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	touch(p)

// 	pause("before unused")
// 	err = SysUnused(p, n)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	touch(p)

// 	pause("before free")
// 	err = SysFree(p)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	pause("after free")
// }
