package clvq

import "C"
import (
	sc "golang.org/x/text/encoding/simplifiedchinese"
	"unsafe"
)

func cInt(i int) C.int {
	return C.int(i)
}

func goInt(ci C.int) int {
	return int(ci)
}

func cBool(b bool) C.int {
	if b {
		return 1
	}
	return 0
}

func cString(str string) *C.char {
	gbstr, _ := sc.GB18030.NewEncoder().String(str)
	return C.CString(gbstr)
}

func goString(str *C.char) string {
	utf8str, _ := sc.GB18030.NewDecoder().String(C.GoString(str))
	return utf8str
}

func str2ptr(s string) uintptr {
	return uintptr(unsafe.Pointer(cString(s)))
}

func int2ptr(i int) uintptr {
	return uintptr(i)
}

func byte2ptr(b []byte) uintptr {
	return uintptr(*((*int32)(unsafe.Pointer(&b))))
}

func bool2ptr(b bool) uintptr {
	if b == true {
		return uintptr(1)
	}
	return uintptr(0)
}

func ptr2str(ptr uintptr) string {
	return goString((*C.char)(unsafe.Pointer(ptr)))
}

func ptr2bool(ptr uintptr) bool {
	i := int(ptr)
	if i == 1 {
		return true
	}
	return false
}

func ptr2int(ptr uintptr) int {
	return int(ptr)
}

func _ptr2str(ptr uintptr) string {
	var b []byte
	for {
		sbyte := *((*byte)(unsafe.Pointer(ptr)))
		if sbyte == 0 {
			break
		}
		b = append(b, sbyte)
		ptr += 1
	}
	return string(b)
}
