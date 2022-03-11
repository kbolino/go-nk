package nk

// #include "nk.h"
import "C"

func (ctx *Context) StyleSetFont(font *UserFont) {
	// void nk_style_set_font(struct nk_context*, const struct nk_user_font*);
	C.nk_style_set_font(ctx.raw(), (*C.struct_nk_user_font)(font))
}
