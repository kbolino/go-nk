package nk_test

import (
	"testing"

	"github.com/kbolino/go-nk"
)

func TestBuffer(t *testing.T) {
	buf1 := nk.NewBuffer()
	t.Logf("buf1 = %p", buf1)
	buf2 := nk.NewBuffer()
	t.Logf("buf2 = %p", buf2)
	if buf1 == buf2 {
		t.Fatal("buffers are not distinct")
	}
}
