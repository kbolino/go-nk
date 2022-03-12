# go-nk

[![Go Reference](https://pkg.go.dev/badge/github.com/kbolino/go-nk.svg)](https://pkg.go.dev/github.com/kbolino/go-nk)

A CGo binding for [Nuklear][nuklear]. The Nuklear header version 4.9.6 is
directly included in this module. All of the following flags are enabled:

- `NK_INCLUDE_DEFAULT_ALLOCATOR`
- `NK_INCLUDE_DEFAULT_FONT` (see `FontAtlas.AddDefaultFont`)
- `NK_INCLUDE_FIXED_TYPES`
- `NK_INCLUDE_FONT_BAKING` (see `FontAtlas` and `Font`)
- `NK_INCLUDE_STANDARD_BOOL`
- `NK_INCLUDE_STANDARD_IO`
- `NK_UINT_DRAW_INDEX`
- `NK_INCLUDE_VERTEX_BUFFER_OUTPUT` (see `Context.Convert` and
  `Context.DrawForEach`)

Most of these flags do not matter to users of this Go library since they only
affect its internals, but `NK_UINT_DRAW_INDEX` is notable as it affects the
output of `Context.Convert`: element indices are 32 bits wide instead of the
default 16 bits wide.

[nuklear]: https://github.com/Immediate-Mode-UI/Nuklear
