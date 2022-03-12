package nk

// #include "nk.h"
// #include <stdlib.h>
import "C"

import "unsafe"

type Buffer C.struct_nk_buffer

func NewBuffer() *Buffer {
	ptr := (*C.struct_nk_buffer)(C.malloc(C.sizeof_struct_nk_buffer))
	// void nk_buffer_init_default(struct nk_buffer*);
	C.nk_buffer_init_default(ptr)
	return (*Buffer)(ptr)
}

func (b *Buffer) Clear() {
	// void nk_buffer_clear(struct nk_buffer*);
	C.nk_buffer_clear(b.raw())
}

func (b *Buffer) Free() {
	if b != nil {
		// void nk_buffer_free(struct nk_buffer*)
		C.nk_buffer_free(b.raw())
		C.free(unsafe.Pointer(b))
	}
}

func (b *Buffer) Memory() []byte {
	ptr, size := b.MemoryUnsafe()
	return fakeByteSlice(ptr, size)
}

func (b *Buffer) MemoryUnsafe() (ptr unsafe.Pointer, size int) {
	// void *nk_buffer_memory(struct nk_buffer*);
	return C.nk_buffer_memory(b.raw()), int(b.needed)
}

func (b *Buffer) raw() *C.struct_nk_buffer {
	return (*C.struct_nk_buffer)(b)
}
