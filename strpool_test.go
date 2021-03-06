package nk

import "testing"

func BenchmarkDefaultCStringPool_Direct(b *testing.B) {
	var stringPool defaultStringPool
	for i := 0; i < b.N; i++ {
		str := stringPool.Get("hello, world")
		stringPool.Release(str)
	}
}

func BenchmarkDefaultCStringPool_Interface(b *testing.B) {
	var stringPool CStringPool = defaultStringPool{}
	for i := 0; i < b.N; i++ {
		str := stringPool.Get("hello, world")
		stringPool.Release(str)
	}
}

func BenchmarkDummyCStringPool_Direct(b *testing.B) {
	var stringPool dummyStringPool
	stringPool.init("hello, world")
	defer stringPool.free()
	for i := 0; i < b.N; i++ {
		str := stringPool.Get("hello, world")
		stringPool.Release(str)
	}
}

func BenchmarkDummyCStringPool_Interface(b *testing.B) {
	var rawStringPool dummyStringPool
	rawStringPool.init("hello, world")
	defer rawStringPool.free()
	var stringPool CStringPool = &rawStringPool
	for i := 0; i < b.N; i++ {
		str := stringPool.Get("hello, world")
		stringPool.Release(str)
	}
}
