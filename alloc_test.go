package nk_test

import (
	"testing"

	"github.com/kbolino/go-nk"
)

func TestInitFree(t *testing.T) {
	nkc, err := nk.NewContext()
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	nkc.Free() // should not panic
}

func TestBufferInitFree(t *testing.T) {
	buf := nk.NewBuffer()
	buf.Free()
}
