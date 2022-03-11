#define NK_IMPLEMENTATION
#include "nk.h"
#include <string.h>

size_t find_vertex_layout_count(struct nk_draw_vertex_layout_element *start) {
    size_t i;
    struct nk_draw_vertex_layout_element end = {NK_VERTEX_LAYOUT_END};
    for (i = 0;; i++) {
        if (memcmp(&start[i], &end, DVLE_SIZE) == 0) break;
    }
    return i;
}
