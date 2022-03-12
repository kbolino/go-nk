package nk

// #include "nk.h"
import "C"

const InputMax = C.NK_INPUT_MAX

// InputBegin calls nk_input_begin which begins the input handling phase of
// each frame.
func (ctx *Context) InputBegin() {
	// nk_input_begin(struct nk_context*);
	C.nk_input_begin(ctx.raw())
}

// InputMotion calls nk_input_motion which records a cursor motion event. Such
// events are produced by e.g. moving the mouse.
func (ctx *Context) InputMotion(x, y int32) {
	// void nk_input_motion(struct nk_context*, int x, int y);
	C.nk_input_motion(ctx.raw(), C.int(x), C.int(y))
}

// InputKey calls nk_input_key which records a key input event. Such events are
// produced by e.g. pressing or releasing a keyboard key.
func (ctx *Context) InputKey(key Key, down bool) {
	// void nk_input_key(struct nk_context*, enum nk_keys, nk_bool down);
	C.nk_input_key(ctx.raw(), C.enum_nk_keys(key), C.nk_bool(down))
}

// InputButton calls nk_input_button which records a button input event.
// Such events are produced by e.g. pressing or releasing a mouse button.
func (ctx *Context) InputButton(button Button, x, y int32, down bool) {
	// void nk_input_button(struct nk_context*, enum nk_buttons, int x, int y, nk_bool down);
	C.nk_input_button(ctx.raw(), C.enum_nk_buttons(button), C.int(x), C.int(y), C.nk_bool(down))
}

// InputScroll calls nk_input_scroll which records a scroll input event.
// Such events are produced by e.g. rolling the scrollwheel of a mouse.
func (ctx *Context) InputScroll(x, y float32) {
	// void nk_input_scroll(struct nk_context*, struct nk_vec2 val);
	C.nk_input_scroll(ctx.raw(), C.nk_vec2(C.float(x), C.float(y)))
}

// InputChar calls nk_input_char which records an ASCII text input event.
func (ctx *Context) InputChar(c byte) {
	// void nk_input_char(struct nk_context*, char);
	C.nk_input_char(ctx.raw(), C.char(c))
}

// InputUnicode calls nk_input_unicode which records a Unicode text input event.
func (ctx *Context) InputUnicode(r rune) {
	// void nk_input_unicode(struct nk_context*, nk_rune rune);
	C.nk_input_unicode(ctx.raw(), C.nk_rune(r))
}

func (ctx *Context) InputEnd() {
	// void nk_input_end(struct nk_context*);
	C.nk_input_end(ctx.raw())
}
