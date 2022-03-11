package nk

import "C"

import "reflect"

func init() {
	// there doesn't seem to be any way to statically assert these equivalences,
	// so we have to do it at runtime instead

	if reflect.TypeOf(C.ushort(0)).Kind() != reflect.Uint16 {
		panic("unsigned short != uint16")
	}
	if reflect.TypeOf(C.uint(0)).Kind() != reflect.Uint32 {
		panic("unsigned int != uint32")
	}
}
