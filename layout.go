package nk

// #include "nk.h"
// #include <stdlib.h>
import "C"

// LayoutRowDynamic calls nk_layout_row_dynamic which starts a layout row of
// dynamic width (i.e., it expands and shrinks with the window).
func (ctx *Context) LayoutRowDynamic(height float32, cols int32) {
	// void nk_layout_row_dynamic(struct nk_context *ctx, float height, int cols);
	C.nk_layout_row_dynamic(
		ctx.raw(),
		C.float(height),
		C.int(cols),
	)
}

// LayoutRowStatic calls nk_layout_row_static which starts a layout row of
// static width (i.e., it doesn't resize with the window).
func (ctx *Context) LayoutRowStatic(height float32, itemWidth, cols int32) {
	// void nk_layout_row_static(struct nk_context *ctx, float height, int item_width, int cols);
	C.nk_layout_row_static(
		ctx.raw(),
		C.float(height),
		C.int(itemWidth),
		C.int(cols),
	)
}
