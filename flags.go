package nk

// #include "nk.h"
import "C"

// enum nk_edit_flags {
//     NK_EDIT_DEFAULT                 = 0,
//     NK_EDIT_READ_ONLY               = NK_FLAG(0),
//     NK_EDIT_AUTO_SELECT             = NK_FLAG(1),
//     NK_EDIT_SIG_ENTER               = NK_FLAG(2),
//     NK_EDIT_ALLOW_TAB               = NK_FLAG(3),
//     NK_EDIT_NO_CURSOR               = NK_FLAG(4),
//     NK_EDIT_SELECTABLE              = NK_FLAG(5),
//     NK_EDIT_CLIPBOARD               = NK_FLAG(6),
//     NK_EDIT_CTRL_ENTER_NEWLINE      = NK_FLAG(7),
//     NK_EDIT_NO_HORIZONTAL_SCROLL    = NK_FLAG(8),
//     NK_EDIT_ALWAYS_INSERT_MODE      = NK_FLAG(9),
//     NK_EDIT_MULTILINE               = NK_FLAG(10),
//     NK_EDIT_GOTO_END_ON_ACTIVATE    = NK_FLAG(11)
// };

type EditFlags uint32

const (
	EditDefault            EditFlags = C.NK_EDIT_DEFAULT
	EditReadOnly           EditFlags = C.NK_EDIT_READ_ONLY
	EditAutoSelect         EditFlags = C.NK_EDIT_AUTO_SELECT
	EditSigEnter           EditFlags = C.NK_EDIT_SIG_ENTER
	EditAllowTab           EditFlags = C.NK_EDIT_ALLOW_TAB
	EditNoCursor           EditFlags = C.NK_EDIT_NO_CURSOR
	EditSelectable         EditFlags = C.NK_EDIT_SELECTABLE
	EditClipboard          EditFlags = C.NK_EDIT_CLIPBOARD
	EditCtrlEnterNewline   EditFlags = C.NK_EDIT_CTRL_ENTER_NEWLINE
	EditNoHorizontalScroll EditFlags = C.NK_EDIT_NO_HORIZONTAL_SCROLL
	EditAlwaysinsertMode   EditFlags = C.NK_EDIT_ALWAYS_INSERT_MODE
	EditMultiline          EditFlags = C.NK_EDIT_MULTILINE
	EditGotoEndOnActivate  EditFlags = C.NK_EDIT_GOTO_END_ON_ACTIVATE
)

// enum nk_edit_types {
//     NK_EDIT_SIMPLE  = NK_EDIT_ALWAYS_INSERT_MODE,
//     NK_EDIT_FIELD   = NK_EDIT_SIMPLE|NK_EDIT_SELECTABLE|NK_EDIT_CLIPBOARD,
//     NK_EDIT_BOX     = NK_EDIT_ALWAYS_INSERT_MODE| NK_EDIT_SELECTABLE| NK_EDIT_MULTILINE|NK_EDIT_ALLOW_TAB|NK_EDIT_CLIPBOARD,
//     NK_EDIT_EDITOR  = NK_EDIT_SELECTABLE|NK_EDIT_MULTILINE|NK_EDIT_ALLOW_TAB| NK_EDIT_CLIPBOARD
// };

const (
	EditSimple EditFlags = C.NK_EDIT_SIMPLE
	EditField  EditFlags = C.NK_EDIT_FIELD
	EditBox    EditFlags = C.NK_EDIT_BOX
	EditEditor EditFlags = C.NK_EDIT_EDITOR
)

// enum nk_edit_events {
//     NK_EDIT_ACTIVE      = NK_FLAG(0), /* edit widget is currently being modified */
//     NK_EDIT_INACTIVE    = NK_FLAG(1), /* edit widget is not active and is not being modified */
//     NK_EDIT_ACTIVATED   = NK_FLAG(2), /* edit widget went from state inactive to state active */
//     NK_EDIT_DEACTIVATED = NK_FLAG(3), /* edit widget went from state active to state inactive */
//     NK_EDIT_COMMITED    = NK_FLAG(4) /* edit widget has received an enter and lost focus */
// };

type EditEvents uint32

const (
	EditActive      EditEvents = C.NK_EDIT_ACTIVE
	EditInactive    EditEvents = C.NK_EDIT_INACTIVE
	EditActivated   EditEvents = C.NK_EDIT_ACTIVATED
	EditDeactivated EditEvents = C.NK_EDIT_DEACTIVATED
	EditCommited    EditEvents = C.NK_EDIT_COMMITED
)
