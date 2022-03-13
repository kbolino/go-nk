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

// ConvertError is any non-successful result from Context.Convert.
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

// ConvertConfigBuilder is used to construct ConvertConfig values.
type ConvertConfigBuilder struct {
	GlobalAlpha        float32                   // Global alpha value
	LineAA, ShapeAA    AntiAliasing              // Antialiasing options for lines and shapes
	CircleSegmentCount uint32                    // Number of segments used to draw circles (recommended 22)
	ArcSegmentCount    uint32                    // Number of segments used to draw arcs (recommended 22)
	CurveSegmentCount  uint32                    // Number of segments used to draw curves (recommended 22)
	Null               DrawNullTexture           // Null texture, used for drawing pixels
	VertexLayout       []DrawVertexLayoutElement // Describes the vertex elements and their packing
	VertexSize         uint32                    // Size of each vertex
	VertexAlignment    uint32                    // Alignment of the vertex type
}

// Build builds a ConvertConfig from ccb. The resulting value is stored in C
// memory and must be released by Free when no longer needed.
func (ccb ConvertConfigBuilder) Build() *ConvertConfig {
	raw := (*C.struct_nk_convert_config)(C.malloc(C.sizeof_struct_nk_convert_config))
	*raw = C.struct_nk_convert_config{
		global_alpha:         C.float(ccb.GlobalAlpha),
		line_AA:              C.enum_nk_anti_aliasing(ccb.LineAA),
		shape_AA:             C.enum_nk_anti_aliasing(ccb.ShapeAA),
		circle_segment_count: C.uint(ccb.CircleSegmentCount),
		arc_segment_count:    C.uint(ccb.ArcSegmentCount),
		curve_segment_count:  C.uint(ccb.CurveSegmentCount),
		null:                 *(*C.struct_nk_draw_null_texture)(unsafe.Pointer((&ccb.Null))),
		vertex_layout:        nil,
		vertex_size:          C.nk_size(ccb.VertexSize),
		vertex_alignment:     C.nk_size(ccb.VertexAlignment),
	}
	raw.vertex_layout = (*C.struct_nk_draw_vertex_layout_element)(C.malloc(
		C.ulong(1+len(ccb.VertexLayout)) * C.sizeof_struct_nk_draw_vertex_layout_element))
	for i := 0; i <= len(ccb.VertexLayout); i++ {
		elem := (*DrawVertexLayoutElement)(unsafe.Pointer(uintptr(unsafe.Pointer(raw.vertex_layout)) +
			uintptr(i)*C.sizeof_struct_nk_draw_vertex_layout_element))
		if i == len(ccb.VertexLayout) {
			*elem = vertexLayoutEnd
		} else {
			*elem = ccb.VertexLayout[i]
		}
	}
	return (*ConvertConfig)(raw)
}

// ConvertConfig is the opaque configuration for the Context.Convert method. To
// create such a value, use ConvertConfigBuilder.
type ConvertConfig C.struct_nk_convert_config

// Free releases memory used by cnf. Free is nil-safe. After this call, if cnf
// was not nil, it is now a dangling pointer.
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

// Convert calls nk_convert which converts the queued commands to draw commands
// operating on vertex buffers. The three supplied buffers are used to store
// the commands, the vertices (including positions, colors, and texcoords), and
// the indices into the vertices for each command, respectively. Convert
// operates according to the supplied configuration. None of the parameters
// can be nil.
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
