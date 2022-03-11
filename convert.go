package nk

// #include "nk.h"
import "C"
import "unsafe"

// enum nk_anti_aliasing {NK_ANTI_ALIASING_OFF, NK_ANTI_ALIASING_ON};

type AntiAliasing uint32

const (
	AntiAliasingOff AntiAliasing = C.NK_ANTI_ALIASING_OFF
	AnriAliasingOn  AntiAliasing = C.NK_ANTI_ALIASING_ON
)

// enum nk_convert_result {
//     NK_CONVERT_SUCCESS = 0,
//     NK_CONVERT_INVALID_PARAM = 1,
//     NK_CONVERT_COMMAND_BUFFER_FULL = NK_FLAG(1),
//     NK_CONVERT_VERTEX_BUFFER_FULL = NK_FLAG(2),
//     NK_CONVERT_ELEMENT_BUFFER_FULL = NK_FLAG(3)
// };

type ConvertError uint32

const (
	ErrConvertInvalidParam      ConvertError = C.NK_CONVERT_INVALID_PARAM
	ErrConvertCommandBufferFull ConvertError = C.NK_CONVERT_COMMAND_BUFFER_FULL
	ErrConvertVertexBufferFull  ConvertError = C.NK_CONVERT_VERTEX_BUFFER_FULL
	ErrConvertElementBufferFull ConvertError = C.NK_CONVERT_ELEMENT_BUFFER_FULL
)

func (e ConvertError) Error() string {
	switch e {
	case ErrConvertInvalidParam:
		return "invalid parameter"
	case ErrConvertCommandBufferFull:
		return "command buffer full"
	case ErrConvertVertexBufferFull:
		return "vertex buffer full"
	case ErrConvertElementBufferFull:
		return "element buffer full"
	default:
		return "unknown error"
	}
}

// struct nk_convert_config {
//     float global_alpha; /* global alpha value */
//     enum nk_anti_aliasing line_AA; /* line anti-aliasing flag can be turned off if you are tight on memory */
//     enum nk_anti_aliasing shape_AA; /* shape anti-aliasing flag can be turned off if you are tight on memory */
//     unsigned circle_segment_count; /* number of segments used for circles: default to 22 */
//     unsigned arc_segment_count; /* number of segments used for arcs: default to 22 */
//     unsigned curve_segment_count; /* number of segments used for curves: default to 22 */
//     struct nk_draw_null_texture null; /* handle to texture with a white pixel for shape drawing */
//     const struct nk_draw_vertex_layout_element *vertex_layout; /* describes the vertex output format and packing */
//     nk_size vertex_size; /* sizeof one vertex for vertex packing */
//     nk_size vertex_alignment; /* vertex alignment: Can be obtained by NK_ALIGNOF */
// };

type ConvertConfig struct {
	GlobalAlpha        float32
	LineAA, ShapeAA    AntiAliasing
	CircleSegmentCount uint32
	ArcSegmentCount    uint32
	CurveSegmentCount  uint32
	Null               DrawNullTexture
	VertexLayout       *DrawVertexLayoutElement
	VertexSize         uintptr
	VertexAlignment    uintptr
}

// struct nk_draw_null_texture {
//     nk_handle texture; /* texture handle to a texture with a white pixel */
//     struct nk_vec2 uv; /* coordinates to a white pixel in the texture  */
// };

type DrawNullTexture struct {
	Texture Handle
	UV      Vec2
}

// struct nk_draw_vertex_layout_element {
//     enum nk_draw_vertex_layout_attribute attribute;
//     enum nk_draw_vertex_layout_format format;
//     nk_size offset;
// };

type DrawVertexLayoutElement struct {
	Attribute DrawVertexLayoutAttribute
	Format    DrawVertexLayoutFormat
	Offset    uintptr
}

// enum nk_draw_vertex_layout_attribute {
//     NK_VERTEX_POSITION,
//     NK_VERTEX_COLOR,
//     NK_VERTEX_TEXCOORD,
//     NK_VERTEX_ATTRIBUTE_COUNT
// };

type DrawVertexLayoutAttribute uint32

const (
	VertexPosition       DrawVertexLayoutAttribute = C.NK_VERTEX_POSITION
	VertexColor          DrawVertexLayoutAttribute = C.NK_VERTEX_COLOR
	VertexTexcoord       DrawVertexLayoutAttribute = C.NK_VERTEX_TEXCOORD
	VertexAttributeCount DrawVertexLayoutAttribute = C.NK_VERTEX_ATTRIBUTE_COUNT
)

// enum nk_draw_vertex_layout_format {
//     NK_FORMAT_SCHAR,
//     NK_FORMAT_SSHORT,
//     NK_FORMAT_SINT,
//     NK_FORMAT_UCHAR,
//     NK_FORMAT_USHORT,
//     NK_FORMAT_UINT,
//     NK_FORMAT_FLOAT,
//     NK_FORMAT_DOUBLE,
//
// NK_FORMAT_COLOR_BEGIN,
//     NK_FORMAT_R8G8B8 = NK_FORMAT_COLOR_BEGIN,
//     NK_FORMAT_R16G15B16,
//     NK_FORMAT_R32G32B32,
//
//     NK_FORMAT_R8G8B8A8,
//     NK_FORMAT_B8G8R8A8,
//     NK_FORMAT_R16G15B16A16,
//     NK_FORMAT_R32G32B32A32,
//     NK_FORMAT_R32G32B32A32_FLOAT,
//     NK_FORMAT_R32G32B32A32_DOUBLE,
//
//     NK_FORMAT_RGB32,
//     NK_FORMAT_RGBA32,
// NK_FORMAT_COLOR_END = NK_FORMAT_RGBA32,
//     NK_FORMAT_COUNT
// };

type DrawVertexLayoutFormat uint32

const (
	FormatSchar  DrawVertexLayoutFormat = C.NK_FORMAT_SCHAR
	FormatSshort DrawVertexLayoutFormat = C.NK_FORMAT_SSHORT
	FormatSint   DrawVertexLayoutFormat = C.NK_FORMAT_SINT
	FormatUchar  DrawVertexLayoutFormat = C.NK_FORMAT_UCHAR
	FormatUshort DrawVertexLayoutFormat = C.NK_FORMAT_USHORT
	FormatUint   DrawVertexLayoutFormat = C.NK_FORMAT_UINT
	FormatFloat  DrawVertexLayoutFormat = C.NK_FORMAT_FLOAT
	FormatDouble DrawVertexLayoutFormat = C.NK_FORMAT_DOUBLE

	FormatColorBegin DrawVertexLayoutFormat = C.NK_FORMAT_COLOR_BEGIN

	FormatR8G8B8    DrawVertexLayoutFormat = C.NK_FORMAT_R8G8B8
	FormatR16G15B16 DrawVertexLayoutFormat = C.NK_FORMAT_R16G15B16
	FormatR32G32B32 DrawVertexLayoutFormat = C.NK_FORMAT_R32G32B32

	FormatR8G8B8A8           DrawVertexLayoutFormat = C.NK_FORMAT_R8G8B8A8
	FormatG8B8R8A8           DrawVertexLayoutFormat = C.NK_FORMAT_B8G8R8A8
	FormatR16G15B16A16       DrawVertexLayoutFormat = C.NK_FORMAT_R16G15B16A16
	FormatR32G32B32A32       DrawVertexLayoutFormat = C.NK_FORMAT_R32G32B32A32
	FormatR32G32B32A32Float  DrawVertexLayoutFormat = C.NK_FORMAT_R32G32B32A32_FLOAT
	FormatR32G32B32A32Double DrawVertexLayoutFormat = C.NK_FORMAT_R32G32B32A32_DOUBLE

	FormatRGB32  DrawVertexLayoutFormat = C.NK_FORMAT_RGB32
	FormatRGBA32 DrawVertexLayoutFormat = C.NK_FORMAT_RGBA32

	FormatColorEnd DrawVertexLayoutFormat = C.NK_FORMAT_COLOR_END

	FormatCount DrawVertexLayoutFormat = C.NK_FORMAT_COUNT
)

func (ctx *Context) Convert(cmds, vertices, elements *Buffer, config *ConvertConfig) error {
	// nk_flags nk_convert(struct nk_context*, struct nk_buffer *cmds, struct nk_buffer *vertices,
	//     struct nk_buffer *elements, const struct nk_convert_config*);
	result := C.nk_convert(
		&ctx.raw,
		&cmds.raw,
		&vertices.raw,
		&elements.raw,
		(*C.struct_nk_convert_config)(unsafe.Pointer(config)),
	)
	if result != C.NK_CONVERT_SUCCESS {
		return ConvertError(result)
	}
	return nil
}
