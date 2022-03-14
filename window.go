package nk

// #include "nk.h"
// #include <stdlib.h>
import "C"

import "unsafe"

func (ctx *Context) Begin(title string, bounds *Rect, flags Flags) bool {
	rawTitle := C.CString(title)
	defer C.free(unsafe.Pointer(rawTitle))
	//nk_bool nk_begin(struct nk_context *ctx, const char *title, struct nk_rect bounds, nk_flags flags);
	return bool(C.nk_begin(
		ctx.raw(),
		rawTitle,
		C.nk_rect(
			C.float(bounds.X),
			C.float(bounds.Y),
			C.float(bounds.W),
			C.float(bounds.H),
		),
		C.nk_flags(flags),
	))
}

func (ctx *Context) BeginTitled(name, title string, bounds *Rect, flags Flags) bool {
	rawName := C.CString(name)
	defer C.free(unsafe.Pointer(rawName))
	rawTitle := C.CString(title)
	defer C.free(unsafe.Pointer(rawTitle))
	// nk_bool nk_begin_titled(struct nk_context *ctx, const char *name, const char *title, struct nk_rect bounds, nk_flags flags)
	return bool(C.nk_begin_titled(
		ctx.raw(),
		rawName,
		rawTitle,
		C.nk_rect(
			C.float(bounds.X),
			C.float(bounds.Y),
			C.float(bounds.W),
			C.float(bounds.H),
		),
		C.nk_flags(flags),
	))
}

func (ctx *Context) End() {
	// void nk_end(struct nk_context *ctx);
	C.nk_end(ctx.raw())
}
