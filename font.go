package nk

// #include "nk.h"
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

// enum nk_font_atlas_format {
//     NK_FONT_ATLAS_ALPHA8,
//     NK_FONT_ATLAS_RGBA32
// };

type FontAtlasFormat uint32

const (
	FontAtlasAlpha8 FontAtlasFormat = C.NK_FONT_ATLAS_ALPHA8
	FontAtlasRGBA32 FontAtlasFormat = C.NK_FONT_ATLAS_RGBA32
)

type Font C.struct_nk_font

type FontAtlas C.struct_nk_font_atlas

func NewFontAtlas() *FontAtlas {
	ptr := (*C.struct_nk_font_atlas)(C.malloc(C.sizeof_struct_nk_font_atlas))
	// void nk_font_atlas_init_default(struct nk_font_atlas*);
	C.nk_font_atlas_init_default(ptr)
	return (*FontAtlas)(ptr)
}

func (a *FontAtlas) Free() {
	if a != nil {
		// void nk_font_atlas_clear(struct nk_font_atlas*);
		C.nk_font_atlas_clear(a.raw())
		C.free(unsafe.Pointer(a))
	}
}

func (a *FontAtlas) Begin() {
	// void nk_font_atlas_begin(struct nk_font_atlas*);
	C.nk_font_atlas_begin(a.raw())
}

func (a *FontAtlas) AddDefaultFont(height float32) *Font {
	// struct nk_font* nk_font_atlas_add_default(struct nk_font_atlas*, float height, const struct nk_font_config*);
	return (*Font)(C.nk_font_atlas_add_default(a.raw(), C.float(height), nil))
}

func (a *FontAtlas) AddFromFile(filePath string, height float32) (*Font, error) {
	rawFilePath := C.CString(filePath)
	defer C.free(unsafe.Pointer(rawFilePath))
	// struct nk_font* nk_font_atlas_add_from_file(struct nk_font_atlas *atlas, const char *file_path,
	//     float height, const struct nk_font_config*);
	font := (*Font)(C.nk_font_atlas_add_from_file(a.raw(), rawFilePath, C.float(height), nil))
	if font == nil {
		return nil, fmt.Errorf("error loading font")
	}
	return font, nil
}

func (a *FontAtlas) Bake(widthIn, heightIn int32, format FontAtlasFormat) (image []byte, widthOut, heightOut int32) {
	var width = C.int(widthIn)
	var height = C.int(heightIn)
	// const void* nk_font_atlas_bake(struct nk_font_atlas*, int *width, int *height, enum nk_font_atlas_format);
	imagePtr := C.nk_font_atlas_bake(a.raw(), &width, &height, C.enum_nk_font_atlas_format(format))
	widthOut = int32(width)
	heightOut = int32(height)
	if imagePtr == nil {
		return
	}
	pixelSize := 0
	switch format {
	case FontAtlasAlpha8:
		pixelSize = 1
	case FontAtlasRGBA32:
		pixelSize = 4
	default:
		panic(fmt.Errorf("unsupported font atlas format %d", format))
	}
	dataSize := pixelSize * int(widthOut) * int(heightOut)
	image = fakeByteSlice((*byte)(imagePtr), dataSize)
	return
}

func (a *FontAtlas) End(tex Handle) DrawNullTexture {
	var null C.struct_nk_draw_null_texture
	// void nk_font_atlas_end(struct nk_font_atlas*, nk_handle tex, struct nk_draw_null_texture*);
	C.nk_font_atlas_end(a.raw(), tex.raw(), &null)
	return *(*DrawNullTexture)(unsafe.Pointer(&null))
}

func (a *FontAtlas) Cleanup() {
	// void nk_font_atlas_cleanup(struct nk_font_atlas *atlas);
	C.nk_font_atlas_cleanup(a.raw())
}

func (a *FontAtlas) raw() *C.struct_nk_font_atlas {
	return (*C.struct_nk_font_atlas)(a)
}
