package nk

// #include "nk.h"
// #include <stdlib.h>
import "C"

// enum nk_text_align {
//     NK_TEXT_ALIGN_LEFT        = 0x01,
//     NK_TEXT_ALIGN_CENTERED    = 0x02,
//     NK_TEXT_ALIGN_RIGHT       = 0x04,
//     NK_TEXT_ALIGN_TOP         = 0x08,
//     NK_TEXT_ALIGN_MIDDLE      = 0x10,
//     NK_TEXT_ALIGN_BOTTOM      = 0x20
// };
// enum nk_text_alignment {
//     NK_TEXT_LEFT        = NK_TEXT_ALIGN_MIDDLE|NK_TEXT_ALIGN_LEFT,
//     NK_TEXT_CENTERED    = NK_TEXT_ALIGN_MIDDLE|NK_TEXT_ALIGN_CENTERED,
//     NK_TEXT_RIGHT       = NK_TEXT_ALIGN_MIDDLE|NK_TEXT_ALIGN_RIGHT
// };

// TextAlign specifies the horizontal and vertical alignment of text inside a
// widget.
type TextAlign uint32

const (
	TextAlignLeft     TextAlign = C.NK_TEXT_ALIGN_LEFT
	TextAlignCentered TextAlign = C.NK_TEXT_ALIGN_CENTERED
	TextAlignRight    TextAlign = C.NK_TEXT_ALIGN_RIGHT
	TextAlignTop      TextAlign = C.NK_TEXT_ALIGN_TOP
	TextAlignMiddle   TextAlign = C.NK_TEXT_ALIGN_MIDDLE
	TextAlignBottom   TextAlign = C.NK_TEXT_ALIGN_BOTTOM

	TextLeft     = TextAlignMiddle | TextAlignLeft
	TextCentered = TextAlignMiddle | TextAlignCentered
	TextRight    = TextAlignMiddle | TextAlignRight
)

// Text calls nk_text which inserts a text widget. Alignment is set according to
// align.
func (ctx *Context) Text(label string, align TextAlign) {
	str := cStringPool.Get(label)
	defer cStringPool.Release(str)
	// void nk_text(struct nk_context*, const char*, int, nk_flags);
	C.nk_text(
		ctx.raw(),
		str,
		C.int(len(label)),
		C.nk_flags(align),
	)
}

// ButtonText calls nk_button_text which inserts a button widget. ButtonText
// returns true if the button was clicked.
func (ctx *Context) ButtonText(title string) bool {
	str := cStringPool.Get(title)
	defer cStringPool.Release(str)
	// NK_API nk_bool nk_button_text(struct nk_context*, const char *title, int len);
	return bool(C.nk_button_text(
		ctx.raw(),
		str,
		C.int(len(title)),
	))
}

// CheckText calls nk_check_text which inserts a checkbox widget. CheckText gets
// its current state from active and returns its new state.
func (ctx *Context) CheckText(text string, active bool) bool {
	str := cStringPool.Get(text)
	defer cStringPool.Release(str)
	// nk_check_text(struct nk_context *ctx, const char *text, int len, nk_bool active)
	return bool(C.nk_check_text(
		ctx.raw(),
		str,
		C.int(len(text)),
		C.nk_bool(active),
	))
}

// OptionText calls nk_option_text which inserts a radiobox widget. OptionText
// gets its current state from active and returns its new state.
func (ctx *Context) OptionText(text string, active bool) bool {
	str := cStringPool.Get(text)
	defer cStringPool.Release(str)
	// nk_bool nk_option_text(struct nk_context*, const char*, int, nk_bool active);
	return bool(C.nk_option_text(
		ctx.raw(),
		str,
		C.int(len(text)),
		C.nk_bool(active),
	))
}

// SlideFloat calls nk_slide_float which inserts a slider widget. The slider
// takes its current value from val and returns its new value. The parameters
// min, max, and step define the boundaries and snap intervals of the slider.
func (ctx *Context) SlideFloat(min, val, max, step float32) float32 {
	// float nk_slide_float(struct nk_context*, float min, float val, float max, float step);
	return float32(C.nk_slide_float(
		ctx.raw(),
		C.float(min),
		C.float(val),
		C.float(max),
		C.float(step),
	))
}

// SlideInt is like SlideFloat except for nk_slide_int with integer parameters.
func (ctx *Context) SlideInt(min, val, max, step int32) int32 {
	// int nk_slide_int(struct nk_context*, int min, int val, int max, int step);
	return int32(C.nk_slide_int(
		ctx.raw(),
		C.int(min),
		C.int(val),
		C.int(max),
		C.int(step),
	))
}

// Prog calls nk_prog which inserts a progress bar widget. The progress bar
// advances from 0 to max. Prog takes its current state from cur, and returns
// its new state, if modifyable.
func (ctx *Context) Prog(cur, max uint, modifyable bool) uint {
	// nk_size nk_prog(struct nk_context*, nk_size cur, nk_size max, nk_bool modifyable)
	return uint(C.nk_prog(
		ctx.raw(),
		C.nk_size(cur),
		C.nk_size(max),
		C.nk_bool(modifyable),
	))
}

// ColorPicker calls nk_color_picker which inserts a color picker widget. The
// color picker takes its current state from color and returns its new state.
// The format controls whether color alpha is enabled.
func (ctx *Context) ColorPicker(color Colorf, format ColorFormat) Colorf {
	// struct nk_colorf nk_color_picker(struct nk_context*, struct nk_colorf, enum nk_color_format);
	return rawColorf(C.nk_color_picker(
		ctx.raw(),
		color.raw(),
		C.enum_nk_color_format(format),
	))
}
