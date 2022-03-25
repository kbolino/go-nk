// Package nk provides the Go API for Nuklear, an immediate-mode GUI library
// with optional GPU-accelerated rendering.
//
// Wherever possible, Go idioms are used to increase the comfortability of this
// library for Go natives. However, the method and type names are kept as close
// as possible to the C API to avoid confusion. It is strongly recommended to
// consult the Nuklear documentation first, or at least as a side-by-side
// reference: https://immediate-mode-ui.github.io/Nuklear/doc/index.html
//
// The entrypoint is the Context type, which is created with NewContext. As an
// immediate-mode library, rendering is done on a "frame-by-frame" basis and
// no state is kept about the windows, layouts, widgets, or other UI elements.
//
// Nuklear does not actually render the UI, it only tells another library how
// to render the UI. That other library can be lower level like OpenGL,
// Direct3D, Metal, Vulkan, mid-level like SDL2, or even something high-level
// like drawing an SVG image. As of SDL 2.0.18, it is possible to GPU accelerate
// Nuklear with the SDL2 renderer thanks to SDL_RenderGeometry and/or
// SDL_RenderGeometryRaw.
//
// At the start, there is an initialization phase, during which the fonts can
// be baked with FontAtlas, and the specific renderer should be created atop
// a library- or platform-specific window. Thereafter, the Context should be
// created and used throughout the life of the UI.
//
// There are three phases each frame: input, layout, and draw. The input phase
// is begun with Context.InputBegin and ended with Context.InputEnd; in between
// these two methods, other Input methods should be called to record input
// events received from the platform. The layout phase is begun with
// Context.Begin and ended with Context.End; in between these two methods,
// layout and widget methods should be called to set the UI contents. Finally,
// the draw phase begins implicitly after Context.End and lasts until
// Context.Clear; in between these two methods, Context.ForEach should be called
// once to iterate over the queued commands. Alternately, if GPU acceleration
// is desired, Context.Convert should be called followed by Context.DrawForEach,
// the former creating the vertex buffers and the latter pushing them to the
// GPU.
//
// Since Nuklear is at heart a C API and since cgo places restrictions on how
// freely Go pointers can be passed to C, many types are allocated in C memory.
// These types are explicitly noted in their documentation and need to be Free'd
// after use. The best way to do this is to simply defer the Free call. C-style
// strings can be pooled using the SetCStringPool function, see its
// documentation for more details.
package nk

// #include "nk.h"
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

// Context is an opaque handle to nk_context, the entrypoint for Nuklear.
// Always create a Context with NewContext, do not use a stack-allocated or
// Go heap-allocated value.
type Context C.struct_nk_context

// NewContext creates and initializes a new Context in C memory and returns a
// pointer to it. The memory should be released with Free when no longer needed.
// Under the hood, NewContext calls nk_init_default.
func NewContext() (*Context, error) {
	ptr := (*C.struct_nk_context)(C.malloc(C.sizeof_struct_nk_context))
	// nk_bool nk_init_default(struct nk_context*, const struct nk_user_font*);
	if !C.nk_init_default(ptr, nil) {
		return nil, fmt.Errorf("nk_init_default returned false")
	}
	return (*Context)(ptr), nil

}

// Free releases all memory used by ctx including ctx itself. Free is nil-safe.
// After the call, if ctx was not nil, it is now a dangling pointer.
func (ctx *Context) Free() {
	if ctx != nil {
		// void nk_free(struct nk_context*);
		C.nk_free(ctx.raw())
		C.free(unsafe.Pointer(ctx))
	}
}

// Clear calls nk_clear which resets the Context for the next frame.
func (ctx *Context) Clear() {
	// void nk_clear(struct nk_context*);
	C.nk_clear(ctx.raw())
}

// ForEach calls f for every queued command in ctx until f returns false. This
// function is equivalent to nk_foreach. The return value is true unless f
// returned false.
//
// ForEach implicitly discards GPU acceleration. To use GPU-accelerated
// rendering instead, see Convert and DrawForEach.
func (ctx *Context) ForEach(f func(Command) bool) bool {
	// const struct nk_command* nk__begin(struct nk_context*);
	// const struct nk_command* nk__next(struct nk_context*, const struct nk_command*);
	for cmd := C.nk__begin(ctx.raw()); cmd != nil; cmd = C.nk__next(ctx.raw(), cmd) {
		if !f(typeSwitchCommand(cmd)) {
			return false
		}
	}
	return true
}

// DrawForEach calls f for every queued draw command in ctx until f returns
// false. There will not be any queued draw commands unless Convert has been
// called first. The buf supplied should be the same as the commands buffer
// passed to Convert. The return value is true unless f returned false.
//
// DrawForEach implicity assumes GPU accleration. To use non-GPU-accelerated
// rendering instead, see ForEach.
func (ctx *Context) DrawForEach(buf *Buffer, f func(cmd *DrawCommand) bool) bool {
	// const struct nk_draw_command* nk__draw_begin(const struct nk_context*, const struct nk_buffer*);
	// const struct nk_draw_command* nk__draw_next(const struct nk_draw_command*, const struct nk_buffer*, const struct nk_context*);
	for cmd := C.nk__draw_begin(ctx.raw(), buf.raw()); cmd != nil; cmd = C.nk__draw_next(cmd, buf.raw(), ctx.raw()) {
		goCmd := (*DrawCommand)(unsafe.Pointer(cmd))
		if !f(goCmd) {
			return false
		}
	}
	return true
}

func (ctx *Context) raw() *C.struct_nk_context {
	return (*C.struct_nk_context)(ctx)
}

// typedef union {void *ptr; int id;} nk_handle;

// Handle is used for opaque references to externally created and managed
// resources, outside of Nuklear's control, such as textures and user-data.
//
// Proper use of Handle when wrapping C memory pointers (e.g. graphics
// driver-specific textures):
//
//	// to create the handle and pass it to C
//	handleIn := nk.Handle(unsafe.Pointer(textureIn))
//	// to recover the pointer after getting it back from C
//	textureOut := (*TextureType)(unsafe.Pointer(handleOut))
//
// Proper use of Handle when wrapping Go memory types or pointers (e.g.
// some user-data struct defined entirely in Go):
//
//	import "runtime/cgo"
//	// to create the handle and pass it to C
//	handleIn := nk.Handle(cgo.NewHandle(userDataIn))
//	// to recover the pointer or value after getting it back from C
//	userDataOut := cgo.Handle(handleOut).Value().(UserDataType)
//
// See the documentation for cgo.Handle as well, as each handle consumes
// resources and should be deleted when no longer needed.
type Handle uintptr

func (h Handle) raw() C.nk_handle {
	return *(*C.nk_handle)(unsafe.Pointer(&h))
}

type ColorFormat uint32

const (
	RGB  ColorFormat = C.NK_RGB
	RGBA ColorFormat = C.NK_RGBA
)

// struct nk_color {nk_byte r,g,b,a;};

type Color struct{ R, G, B, A uint8 }

// struct nk_colorf {float r,g,b,a;};

type Colorf struct{ R, G, B, A float32 }

func rawColorf(in C.struct_nk_colorf) Colorf {
	return *(*Colorf)(unsafe.Pointer(&in))
}

func (c Colorf) raw() C.struct_nk_colorf {
	return *(*C.struct_nk_colorf)(unsafe.Pointer(&c))
}

// struct nk_vec2 {float x,y;};

type Vec2 struct{ X, Y float32 }

// struct nk_vec2i {short x, y;};

type Vec2i struct{ X, Y int16 }

// struct nk_rect {float x,y,w,h;};

type Rect struct{ X, Y, W, H float32 }

// struct nk_recti {short x,y,w,h;};

type Recti struct{ X, Y, W, H int16 }

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
