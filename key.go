package nk

// #include "nk.h"
import "C"

// Key represents all key input types that are meaningful to Nuklear.
type Key int32

const (
	KeyNone      Key = C.NK_KEY_NONE
	KeyShift     Key = C.NK_KEY_SHIFT
	KeyCtrl      Key = C.NK_KEY_CTRL
	KeyDel       Key = C.NK_KEY_DEL
	KeyEnter     Key = C.NK_KEY_ENTER
	KeyTab       Key = C.NK_KEY_TAB
	KeyBackspace Key = C.NK_KEY_BACKSPACE
	KeyCopy      Key = C.NK_KEY_COPY
	KeyCut       Key = C.NK_KEY_CUT
	KeyPaste     Key = C.NK_KEY_PASTE
	KeyUp        Key = C.NK_KEY_UP
	KeyDown      Key = C.NK_KEY_DOWN
	KeyLeft      Key = C.NK_KEY_LEFT
	KeyRight     Key = C.NK_KEY_RIGHT

	KeyTextInsertMode  Key = C.NK_KEY_TEXT_INSERT_MODE
	KeyTextReplaceMode Key = C.NK_KEY_TEXT_REPLACE_MODE
	KeyTextResetMode   Key = C.NK_KEY_TEXT_RESET_MODE
	KeyTextLineStart   Key = C.NK_KEY_TEXT_LINE_START
	KeyTextLineEnd     Key = C.NK_KEY_TEXT_LINE_END
	KeyTextStart       Key = C.NK_KEY_TEXT_START
	KeyTextEnd         Key = C.NK_KEY_TEXT_END
	KeyTextUndo        Key = C.NK_KEY_TEXT_UNDO
	KeyTextRedo        Key = C.NK_KEY_TEXT_REDO
	KeyTextSelectAll   Key = C.NK_KEY_TEXT_SELECT_ALL
	KeyTextWordLeft    Key = C.NK_KEY_TEXT_WORD_LEFT
	KeyTextWordRight   Key = C.NK_KEY_TEXT_WORD_RIGHT

	KeyScrollStart Key = C.NK_KEY_SCROLL_START
	KeyScrollEnd   Key = C.NK_KEY_SCROLL_END
	KeyScrollDown  Key = C.NK_KEY_SCROLL_DOWN
	KeyScrollUp    Key = C.NK_KEY_SCROLL_UP

	KeyMax Key = C.NK_KEY_MAX
)
