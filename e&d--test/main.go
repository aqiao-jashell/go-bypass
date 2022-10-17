package main

import (
	bases "bypass/bagua"
	"fmt"
)

func main() {
	fmt.Println("-------------------")
	fmt.Println("cs-shellcode加解密")

	var str = "cs-shellcode加解密"
	strbytes := []byte(str)
	red := bases.Bagua_en(strbytes)
	fmt.Println("加密", red)

	ref := bases.Bagua_de(red)
	fmt.Println("解密", string(ref))
}
