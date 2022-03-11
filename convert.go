package nk

// #include "nk.h"
// #include <stdlib.h>
import "C"
import "unsafe"

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

// CConvertConfig is the direct analogue of nk_convert_config. However, because
// nk_convert_config expects a C-style array of vertex layout elements, that
// field is present but not exported in CConvertConfig to avoid allocation
// issues. Instead, to create a "real" ConvertConfig, use ConvertConfigBuilder.
type CConvertConfig struct {
	GlobalAlpha        float32
	LineAA, ShapeAA    AntiAliasing
	CircleSegmentCount uint32
	ArcSegmentCount    uint32
	CurveSegmentCount  uint32
	Null               DrawNullTexture
	vertexLayout       unsafe.Pointer
	VertexSize         uintptr
	VertexAlignment    uintptr
}

// ConvertConfigBuilder is used to build ConvertConfig values.
// It adds a Go-style slice of DrawVertexLayoutElement values to CConvertConfig
// so that the hidden C-style array can be filled in by Build.
type ConvertConfigBuilder struct {
	CConvertConfig
	VertexLayout []DrawVertexLayoutElement
}

// Build builds a ConvertConfig from ccb. The resulting value must be released
// by Free.
func (ccb *ConvertConfigBuilder) Build() *ConvertConfig {
	ptr := C.malloc(C.sizeof_struct_nk_convert_config)
	raw := (*CConvertConfig)(unsafe.Pointer(ptr))
	*raw = ccb.CConvertConfig
	raw.vertexLayout = C.malloc(C.ulong(1+len(ccb.VertexLayout)) * C.sizeof_struct_nk_draw_vertex_layout_element)
	for i := 0; i <= len(ccb.VertexLayout); i++ {
		dvlePtr := unsafe.Pointer(uintptr(raw.vertexLayout) + uintptr(i)*C.sizeof_struct_nk_draw_vertex_layout_element)
		dvleRaw := (*DrawVertexLayoutElement)(dvlePtr)
		if i == len(ccb.VertexLayout) {
			*dvleRaw = vertexLayoutEnd
		} else {
			*dvleRaw = ccb.VertexLayout[i]
		}
	}
	return (*ConvertConfig)(ptr)
}

type ConvertConfig C.struct_nk_convert_config

func (cnf *ConvertConfig) Free() {
	if cnf != nil {
		if cnf.vertex_layout != nil {
			C.free(unsafe.Pointer(cnf.vertex_layout))
		}
		C.free(unsafe.Pointer(cnf))
	}
}

func (cnf *ConvertConfig) raw() *C.struct_nk_convert_config {
	return (*C.struct_nk_convert_config)(cnf)
}

func (ctx *Context) Convert(cmds, vertices, elements *Buffer, config *ConvertConfig) error {
	// nk_flags nk_convert(struct nk_context*, struct nk_buffer *cmds, struct nk_buffer *vertices,
	//     struct nk_buffer *elements, const struct nk_convert_config*);
	result := C.nk_convert(
		ctx.raw(),
		cmds.raw(),
		vertices.raw(),
		elements.raw(),
		config.raw(),
	)
	if result != C.NK_CONVERT_SUCCESS {
		return ConvertError(result)
	}
	return nil
}
