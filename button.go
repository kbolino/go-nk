package nk

// #include "nk.h"
import "C"

// Button represents all of the button input events that are meaningful to
// Nuklear.
type Button int32

const (
	ButtonLeft   Button = C.NK_BUTTON_LEFT
	ButtonMiddle Button = C.NK_BUTTON_MIDDLE
	ButtonRight  Button = C.NK_BUTTON_RIGHT
	ButtonDouble Button = C.NK_BUTTON_DOUBLE
	ButtonMax    Button = C.NK_BUTTON_MAX
)
