package nk

// #include "nk.h"
import "C"

import (
	"fmt"
	"unsafe"
)

type UserFont struct {
	raw C.struct_nk_user_font
}

type Context struct {
	raw C.struct_nk_context
	mem []byte // go reference to keep fixed memory alive
}

func (ctx *Context) InitFixed(memSize int) error {
	ctx.mem = make([]byte, memSize)
	// nk_bool nk_init_fixed(struct nk_context*, void *memory, nk_size size, const struct nk_user_font*);
	if !C.nk_init_fixed(&ctx.raw, unsafe.Pointer(&ctx.mem[0]), C.nk_size(memSize), nil) {
		return fmt.Errorf("nk_init_fixed returned false")
	}
	return nil
}

func (ctx *Context) Clear() {
	// void nk_clear(struct nk_context*);
	C.nk_clear(&ctx.raw)
}

func (ctx *Context) Free() {
	// void nk_free(struct nk_context*);
	C.nk_free(&ctx.raw)
	ctx.mem = nil
}

func (ctx *Context) ForEach(f func(Command) bool) bool {
	// const struct nk_command* nk__begin(struct nk_context*);
	// const struct nk_command* nk__next(struct nk_context*, const struct nk_command*);
	for cmd := C.nk__begin(&ctx.raw); cmd != nil; cmd = C.nk__next(&ctx.raw, cmd) {
		if !f(typeSwitchCommand(cmd)) {
			return false
		}
	}
	return true
}

// struct nk_color {nk_byte r,g,b,a;};

type Color struct{ R, G, B, A uint8 }

// struct nk_colorf {float r,g,b,a;};

type Colorf struct{ R, G, B, A float32 }

// struct nk_vec2 {float x,y;};

type Vec2 struct{ X, Y float32 }

// struct nk_vec2i {short x, y;};

type Vec2i struct{ X, Y int16 }

// struct nk_rect {float x,y,w,h;};

type Rect struct{ X, Y, W, H float32 }

// struct nk_recti {short x,y,w,h;};

type Recti struct{ X, Y, W, H int16 }

// typedef union {void *ptr; int id;} nk_handle;

type Handle uintptr

// struct nk_image {nk_handle handle; nk_ushort w, h; nk_ushort region[4];};

type Image struct {
	Handle Handle
	W, H   uint16
	Region [4]uint16
}

// struct nk_nine_slice {struct nk_image img; nk_ushort l, t, r, b;};

type NineSlice struct {
	Img        Image
	L, T, R, B uint16
}

// struct nk_cursor {struct nk_image img; struct nk_vec2 size, offset;};

type Cursor struct {
	Img          Image
	Size, Offset Vec2
}

// struct nk_scroll {nk_uint x, y;};

type Scroll struct {
	X, Y uint32
}
