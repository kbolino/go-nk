package nk

// #include "nk.h"
import "C"

type Flags uint32

// enum nk_panel_flags {
//     NK_WINDOW_BORDER            = NK_FLAG(0),
//     NK_WINDOW_MOVABLE           = NK_FLAG(1),
//     NK_WINDOW_SCALABLE          = NK_FLAG(2),
//     NK_WINDOW_CLOSABLE          = NK_FLAG(3),
//     NK_WINDOW_MINIMIZABLE       = NK_FLAG(4),
//     NK_WINDOW_NO_SCROLLBAR      = NK_FLAG(5),
//     NK_WINDOW_TITLE             = NK_FLAG(6),
//     NK_WINDOW_SCROLL_AUTO_HIDE  = NK_FLAG(7),
//     NK_WINDOW_BACKGROUND        = NK_FLAG(8),
//     NK_WINDOW_SCALE_LEFT        = NK_FLAG(9),
//     NK_WINDOW_NO_INPUT          = NK_FLAG(10)
// };

// enum nk_panel_flags constants
const (
	WindowBorder         = C.NK_WINDOW_BORDER
	WindowMovable        = C.NK_WINDOW_MOVABLE
	WindowScalable       = C.NK_WINDOW_SCALABLE
	WindowClosable       = C.NK_WINDOW_CLOSABLE
	WindowMinimizable    = C.NK_WINDOW_MINIMIZABLE
	WindowNoScrollbar    = C.NK_WINDOW_NO_SCROLLBAR
	WindowTitle          = C.NK_WINDOW_TITLE
	WindowScrollAutoHide = C.NK_WINDOW_SCROLL_AUTO_HIDE
	WindowBackground     = C.NK_WINDOW_BACKGROUND
	WindowScaleLeft      = C.NK_WINDOW_SCALE_LEFT
	WindowNoInput        = C.NK_WINDOW_NO_INPUT
)

// enum nk_window_flags {
//     NK_WINDOW_PRIVATE       = NK_FLAG(11),
//     NK_WINDOW_DYNAMIC       = NK_WINDOW_PRIVATE,
//     /* special window type growing up in height while being filled to a certain maximum height */
//     NK_WINDOW_ROM           = NK_FLAG(12),
//     /* sets window widgets into a read only mode and does not allow input changes */
//     NK_WINDOW_NOT_INTERACTIVE = NK_WINDOW_ROM|NK_WINDOW_NO_INPUT,
//     /* prevents all interaction caused by input to either window or widgets inside */
//     NK_WINDOW_HIDDEN        = NK_FLAG(13),
//     /* Hides window and stops any window interaction and drawing */
//     NK_WINDOW_CLOSED        = NK_FLAG(14),
//     /* Directly closes and frees the window at the end of the frame */
//     NK_WINDOW_MINIMIZED     = NK_FLAG(15),
//     /* marks the window as minimized */
//     NK_WINDOW_REMOVE_ROM    = NK_FLAG(16)
//     /* Removes read only mode at the end of the window */
// };

// enum nk_window_flags constants
const (
	WindowPrivate        = C.NK_WINDOW_PRIVATE
	WindowDynamic        = C.NK_WINDOW_DYNAMIC
	WindowROM            = C.NK_WINDOW_ROM
	WindowNotInteractive = C.NK_WINDOW_NOT_INTERACTIVE
	WindowHidden         = C.NK_WINDOW_HIDDEN
	WindowClosed         = C.NK_WINDOW_CLOSED
	WindowMinimized      = C.NK_WINDOW_MINIMIZED
	WindowRemoveROM      = C.NK_WINDOW_REMOVE_ROM
)
