package nk

// #include <stdlib.h>
import "C"

import "unsafe"

// CStringPool defines the interface for string pooling. Use of this interface
// and SetCStringPool is entirely optional. The primary purpose of string
// pooling is to reduce unnecessary heap allocations.
type CStringPool interface {
	// Get provides the C-style string for s. Requirements for the return value:
	//
	//    - must not be nil (if allocation fails, panic instead)
	//    - target must be stored outside Go memory (e.g. obtained from malloc)
	//    - target must be identical to s for its first len(s) bytes
	//    - target must be at least 1 byte larger in size than s
	//    - target must be NUL-terminated, i.e. the byte at index len(s) should be 0
	//
	// These requirements are already met by the cgo built-in C.CString
	// function.
	//
	// The returned value must be given back to the pool when no longer needed
	// using Release.
	Get(s string) *C.char

	// Release returns the C-style string cs back to the pool.
	//
	// The pointer cs must have been returned by a previous call to cs, and the
	// same pointer should not be returned to the pool more than once. That is,
	// every call to Get should be matched by a single call to Release, and vice
	// versa.
	Release(cs *C.char)
}

// SetCStringPool sets the string pool used to obtain C-style strings whenever
// they are needed. If set to nil or left unset, then C-style strings are
// created from Go strings using cgo's C.CString function, which copies the
// string into C memory, and then freed at the end of the enclosing call.
//
// SetCStringPool should be called to set the string pool once, before creating
// any types or calling any other functions.
//
// Concurrency: Whether pool is safe for concurrent use by multiple goroutines
// simultaneously is up to the implementation. If the other functions in this
// library are only ever called from the same goroutine, then the pool does not
// need to be concurrency-safe. The default implementation is stateless and thus
// its concurrency-safety is the same as for the C functions malloc and free.
func SetCStringPool(pool CStringPool) {
	if pool == nil {
		cStringPool = defaultStringPool{}
	} else {
		cStringPool = pool
	}
}

var cStringPool CStringPool = defaultStringPool{}

// defaultStringPool implements CStringPool by calling malloc/free every time.
type defaultStringPool struct{}

func (defaultStringPool) Get(s string) *C.char { return C.CString(s) }
func (defaultStringPool) Release(cs *C.char)   { C.free(unsafe.Pointer(cs)) }
