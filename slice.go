package nk

import (
	"reflect"
	"unsafe"
)

func fakeByteSlice(ptr unsafe.Pointer, length int) []byte {
	fakeSliceHeader := reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  length,
		Cap:  length,
	}
	return *(*[]byte)(unsafe.Pointer(&fakeSliceHeader))
}

func fakePointsSlice(ptr unsafe.Pointer, length int) []Vec2i {
	fakeSliceHeader := reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  length,
		Cap:  length,
	}
	return *(*[]Vec2i)(unsafe.Pointer(&fakeSliceHeader))
}
