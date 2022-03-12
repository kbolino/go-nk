package nk

// #include "nk.h"
// #include <stdlib.h>
import "C"

import "unsafe"

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
// current row. The label is copied into C memory and freed at the end of
// the call. Alignment is set according to align.
func (ctx *Context) Text(label string, align TextAlign) {
	str := C.CString(label)
	defer C.free(unsafe.Pointer(str))
	// void nk_text(struct nk_context*, const char*, int, nk_flags);
	C.nk_text(ctx.raw(), str, C.int(len(label)), C.nk_flags(align))
}

// TextBytes is like Text except that it doesn't have to copy label.
func (ctx *Context) TextBytes(label []byte, align TextAlign) {
	// void nk_text(struct nk_context*, const char*, int, nk_flags);
	C.nk_text(ctx.raw(), (*C.char)(unsafe.Pointer(&label[0])), C.int(len(label)), C.nk_flags(align))
}

// ButtonText calls nk_button_text which inserts a button widget at the next
// column of the current row. The title is copied into C memory that's freed at
// the end of the call. ButtonText returns true if the button was clicked.
func (ctx *Context) ButtonText(title string) bool {
	str := C.CString(title)
	defer C.free(unsafe.Pointer(str))
	// NK_API nk_bool nk_button_text(struct nk_context*, const char *title, int len);
	return bool(C.nk_button_text(ctx.raw(), str, C.int(len(title))))
}

// ButtonText bytes is like ButtonText except that it doesn't have to copy
// title.
func (ctx *Context) ButtonTextBytes(title []byte) bool {
	// NK_API nk_bool nk_button_text(struct nk_context*, const char *title, int len);
	return bool(C.nk_button_text(ctx.raw(), (*C.char)(unsafe.Pointer(&title[0])), C.int(len(title))))
}

// CheckText calls nk_check_text which inserts a checkbox widget at the next
// column of the current row. The text is copied into C memory that's freed at
// the end of the call. CheckText gets its initial state from active and returns
// its current state.
func (ctx *Context) CheckText(text string, active bool) bool {
	str := C.CString(text)
	defer C.free(unsafe.Pointer(str))
	// nk_check_text(struct nk_context *ctx, const char *text, int len, nk_bool active)
	return bool(C.nk_check_text(ctx.raw(), str, C.int(len(text)), C.nk_bool(active)))
}

// CheckTextBytes is like CheckText except that it doesn't have to copy text.
func (ctx *Context) CheckTextBytes(text []byte, active bool) bool {
	return bool(C.nk_check_text(ctx.raw(), (*C.char)(unsafe.Pointer(&text[0])), C.int(len(text)), C.nk_bool(active)))
}
