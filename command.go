package nk

// #include "nk.h"
import "C"

import (
	"fmt"
	"unsafe"
)

type CommandType uint32

const (
	CommandTypeNop            CommandType = C.NK_COMMAND_NOP
	CommandTypeScissor        CommandType = C.NK_COMMAND_SCISSOR
	CommandTypeLine           CommandType = C.NK_COMMAND_LINE
	CommandTypeCurve          CommandType = C.NK_COMMAND_CURVE
	CommandTypeRect           CommandType = C.NK_COMMAND_RECT
	CommandTypeRectFilled     CommandType = C.NK_COMMAND_RECT_FILLED
	CommandTypeRectMultiColor CommandType = C.NK_COMMAND_RECT_MULTI_COLOR
	CommandTypeCircle         CommandType = C.NK_COMMAND_CIRCLE
	CommandTypeCircleFilled   CommandType = C.NK_COMMAND_CIRCLE_FILLED
	CommandTypeArc            CommandType = C.NK_COMMAND_ARC
	CommandTypeArcFilled      CommandType = C.NK_COMMAND_ARC_FILLED
	CommandTypeTriangle       CommandType = C.NK_COMMAND_TRIANGLE
	CommandTypeTriangleFilled CommandType = C.NK_COMMAND_TRIANGLE_FILLED
	CommandTypePolygon        CommandType = C.NK_COMMAND_POLYGON
	CommandTypePolygonFilled  CommandType = C.NK_COMMAND_POLYGON_FILLED
	CommandTypePolyline       CommandType = C.NK_COMMAND_POLYLINE
	CommandTypeText           CommandType = C.NK_COMMAND_TEXT
	CommandTypeImage          CommandType = C.NK_COMMAND_IMAGE
	CommandTypeCustom         CommandType = C.NK_COMMAND_CUSTOM
)

type Command interface {
	CommandType() CommandType
}

// struct nk_command {
//     enum nk_command_type type;
//     nk_size next;
// #ifdef NK_INCLUDE_COMMAND_USERDATA
//     nk_handle userdata;
// #endif
// };

type CommandHeader struct {
	Type CommandType
	Next uintptr
}

func (h *CommandHeader) CommandType() CommandType {
	return h.Type
}

type CommandNop struct {
	CommandHeader
}

// struct nk_command_scissor {
//     struct nk_command header;
//     short x, y;
//     unsigned short w, h;
// };

type CommandScissor struct {
	CommandHeader
	X, Y int16
	W, H uint16
}

// struct nk_command_line {
//     struct nk_command header;
//     unsigned short line_thickness;
//     struct nk_vec2i begin;
//     struct nk_vec2i end;
//     struct nk_color color;
// };

type CommandLine struct {
	CommandHeader
	LineThickness uint16
	Begin, End    Vec2i
	Color         Color
}

// struct nk_command_curve {
//     struct nk_command header;
//     unsigned short line_thickness;
//     struct nk_vec2i begin;
//     struct nk_vec2i end;
//     struct nk_vec2i ctrl[2];
//     struct nk_color color;
// };

type CommandCurve struct {
	CommandHeader
	LineThickness uint16
	Begin, End    Vec2i
	Ctrl          [2]Vec2i
	Color         Color
}

// struct nk_command_rect {
//     struct nk_command header;
//     unsigned short rounding;
//     unsigned short line_thickness;
//     short x, y;
//     unsigned short w, h;
//     struct nk_color color;
// };

type CommandRect struct {
	CommandHeader
	Rounding      uint16
	LineThickness uint16
	X, Y          int16
	W, H          uint16
	Color         Color
}

// struct nk_command_rect_filled {
//     struct nk_command header;
//     unsigned short rounding;
//     short x, y;
//     unsigned short w, h;
//     struct nk_color color;
// };

type CommandRectFilled struct {
	CommandHeader
	Rounding uint16
	X, Y     int16
	W, H     uint16
	Color    Color
}

// struct nk_command_rect_multi_color {
//     struct nk_command header;
//     short x, y;
//     unsigned short w, h;
//     struct nk_color left;
//     struct nk_color top;
//     struct nk_color bottom;
//     struct nk_color right;
// };

type CommandRectMultiColor struct {
	CommandHeader
	X, Y   int16
	W, H   uint16
	Left   Color
	Top    Color
	Bottom Color
	Right  Color
}

// struct nk_command_triangle {
//     struct nk_command header;
//     unsigned short line_thickness;
//     struct nk_vec2i a;
//     struct nk_vec2i b;
//     struct nk_vec2i c;
//     struct nk_color color;
// };

type CommandTriangle struct {
	CommandHeader
	LineThickness uint16
	A, B, C       Vec2i
	Color         Color
}

// struct nk_command_triangle_filled {
//     struct nk_command header;
//     struct nk_vec2i a;
//     struct nk_vec2i b;
//     struct nk_vec2i c;
//     struct nk_color color;
// };

type CommandTriangleFilled struct {
	CommandHeader
	A, B, C Vec2i
	Color   Color
}

// struct nk_command_circle {
//     struct nk_command header;
//     short x, y;
//     unsigned short line_thickness;
//     unsigned short w, h;
//     struct nk_color color;
// };

type CommandCircle struct {
	CommandHeader
	X, Y          int16
	LineThickness uint16
	W, H          uint16
	Color         Color
}

// struct nk_command_circle_filled {
//     struct nk_command header;
//     short x, y;
//     unsigned short w, h;
//     struct nk_color color;
// };

type CommandCircleFilled struct {
	CommandHeader
	X, Y  int16
	W, H  uint16
	Color Color
}

// struct nk_command_arc {
//     struct nk_command header;
//     short cx, cy;
//     unsigned short r;
//     unsigned short line_thickness;
//     float a[2];
//     struct nk_color color;
// };

type CommandArc struct {
	CommandHeader
	CX, CY        int16
	R             uint16
	LineThickness uint16
	A             [2]float32
	Color         Color
}

// struct nk_command_arc_filled {
//     struct nk_command header;
//     short cx, cy;
//     unsigned short r;
//     float a[2];
//     struct nk_color color;
// };

type CommandArcFilled struct {
	CommandHeader
	CX, CY int16
	R      uint16
	A      [2]float32
	Color  Color
}

// struct nk_command_polygon {
//     struct nk_command header;
//     struct nk_color color;
//     unsigned short line_thickness;
//     unsigned short point_count;
//     struct nk_vec2i points[1];
// };

type CommandPolygon struct {
	CommandHeader
	Color         Color
	LineThickness uint16
	PointCount    uint16
	FirstPoint    Vec2i
}

func (p *CommandPolygon) Points() []Vec2i {
	return fakeSlice(&p.FirstPoint, int(p.PointCount))
}

// struct nk_command_polygon_filled {
//     struct nk_command header;
//     struct nk_color color;
//     unsigned short point_count;
//     struct nk_vec2i points[1];
// };

type CommandPolygonFilled struct {
	CommandHeader
	Color      Color
	PointCount uint16
	FirstPoint Vec2i
}

func (p *CommandPolygonFilled) Points() []Vec2i {
	return fakeSlice(&p.FirstPoint, int(p.PointCount))
}

// struct nk_command_polyline {
//     struct nk_command header;
//     struct nk_color color;
//     unsigned short line_thickness;
//     unsigned short point_count;
//     struct nk_vec2i points[1];
// };

type CommandPolyline struct {
	CommandHeader
	Color         Color
	LineThickness uint16
	PointCoint    uint16
	FirstPoint    Vec2i
}

func (p *CommandPolyline) Points() []Vec2i {
	return fakeSlice(&p.FirstPoint, int(p.PointCoint))
}

// struct nk_command_image {
//     struct nk_command header;
//     short x, y;
//     unsigned short w, h;
//     struct nk_image img;
//     struct nk_color col;
// };

type CommandImage struct {
	CommandHeader
	X, Y int16
	W, H uint16
	Img  Image
	Col  Color
}

// struct nk_command_text {
//     struct nk_command header;
//     const struct nk_user_font *font;
//     struct nk_color background;
//     struct nk_color foreground;
//     short x, y;
//     unsigned short w, h;
//     float height;
//     int length;
//     char string[1];
// };

type CommandText struct {
	CommandHeader
	Font       *UserFont
	Background Color
	Foreground Color
	X, Y       int16
	W, H       uint16
	Height     float32
	Length     int32
	FirstByte  byte
}

func (t *CommandText) Bytes() []byte {
	return fakeSlice(&t.FirstByte, int(t.Length))
}

func typeSwitchCommand(cmd *C.struct_nk_command) Command {
	rawPtr := unsafe.Pointer(cmd)
	switch CommandType(cmd._type) {
	case CommandTypeNop:
		return (*CommandNop)(rawPtr)
	case CommandTypeScissor:
		return (*CommandScissor)(rawPtr)
	case CommandTypeLine:
		return (*CommandLine)(rawPtr)
	case CommandTypeCurve:
		return (*CommandCurve)(rawPtr)
	case CommandTypeRect:
		return (*CommandRect)(rawPtr)
	case CommandTypeRectFilled:
		return (*CommandRectFilled)(rawPtr)
	case CommandTypeRectMultiColor:
		return (*CommandRectMultiColor)(rawPtr)
	case CommandTypeTriangle:
		return (*CommandTriangle)(rawPtr)
	case CommandTypeTriangleFilled:
		return (*CommandTriangleFilled)(rawPtr)
	case CommandTypeCircle:
		return (*CommandCircle)(rawPtr)
	case CommandTypeCircleFilled:
		return (*CommandCircleFilled)(rawPtr)
	case CommandTypeArc:
		return (*CommandArc)(rawPtr)
	case CommandTypeArcFilled:
		return (*CommandArcFilled)(rawPtr)
	case CommandTypePolygon:
		return (*CommandPolygon)(rawPtr)
	case CommandTypePolygonFilled:
		return (*CommandPolygonFilled)(rawPtr)
	case CommandTypePolyline:
		return (*CommandPolyline)(rawPtr)
	case CommandTypeImage:
		return (*CommandImage)(rawPtr)
	case CommandTypeCustom:
		panic("custom command type unsupported")
	case CommandTypeText:
		return (*CommandText)(rawPtr)
	default:
		panic(fmt.Errorf("unsupported command type %d", cmd._type))
	}
}
