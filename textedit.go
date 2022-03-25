package nk

// #include "nk.h"
import "C"
import "unsafe"

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

// EditFlags specify the options for text edit widgets
type EditFlags uint32

// enum nk_edit_flags constants; individually toggleable flags
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

// enum nk_edit_types constants; prepackaged combinations of flags
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

// enum nk_edit_events constants; individually toggleable flags
const (
	EditActive      EditEvents = C.NK_EDIT_ACTIVE
	EditInactive    EditEvents = C.NK_EDIT_INACTIVE
	EditActivated   EditEvents = C.NK_EDIT_ACTIVATED
	EditDeactivated EditEvents = C.NK_EDIT_DEACTIVATED
	EditCommited    EditEvents = C.NK_EDIT_COMMITED
)

// PluginFilter controls which characters may be inserted into a text edit
// widget. Currently, only the predefined values are supported.
type PluginFilter C.nk_plugin_filter

// nk_bool nk_filter_default(const struct nk_text_edit*, nk_rune unicode);
// nk_bool nk_filter_ascii(const struct nk_text_edit*, nk_rune unicode);
// nk_bool nk_filter_float(const struct nk_text_edit*, nk_rune unicode);
// nk_bool nk_filter_decimal(const struct nk_text_edit*, nk_rune unicode);
// nk_bool nk_filter_hex(const struct nk_text_edit*, nk_rune unicode);
// nk_bool nk_filter_oct(const struct nk_text_edit*, nk_rune unicode);
// nk_bool nk_filter_binary(const struct nk_text_edit*, nk_rune unicode);

// Predefined filters
var (
	FilterDefault = PluginFilter(C.nk_filter_default)
	FilterASCII   = PluginFilter(C.nk_filter_ascii)
	FilterFloat   = PluginFilter(C.nk_filter_float)
	FilterDecimal = PluginFilter(C.nk_filter_decimal)
	FilterHex     = PluginFilter(C.nk_filter_hex)
	FilterOct     = PluginFilter(C.nk_filter_oct)
	FilterBinary  = PluginFilter(C.nk_filter_binary)
)

// EditString calls nk_edit_string which inserts a text edit widget.
// The initial string to edit is taken from the first n bytes of buffer, the
// maximum string length allowed is taken from len(buffer), and the resulting
// string length is returned. The filter paramater controls the allowable
// input. The flags input sets the options for the widget. The EditEvents return
// provides the current widget state.
func (ctx *Context) EditString(flags EditFlags, buffer []byte, n int, filter PluginFilter) (int, EditEvents) {
	length := C.int(n)
	// nk_flags nk_edit_string(struct nk_context*, nk_flags, char *buffer, int *len, int max, nk_plugin_filter)
	flagsOut := EditEvents(C.nk_edit_string(
		ctx.raw(),
		C.nk_flags(flags),
		(*C.char)(unsafe.Pointer(&buffer[0])),
		&length,
		C.int(len(buffer)),
		C.nk_plugin_filter(filter),
	))
	return int(length), flagsOut
}
