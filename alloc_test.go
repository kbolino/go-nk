package nk_test

import (
	"testing"

	"github.com/kbolino/go-nk"
)

func TestFixedAlloc(t *testing.T) {
	var nkc nk.Context
	nkc.InitFixed(10)
	defer nkc.Free() // should not panic
}
