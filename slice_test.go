package nk

import (
	"testing"
	"unsafe"
)

const (
	length = 5
)

// TestFakeByteSlice is a regression test for an issue where fakeByteSlice
// was not setting the pointer of the slice correctly.
func TestFakeByteSlice(t *testing.T) {
	var b [length]byte
	slice := fakeByteSlice(unsafe.Pointer(&b[0]), length)
	if &b[0] != &slice[0] {
		t.Fatal("slice unequal")
	}
}
