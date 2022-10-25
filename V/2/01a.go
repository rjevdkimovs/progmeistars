package main

import  (
	"fmt"
	"strings"
	"unsafe"
)	

type (
	stringHeader struct {
		start   uintptr
		length  uint
	}
)	

func main() {  
    str:= strings.Repeat("abc",4) 
	fmt.Println(str)                                    // abcabcabcabc
    fmt.Println(&str)                                   // 0xc00002c1f0
    fmt.Println(unsafe.Pointer(&str))                   // 0xc00002c1f0
    fmt.Println(unsafe.Sizeof(str))                     // 16
	strarr := (*[2]uintptr)(unsafe.Pointer(&str))  
	fmt.Println(*strarr)                                // [824633778352 12]
	fmt.Printf("%x %d\n", (*strarr)[0], (*strarr)[1])   // c00000e0b0 12
	// all together - struct
	fmt.Println(*(*stringHeader)(unsafe.Pointer(&str))) // {824633778352 12}
	str2:= str
    fmt.Println(&str2)                                  // 0xc00002c210
	fmt.Println(*(*stringHeader)(unsafe.Pointer(&str2)))// {824633778352 12}
}
