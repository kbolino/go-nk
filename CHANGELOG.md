# Change Log

## v0.7.0 (2022-03-12)

- Breaking API change: Set flag `NK_UINT_DRAW_INDEX` in `nk.h` to make
 `nk_draw_index` an `unsigned int` instead of an `unsigned short`; this more
  closely matches what SDL2 expects, but causes the output from
  `Context.Convert` to the `elements` buffer to change incompatibly
- Breaking API change: the `DrawList` struct is no longer exposed, as it is not
  currently used by the Go API in any way
- Several types and methods have been documented

## v0.6.1 (2022-03-11)

- Added the `CheckText` and `CheckTextBytes` methods to `Context`

## v0.6.0 (2022-03-11)

- Breaking API change: Fixed the spelling of `AntiAliasingOn` constant

## v0.5.2 (2022-03-11)

- Bug fix: Fixed the broken `fakeByteSlice` implementation, which was causing
  `Buffer.Memory` to return an invalid slice; added regression test as well

## v0.5.1 (2022-03-11)

- Added `Font.Handle` to get the associated `UserFont` handle, and added
  `Context.StyleSetFont` to actually use the font

## v0.5.0 (2022-03-11)

- Breaking API change: The `width` and `height` of a baked font image are
  strictly out parameters, so `widthIn` and `heightIn` were completely ignored;
  they have been dropped from the `FontAtlas.Bake` API signature accordingly

## v0.4.3 (2022-03-11)

- Allow `Context.Begin` to take more generic `PanelFlags` instead of more
  specific `WindowFlags`

## v0.4.2 (2022-03-11)

- Bug fix: Use `nk_buffer->needed` instead of `nk_buffer->size` as the "size"
  of `Buffer` memory, to match nuklear's bundled examples

## v0.4.1 (2022-03-11)

- Added a way to retrieve `Buffer` memory

## v0.4.0 (2022-03-11)

- Breaking API change: all memory is now allocated with the C stdlib, including
  for apparent Go types; `InitDefault` has been dropped and structs should be
  created with `NewXXX` methods instead (which return pointers); this is
  because even stack-allocated Go structs still count as "in Go memory" and so
  pointers to them which get nested in other structs will still violate cgo
  rules

## v0.3.0 (2022-03-11)

- Breaking API change: moved the `DrawVertexListElement` array in
  `ConvertConfig` to C memory, again to deal with the cgo rules

## v0.2.0 (2022-03-11)

- Breaking API change: `InitFixed` dropped and replaced with `InitDefault`
- The fixed allocator implementation was breaking cgo's rule about passing
  nested Go pointers

## v0.1.1 (2022-03-11)

- Bug fix: Prevent a panic when calling `Free` on `Buffer` and `Context`

## v0.1.0 (2022-03-11)

- Initial implementation