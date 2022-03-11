package nk

// #include "nk.h"
import "C"

type Buffer struct {
	raw C.struct_nk_buffer
}

func (b *Buffer) InitDefault() {
	// void nk_buffer_init_default(struct nk_buffer*);
	C.nk_buffer_init_default(&b.raw)
}

func (b *Buffer) Clear() {
	// void nk_buffer_clear(struct nk_buffer*);
	C.nk_buffer_clear(&b.raw)
}

func (b *Buffer) Free() {
	// void nk_buffer_free(struct nk_buffer*)
	C.nk_buffer_free(&b.raw)
}
