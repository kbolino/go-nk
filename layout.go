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

// LayoutRowDynamic calls nk_layout_row_dynamic which starts a layout row of
// dynamic width (i.e., it expands and shrinks with the window).
func (ctx *Context) LayoutRowDynamic(height float32, cols int32) {
	// void nk_layout_row_dynamic(struct nk_context *ctx, float height, int cols);
	C.nk_layout_row_dynamic(ctx.raw(), C.float(height), C.int(cols))
}

// LayoutRowStatic calls nk_layout_row_static which starts a layout row of
// static width (i.e., it doesn't resize with the window).
func (ctx *Context) LayoutRowStatic(height float32, itemWidth, cols int32) {
	// void nk_layout_row_static(struct nk_context *ctx, float height, int item_width, int cols);
	C.nk_layout_row_static(ctx.raw(), C.float(height), C.int(itemWidth), C.int(cols))
}

// Text calls nk_text which inserts a text widget at the next column of the
// current row. Alignment is set according to align.
func (ctx *Context) Text(label string, align TextAlign) {
	str := cStringPool.Get(label)
	defer cStringPool.Release(str)
	// void nk_text(struct nk_context*, const char*, int, nk_flags);
	C.nk_text(ctx.raw(), str, C.int(len(label)), C.nk_flags(align))
}

// ButtonText calls nk_button_text which inserts a button widget at the next
// column of the current row. ButtonText returns true if the button was clicked.
func (ctx *Context) ButtonText(title string) bool {
	str := cStringPool.Get(title)
	defer cStringPool.Release(str)
	// NK_API nk_bool nk_button_text(struct nk_context*, const char *title, int len);
	return bool(C.nk_button_text(ctx.raw(), str, C.int(len(title))))
}

// CheckText calls nk_check_text which inserts a checkbox widget at the next
// column of the current row. CheckText gets its initial state from active and
// returns its new state.
func (ctx *Context) CheckText(text string, active bool) bool {
	str := cStringPool.Get(text)
	defer cStringPool.Release(str)
	// nk_check_text(struct nk_context *ctx, const char *text, int len, nk_bool active)
	return bool(C.nk_check_text(ctx.raw(), str, C.int(len(text)), C.nk_bool(active)))
}
