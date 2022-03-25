package nk

import "C"

import (
	"fmt"
	"unsafe"
)

func init() {
	// there doesn't seem to be any way to statically assert these equivalences,
	// so we have to do it at runtime instead

	if sizeofShort := unsafe.Sizeof(C.short(0)); sizeofShort != 2 {
		panic(fmt.Errorf("type convertibility assertion failed: size of C short (%d) != size of Go int16 (2)",
			sizeofShort))
	}
	if sizeofInt := unsafe.Sizeof(C.int(0)); sizeofInt != 4 {
		panic(fmt.Errorf("type convertibility assertion failed: size of C int (%d) != size of Go int32 (4)",
			sizeofInt))
	}
	if sizeofSizeT, sizeofUint := unsafe.Sizeof(C.size_t(0)), unsafe.Sizeof(uint(0)); sizeofSizeT != sizeofUint {
		panic(fmt.Errorf("type convertibility assertion failed: size of C size_t (%d) != size of Go uint (%d)",
			sizeofSizeT, sizeofUint))
	}
}
