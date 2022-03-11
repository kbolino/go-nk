package nk

// #include "nk.h"
// #include <stdlib.h>
import "C"

import "unsafe"

// enum nk_panel_flags {
//     NK_WINDOW_BORDER            = NK_FLAG(0),
//     NK_WINDOW_MOVABLE           = NK_FLAG(1),
//     NK_WINDOW_SCALABLE          = NK_FLAG(2),
//     NK_WINDOW_CLOSABLE          = NK_FLAG(3),
//     NK_WINDOW_MINIMIZABLE       = NK_FLAG(4),
//     NK_WINDOW_NO_SCROLLBAR      = NK_FLAG(5),
//     NK_WINDOW_TITLE             = NK_FLAG(6),
//     NK_WINDOW_SCROLL_AUTO_HIDE  = NK_FLAG(7),
//     NK_WINDOW_BACKGROUND        = NK_FLAG(8),
//     NK_WINDOW_SCALE_LEFT        = NK_FLAG(9),
//     NK_WINDOW_NO_INPUT          = NK_FLAG(10)
// };

type PanelFlags uint32

const (
	WindowBorder         PanelFlags = C.NK_WINDOW_BORDER
	WindowMovable        PanelFlags = C.NK_WINDOW_MOVABLE
	WindowScalable       PanelFlags = C.NK_WINDOW_SCALABLE
	WindowClosable       PanelFlags = C.NK_WINDOW_CLOSABLE
	WindowMinimizable    PanelFlags = C.NK_WINDOW_MINIMIZABLE
	WindowNoScrollbar    PanelFlags = C.NK_WINDOW_NO_SCROLLBAR
	WindowTitle          PanelFlags = C.NK_WINDOW_TITLE
	WindowScrollAutoHide PanelFlags = C.NK_WINDOW_SCROLL_AUTO_HIDE
	WindowBackground     PanelFlags = C.NK_WINDOW_BACKGROUND
	WindowScaleLeft      PanelFlags = C.NK_WINDOW_SCALE_LEFT
	WindowNoInput        PanelFlags = C.NK_WINDOW_NO_INPUT
)

// enum nk_window_flags {
//     NK_WINDOW_PRIVATE       = NK_FLAG(11),
//     NK_WINDOW_DYNAMIC       = NK_WINDOW_PRIVATE,
//     /* special window type growing up in height while being filled to a certain maximum height */
//     NK_WINDOW_ROM           = NK_FLAG(12),
//     /* sets window widgets into a read only mode and does not allow input changes */
//     NK_WINDOW_NOT_INTERACTIVE = NK_WINDOW_ROM|NK_WINDOW_NO_INPUT,
//     /* prevents all interaction caused by input to either window or widgets inside */
//     NK_WINDOW_HIDDEN        = NK_FLAG(13),
//     /* Hides window and stops any window interaction and drawing */
//     NK_WINDOW_CLOSED        = NK_FLAG(14),
//     /* Directly closes and frees the window at the end of the frame */
//     NK_WINDOW_MINIMIZED     = NK_FLAG(15),
//     /* marks the window as minimized */
//     NK_WINDOW_REMOVE_ROM    = NK_FLAG(16)
//     /* Removes read only mode at the end of the window */
// };

type WindowFlags PanelFlags

const (
	WindowPrivate        WindowFlags = C.NK_WINDOW_PRIVATE
	WindowDynamic        WindowFlags = C.NK_WINDOW_DYNAMIC
	WindowROM            WindowFlags = C.NK_WINDOW_ROM
	WindowNotInteractive WindowFlags = C.NK_WINDOW_NOT_INTERACTIVE
	WindowHidden         WindowFlags = C.NK_WINDOW_HIDDEN
	WindowClosed         WindowFlags = C.NK_WINDOW_CLOSED
	WindowMinimized      WindowFlags = C.NK_WINDOW_MINIMIZED
	WindowRemoveROM      WindowFlags = C.NK_WINDOW_REMOVE_ROM
)

func (ctx *Context) Begin(title string, bounds *Rect, flags PanelFlags) bool {
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

func (ctx *Context) BeginTitled(name, title string, bounds *Rect, flags PanelFlags) bool {
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
