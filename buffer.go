package nk

// #include "nk.h"
// #include <stdlib.h>
import "C"

import "unsafe"

// Buffer is an opaque handle to nk_buffer, which holds a defined and possibly
// expandable amount of memory. As with Context, the only safe way to create a
// Buffer is with NewBuffer.
type Buffer C.struct_nk_buffer

// NewBuffer creates and initializes a new Buffer in C memory and returns a
// pointer to it. The memory should be released with Free when no longer needed.
// Under the hood, NewBuffer calls nk_buffer_init_default.
func NewBuffer() *Buffer {
	ptr := (*C.struct_nk_buffer)(C.malloc(C.sizeof_struct_nk_buffer))
	// void nk_buffer_init_default(struct nk_buffer*);
	C.nk_buffer_init_default(ptr)
	return (*Buffer)(ptr)
}

// Free releases the memory used by b including b itself. Free is nil-safe.
// After the call, if b was non-nil, it is now a dangling pointer.
func (b *Buffer) Free() {
	if b != nil {
		// void nk_buffer_free(struct nk_buffer*)
		C.nk_buffer_free(b.raw())
		C.free(unsafe.Pointer(b))
	}
}

// Clear clears the contents of a buffer without releasing its memory.
func (b *Buffer) Clear() {
	// void nk_buffer_clear(struct nk_buffer*);
	C.nk_buffer_clear(b.raw())
}

// Memory returns the "needed" memory held by b as a slice of bytes. Despite
// appearances, the returned slice is not really Go-managed memory and may
// contain a dangling pointer if the buffer is altered after retrieving it.
// Generally speaking, the return value of Memory should only be held until
// it has been fully consumed, and then re-obtained for future use.
func (b *Buffer) Memory() []byte {
	ptr, size := b.MemoryUnsafe()
	return fakeByteSlice(ptr, size)
}

// MemoryUnsafe is the less-safe version of Memory. It returns a pointer to the
// first byte and the size of the "needed" portion. The same caveats about the
// validity of ptr apply as for Memory.
func (b *Buffer) MemoryUnsafe() (ptr unsafe.Pointer, size int) {
	// void *nk_buffer_memory(struct nk_buffer*);
	return C.nk_buffer_memory(b.raw()), int(b.needed)
}

func (b *Buffer) raw() *C.struct_nk_buffer {
	return (*C.struct_nk_buffer)(b)
}
