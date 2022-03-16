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

// Handle returns the UserFont handle for f.
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
// to a and returns a pointer to it. The config parameter may be nil.
func (a *FontAtlas) AddDefaultFont(height float32, config *FontConfig) *Font {
	// struct nk_font* nk_font_atlas_add_default(struct nk_font_atlas*, float height, const struct nk_font_config*);
	return (*Font)(C.nk_font_atlas_add_default(a.raw(), C.float(height), (*C.struct_nk_font_config)(config)))
}

// AddFromFile calls nk_font_atlas_add_from_file which adds a TTF font from the
// file at the given path. This function uses the C standard library for I/O,
// and moreover has very limited error handling; it either succeeds and returns
// a non-nil error, or fails with a generic error message. The config parameter
// may be nil.
func (a *FontAtlas) AddFromFile(filePath string, height float32, config *FontConfig) (*Font, error) {
	rawFilePath := cStringPool.Get(filePath)
	defer cStringPool.Release(rawFilePath)
	// struct nk_font* nk_font_atlas_add_from_file(struct nk_font_atlas *atlas, const char *file_path,
	//     float height, const struct nk_font_config*);
	font := (*Font)(C.nk_font_atlas_add_from_file(a.raw(), rawFilePath, C.float(height),
		(*C.struct_nk_font_config)(config)))
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

// struct nk_font_config {
//     struct nk_font_config *next;
//     /* NOTE: only used internally */
//     void *ttf_blob;
//     /* pointer to loaded TTF file memory block.
//      * NOTE: not needed for nk_font_atlas_add_from_memory and nk_font_atlas_add_from_file. */
//     nk_size ttf_size;
//     /* size of the loaded TTF file memory block
//      * NOTE: not needed for nk_font_atlas_add_from_memory and nk_font_atlas_add_from_file. */
//
//     unsigned char ttf_data_owned_by_atlas;
//     /* used inside font atlas: default to: 0*/
//     unsigned char merge_mode;
//     /* merges this font into the last font */
//     unsigned char pixel_snap;
//     /* align every character to pixel boundary (if true set oversample (1,1)) */
//     unsigned char oversample_v, oversample_h;
//     /* rasterize at high quality for sub-pixel position */
//     unsigned char padding[3];
//
//     float size;
//     /* baked pixel height of the font */
//     enum nk_font_coord_type coord_type;
//     /* texture coordinate format with either pixel or UV coordinates */
//     struct nk_vec2 spacing;
//     /* extra pixel spacing between glyphs  */
//     const nk_rune *range;
//     /* list of unicode ranges (2 values per range, zero terminated) */
//     struct nk_baked_font *font;
//     /* font to setup in the baking process: NOTE: not needed for font atlas */
//     nk_rune fallback_glyph;
//     /* fallback glyph to use if a given rune is not found */
//     struct nk_font_config *n;
//     struct nk_font_config *p;
// };

// FontConfig is an opaque handle to the configuration for loading a font.
// To create a FontConfig, use FontConfigBuilder. FontConfig is stored in C
// memory which must be released by Free.
type FontConfig C.struct_nk_font_config

// Free releases the memory used by fc, including fc itself. Free is nil-safe.
// After the call, if fc was not nil, it is now a dangling pointer.
func (fc *FontConfig) Free() {
	if fc != nil {
		C.free(unsafe.Pointer(fc))
	}
}

// FontConfigBuilder is used to construct FontConfig values.
type FontConfigBuilder struct {
	// MergeMode specifies whether this font should be merged into the previous
	// font when baking.
	MergeMode bool
	// PixelSnap specifies whether every character of this font should be
	// aligned at a pixel boundary. If set to true, then oversampling should
	// be set to (1, 1).
	PixelSnap bool
	// OversampleH and OversampleV are the degrees to which the font should be
	// oversampled (rendered above its nominal resolution) in the horizontal
	// and vertical directions, respectively. Oversampling is useful to obtain
	// subpixel resolution and for high-DPI rendering.
	OversampleH, OversampleV uint8
	// CoordType is the kind of coordinates to use with the baked texture.
	CoordType FontCoordType
	// FallbackGlyph is the alternate code point to use when no glyph exists
	// for another code point.
	FallbackGlyph rune
}

// Build creates a new FontConfig from the state of fcb. The resulting value is
// stored in C memory and must be freed after use.
func (fcb FontConfigBuilder) Build() *FontConfig {
	raw := (*C.struct_nk_font_config)(C.malloc(C.sizeof_struct_nk_font_config))
	mergeMode := C.nk_false
	if fcb.MergeMode {
		mergeMode = C.nk_true
	}
	pixelSnap := C.nk_false
	if fcb.PixelSnap {
		pixelSnap = C.nk_true
	}
	*raw = C.struct_nk_font_config{
		merge_mode:     C.uchar(mergeMode),
		pixel_snap:     C.uchar(pixelSnap),
		oversample_h:   C.uchar(fcb.OversampleH),
		oversample_v:   C.uchar(fcb.OversampleV),
		coord_type:     C.enum_nk_font_coord_type(fcb.CoordType),
		fallback_glyph: C.nk_rune(fcb.FallbackGlyph),
	}
	return (*FontConfig)(raw)
}

// FontCoordType specifies the type of coordinates to use for baked fonts.
type FontCoordType uint32

const (
	CoordUV    FontCoordType = C.NK_COORD_UV    // UV coordinates, i.e. scaled from 0 to 1
	CoordPixel FontCoordType = C.NK_COORD_PIXEL // absolute pixel coordinates
)
