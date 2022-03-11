package nk

// #include "nk.h"
import "C"

import "unsafe"

type Buffer struct {
	raw C.struct_nk_buffer
	mem []byte
}

func (b *Buffer) InitFixed(memSize int) {
	mem := make([]byte, memSize)
	// void nk_buffer_init_fixed(struct nk_buffer*, void *memory, nk_size size);
	C.nk_buffer_init_fixed(&b.raw, unsafe.Pointer(&mem[0]), C.nk_size(memSize))
}

func (b *Buffer) Clear() {
	C.nk_buffer_clear(&b.raw)
}

func (b *Buffer) Free() {
	if b.mem != nil {
		b.mem = nil
	} else {
		C.nk_buffer_free(&b.raw)
	}
}
