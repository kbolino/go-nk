package nk_test

import (
	"testing"

	"github.com/kbolino/go-nk"
)

func TestInitFree(t *testing.T) {
	var nkc nk.Context
	if err := nkc.InitDefault(); err != nil {
		t.Fatal("unexpected error:", err)
	}
	defer nkc.Free() // should not panic
}

func TestBufferInitFree(t *testing.T) {
	var buf nk.Buffer
	buf.InitDefault()
	defer buf.Free()
}
