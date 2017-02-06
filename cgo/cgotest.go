package main

//#include <stdio.h>
//#include <stdlib.h>
/*
void printText(char* str){
	printf("%s\n", str);
}
*/
import "C"
import "unsafe"

func main() {
	str := "hello"
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.printText(cstr)
}
