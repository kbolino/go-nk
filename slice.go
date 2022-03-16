package nk

import (
	"reflect"
	"unsafe"
)

func fakeSlice[T any](ptr *T, length int) []T {
	fakeSliceHeader := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ptr)),
		Len:  length,
		Cap:  length,
	}
	return *(*[]T)(unsafe.Pointer(&fakeSliceHeader))
}
