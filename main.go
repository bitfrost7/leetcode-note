package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 1
	b := int8(1)
	c := int16(1)
	d := int32(1)
	e := int64(1)
	u_a := uint(1)
	u_b := uint8(1)
	u_c := uint16(1)
	u_d := uint32(1)
	u_e := uint64(1)
	fmt.Println("int类型大小:", unsafe.Sizeof(a), "字节")
	fmt.Println("int8类型大小:", unsafe.Sizeof(b), "字节")
	fmt.Println("int16类型大小:", unsafe.Sizeof(c), "字节")
	fmt.Println("int32类型大小:", unsafe.Sizeof(d), "字节")
	fmt.Println("int64类型大小:", unsafe.Sizeof(e), "字节")
	fmt.Println("uint类型大小:", unsafe.Sizeof(u_a), "字节")
	fmt.Println("uint8类型大小:", unsafe.Sizeof(u_b), "字节")
	fmt.Println("uint16类型大小:", unsafe.Sizeof(u_c), "字节")
	fmt.Println("uint32类型大小:", unsafe.Sizeof(u_d), "字节")
	fmt.Println("uint64类型大小:", unsafe.Sizeof(u_e), "字节")
	type mystruct struct {
		a int8
		d int64
		c int32
		b int16
	}
	f := mystruct{}
	fmt.Println("mystruct类型大小:", unsafe.Sizeof(f), "字节")
}
