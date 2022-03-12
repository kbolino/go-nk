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

// FontAtlasFormat represents the ways a font can be baked.
type FontAtlasFormat uint32

const (
	FontAtlasAlpha8 FontAtlasFormat = C.NK_FONT_ATLAS_ALPHA8
	FontAtlasRGBA32 FontAtlasFormat = C.NK_FONT_ATLAS_RGBA32
)

// Font is an opaque handle to a font in a font atlas.
type Font C.struct_nk_font

func (f *Font) Handle() *UserFont {
	return (*UserFont)(&f.handle)
}

// FontAtlas is an opaque handle to the nk_font_atlas struct, which contains an
// "atlas" of loaded fonts for baking. As with Context, the only safe way to
// create a FontAtlas is with NewFontAtlas.
type FontAtlas C.struct_nk_font_atlas

// NewFontAtlas creates and initializes a new FontAtlas in C memory and returns
// a pointer to it. The memory should be released with Free when no longer
// needed. Under the hood, NewFontAtlas calls nk_font_atlas_init_default.
func NewFontAtlas() *FontAtlas {
	ptr := (*C.struct_nk_font_atlas)(C.malloc(C.sizeof_struct_nk_font_atlas))
	// void nk_font_atlas_init_default(struct nk_font_atlas*);
	C.nk_font_atlas_init_default(ptr)
	return (*FontAtlas)(ptr)
}

// Free releases the memory used by a including a itself. Free is nil-safe.
// After the call, if a was not nil, it is now a dangling pointer.
func (a *FontAtlas) Free() {
	if a != nil {
		// void nk_font_atlas_clear(struct nk_font_atlas*);
		C.nk_font_atlas_clear(a.raw())
		C.free(unsafe.Pointer(a))
	}
}

// Begin calls nk_font_atlas_begin which allows fonts to be added to a.
func (a *FontAtlas) Begin() {
	// void nk_font_atlas_begin(struct nk_font_atlas*);
	C.nk_font_atlas_begin(a.raw())
}

// AddDefaultFont calls nk_font_atlas_add_default which adds the default font
// to a and returns a pointer to it.
func (a *FontAtlas) AddDefaultFont(height float32) *Font {
	// struct nk_font* nk_font_atlas_add_default(struct nk_font_atlas*, float height, const struct nk_font_config*);
	return (*Font)(C.nk_font_atlas_add_default(a.raw(), C.float(height), nil))
}

// AddFromFile calls nk_font_atlas_add_from_file which adds a TTF font from the
// file at the given path. This function uses the C standard library for I/O,
// and moreover has very limited error handling; it either succeeds and returns
// a non-nil error, or fails with a generic error message.
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

// Bake calls nk_font_atlas_bake which "bakes" the loaded fonts into an image.
// The number of bytes in image will be equal to the width times the height for
// Alpha8 images, or 4 times larger than that for RGBA32 images. The memory used
// by the image will remain allocated until Cleanup is called.
func (a *FontAtlas) Bake(format FontAtlasFormat) (image []byte, width, height int32) {
	var rawWidth, rawHeight C.int
	// const void* nk_font_atlas_bake(struct nk_font_atlas*, int *width, int *height, enum nk_font_atlas_format);
	imagePtr := C.nk_font_atlas_bake(a.raw(), &rawWidth, &rawHeight, C.enum_nk_font_atlas_format(format))
	width = int32(rawWidth)
	height = int32(rawHeight)
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
	dataSize := pixelSize * int(width) * int(height)
	image = fakeByteSlice(imagePtr, dataSize)
	return
}

// End calls nk_font_atlas_end which completes the font baking process and
// associates tex with the baked image. After this point, the loaded fonts can
// be supplied to Context.StyleSetFont. End does not free memory which is no
// longer needed, use Cleanup for that.
//
// What tex means is up to the rendering implementation. In general, it should
// be a handle or pointer to a texture that the renderer created and knows how
// to use.
//
// End returns a DrawNullTexture which can be used with Context.Convert.
func (a *FontAtlas) End(tex Handle) DrawNullTexture {
	var null C.struct_nk_draw_null_texture
	// void nk_font_atlas_end(struct nk_font_atlas*, nk_handle tex, struct nk_draw_null_texture*);
	C.nk_font_atlas_end(a.raw(), tex.raw(), &null)
	return *(*DrawNullTexture)(unsafe.Pointer(&null))
}

// Cleanup frees memory which was useful during the font baking process but is
// not needed any longer after End. Note that this includes the baked image,
// which should have been converted to a texture in the rendering system.
func (a *FontAtlas) Cleanup() {
	// void nk_font_atlas_cleanup(struct nk_font_atlas *atlas);
	C.nk_font_atlas_cleanup(a.raw())
}

func (a *FontAtlas) raw() *C.struct_nk_font_atlas {
	return (*C.struct_nk_font_atlas)(a)
}
