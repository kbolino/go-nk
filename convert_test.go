package nk_test

import (
	"testing"
	"unsafe"

	"github.com/kbolino/go-nk"
)

func TestConvert(t *testing.T) {
	nkc, err := nk.NewContext()
	if err != nil {
		t.Fatal("unexpected init error: %w", err)
	}
	defer nkc.Free()
	cbuf := nk.NewBuffer()
	defer cbuf.Free()
	ebuf := nk.NewBuffer()
	defer ebuf.Free()
	vbuf := nk.NewBuffer()
	defer vbuf.Free()
	configBuilder := nk.ConvertConfigBuilder{
		CConvertConfig: nk.CConvertConfig{
			VertexSize:      20, // sizeof(sdl.Vertex)
			VertexAlignment: 4,  // alignof(sdl.Vertex)
			Null: nk.DrawNullTexture{
				Texture: nk.Handle(unsafe.Pointer(nil)),
				UV:      nk.Vec2{X: 0, Y: 0},
			},
			CircleSegmentCount: 10,
			CurveSegmentCount:  10,
			ArcSegmentCount:    10,
			GlobalAlpha:        1.0,
			ShapeAA:            nk.AntiAliasingOff,
			LineAA:             nk.AntiAliasingOff,
		},
		VertexLayout: []nk.DrawVertexLayoutElement{
			{Attribute: nk.VertexPosition, Format: nk.FormatFloat, Offset: 0},
			{Attribute: nk.VertexTexcoord, Format: nk.FormatFloat, Offset: 8},
			{Attribute: nk.VertexColor, Format: nk.FormatR8G8B8A8, Offset: 16},
		},
	}
	config := configBuilder.Build()
	defer config.Free()
	if err := nkc.Convert(cbuf, vbuf, ebuf, config); err != nil {
		t.Fatal("unexpected convert error:", err)
	}
}
