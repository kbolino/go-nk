#ifndef NK_H_
#define NK_H_

#define NK_INCLUDE_DEFAULT_ALLOCATOR
#define NK_INCLUDE_DEFAULT_FONT
#define NK_INCLUDE_FIXED_TYPES
#define NK_INCLUDE_FONT_BAKING
#define NK_INCLUDE_STANDARD_BOOL
#define NK_INCLUDE_VERTEX_BUFFER_OUTPUT
#include "nuklear.h"

// DVLE_SIZE is the size of the nk_draw_vertex_layout_element struct.
#define DVLE_SIZE sizeof(struct nk_draw_vertex_layout_element)

// find_vertex_layout_count returns the length of the vertex layout element
// array starting with start.
size_t find_vertex_layout_count(struct nk_draw_vertex_layout_element *start);

#endif
