package nk

// #include "nk.h"
// #include <stdlib.h>
import "C"

// Begin is equivalent to BeginTitled(title, title, bounds, flags).
func (ctx *Context) Begin(title string, bounds *Rect, flags Flags) bool {
	// nk_begin calls nk_begin_titled under the hood with name == title
	rawTitle := cStringPool.Get(title)
	defer cStringPool.Release(rawTitle)
	// nk_bool nk_begin(struct nk_context *ctx, const char *title, struct nk_rect bounds, nk_flags flags);
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

// BeginTitled starts a new window. If the window already existed, it will be
// carried over from the previous frame with its old size and state. The return
// value indicates whether the window is visible; if not, no layouts or widgets
// should be added to it. BeginTitled must be matched with a call to End once
// all the layouts and widgets have been added.
func (ctx *Context) BeginTitled(name, title string, bounds *Rect, flags Flags) bool {
	rawName := cStringPool.Get(name)
	// releasing name is okay since the context will copy the name if needed
	defer cStringPool.Release(rawName)
	rawTitle := cStringPool.Get(title)
	// releasing title is okay since the context doesn't keep the title
	defer cStringPool.Release(rawTitle)
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

// End completes the current window, which must have been started with Begin
// or BeginTitled.
func (ctx *Context) End() {
	// void nk_end(struct nk_context *ctx);
	C.nk_end(ctx.raw())
}
