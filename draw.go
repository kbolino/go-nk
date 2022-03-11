package nk

// #include "nk.h"
import "C"

// enum nk_anti_aliasing {NK_ANTI_ALIASING_OFF, NK_ANTI_ALIASING_ON};

type AntiAliasing uint32

const (
	AntiAliasingOff AntiAliasing = C.NK_ANTI_ALIASING_OFF
	AnriAliasingOn  AntiAliasing = C.NK_ANTI_ALIASING_ON
)

// enum nk_draw_list_stroke {
//     NK_STROKE_OPEN = nk_false,
//     /* build up path has no connection back to the beginning */
//     NK_STROKE_CLOSED = nk_true
//     /* build up path has a connection back to the beginning */
// };

type DrawListStroke uint32

const (
	StrokeOpen   DrawListStroke = C.NK_STROKE_OPEN
	StrokeClosed DrawListStroke = C.NK_STROKE_CLOSED
)

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

// #define NK_VERTEX_LAYOUT_END NK_VERTEX_ATTRIBUTE_COUNT,NK_FORMAT_COUNT,0

var VertexLayoutEnd = DrawVertexLayoutElement{
	Attribute: VertexAttributeCount,
	Format:    FormatCount,
	Offset:    0,
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

// struct nk_draw_command {
//     unsigned int elem_count;
//     /* number of elements in the current draw batch */
//     struct nk_rect clip_rect;
//     /* current screen clipping rectangle */
//     nk_handle texture;
//     /* current texture to set */
// #ifdef NK_INCLUDE_COMMAND_USERDATA
//     nk_handle userdata;
// #endif
// };

type DrawCommand struct {
	ElemCount uint32
	ClipRect  Rect
	Texture   Handle
}

// struct nk_draw_list {
//     struct nk_rect clip_rect;
//     struct nk_vec2 circle_vtx[12];
//     struct nk_convert_config config;
//
//     struct nk_buffer *buffer;
//     struct nk_buffer *vertices;
//     struct nk_buffer *elements;
//
//     unsigned int element_count;
//     unsigned int vertex_count;
//     unsigned int cmd_count;
//     nk_size cmd_offset;
//
//     unsigned int path_count;
//     unsigned int path_offset;
//
//     enum nk_anti_aliasing line_AA;
//     enum nk_anti_aliasing shape_AA;
//
// #ifdef NK_INCLUDE_COMMAND_USERDATA
//     nk_handle userdata;
// #endif
// };

type DrawList struct {
	ClipRect     Rect
	CircleVtx    [12]Vec2
	Config       ConvertConfig
	Buffer       *Buffer
	Vertices     *Buffer
	Elements     *Buffer
	ElementCount uint32
	VertexCount  uint32
	CommandCount uint32
	CmdOffset    uintptr
	PathCount    uint32
	PathOffset   uint32
	LineAA       AntiAliasing
	ShapeAA      AntiAliasing
}
