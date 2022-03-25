package nk

// #include "nk.h"
// #include <stdlib.h>
import "C"

// LayoutSetMinRowHeight sets the currently used minimum row height to the
// specified value.
func (ctx *Context) LayoutSetMinRowHeight(height float32) {
	// void nk_layout_set_min_row_height(struct nk_context*, float height);
	C.nk_layout_set_min_row_height(
		ctx.raw(),
		C.float(height),
	)
}

// LayoutResetMinRowHeight resets the currently used minimum row height to font
// height.
func (ctx *Context) LayoutResetMinRowHeight() {
	// void nk_layout_reset_min_row_height(struct nk_context*);
	C.nk_layout_reset_min_row_height(ctx.raw())
}

// LayoutWidgetBounds calculates current width a static layout row can fit
// inside a window.
func (ctx *Context) LayoutWidgetBounds() Rect {
	// struct nk_rect nk_layout_widget_bounds(struct nk_context*);
	return rawRect(C.nk_layout_widget_bounds(ctx.raw()))
}

// LayoutRatioFromPixel calculates window ratio from pixel size.
func (ctx *Context) LayoutRatioFromPixel(pixelWidth float32) float32 {
	// float nk_layout_ratio_from_pixel(struct nk_context*, float pixel_width);
	return float32(C.nk_layout_ratio_from_pixel(
		ctx.raw(),
		C.float(pixelWidth),
	))
}

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
