# Change Log

## v0.14.0 (2022-03-25)

- Breaking API change: Removed `ScaleHeight` from `Font`, replaced with `Height`
  and `SetHeight` on `UserFont` to more closely match the C API

## v0.13.1 (2022-03-23)

- Added `DefaultSegmentCount` constant

## v0.13.0 (2022-03-17)

- Minimum required Go version is now 1.18.
- Removed `HandlePtr` and `GetHandlePtr` functions due to a
  [bug in Go][golang-51733]. Instead, proper use of `Handle` for both C memory
  and Go memory has been fully documented.

[golang-51733]: https://github.com/golang/go/issues/51733

## v0.12.0 (2022-03-16)

**RESCINDED RELEASE:** Use v0.13.0 or newer instead

- Minimum required Go version is now 1.18
- Added `HandlePtr` and `GetHandlePtr` generic functions to ergonomically and
  safely use `Handle`s

## v0.11.0 (2022-03-16)

- Breaking API change: The `XXXBytes` methods have been dropped. The maintenance
  burden is no longer worth the benefit now that string pooling is possible.
- Added `ScaleHeight` method on `Font`, which does what the old `SetScale` meant
  to do. See [this Nuklear PR][nk-pr-427] for details.

[nk-pr-427]: https://github.com/Immediate-Mode-UI/Nuklear/pull/427

## v0.10.1 (2022-03-16)

- New feature: C-style strings can be pooled using the new `CStringPool`
  interface and `SetCStringPool` function; the default implementation is no
  different than previous behavior (copy/free a C-style string with every
  function call)

## v0.10.0 (2022-03-14)

- Breaking API change: removed all specialized `XXXFlags` types, made their
  constants untyped, and replaced the types with generic `Flags` instead: this
  more closely matches the C API and prevents flag type issues from arising
  again like e.g. `PanelFlags` vs. `WindowFlags`

## v0.9.1 (2022-03-13)

- Fixed type passed to `malloc` in `ConvertConfigBuilder.Build`

## v0.9.0 (2022-03-13)

- Breaking API change: remodeled `ConvertConfigBuilder` to match the pattern
  used by `FontConfigBuilder`: it now takes a value receiver and doesn't have
  a nested `CConvertConfig`; `CConvertConfig` type is removed entirely
- Added `CoordType` and `FallbackGlypth` to `FontConfigBuilder`

## v0.8.1 (2022-03-13)

- Use value receiver on `FontConfigBuilder.Build` so that it can be called from
  inline builder value

## v0.8.0 (2022-03-13)

- Breaking API change: removed `Scale` and `SetScale` from `Font` as they
  didn't seem to do anything
- Breaking API change: added `config *FontConfig` parameter to
  `FontAtlas.AddXXX` methods
- Added `FontConfig` and `FontConfigBuilder` types to control how fonts are
  baked

## v0.7.1 (2022-03-13)

- Add `Scale` and `SetScale` to `Font` so that it can be scaled for high-DPI
  displays

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
