package key

import (
	"strings"
)

type Key string

const (
    Escape Key = "escape"

    Up Key = "up"
    Down Key = "down"
    Right Key = "right"
    Left Key = "left"

    Home Key = "home"
    End Key = "end"
    Delete Key = "delete"
    ControlDelete Key = "c-delete"
    PageUp Key = "pageup"
    PageDown Key = "pagedown"
    Insert Key = "insert"

    ControlAt Key = "c-@"  // c-space
    ControlA Key = "c-a"
    ControlB Key = "c-b"
    ControlC Key = "c-c"
    ControlD Key = "c-d"
    ControlE Key = "c-e"
    ControlF Key = "c-f"
    ControlG Key = "c-g"
    ControlH Key = "c-h"
    ControlI Key = "c-i"  // tab
    ControlJ Key = "c-j"  // newline
    ControlK Key = "c-k"
    ControlL Key = "c-l"
    ControlM Key = "c-m"  // cr
    ControlN Key = "c-n"
    ControlO Key = "c-o"
    ControlP Key = "c-p"
    ControlQ Key = "c-q"
    ControlR Key = "c-r"
    ControlS Key = "c-s"
    ControlT Key = "c-t"
    ControlU Key = "c-u"
    ControlV Key = "c-v"
    ControlW Key = "c-w"
    ControlX Key = "c-x"
    ControlY Key = "c-y"
    ControlZ Key = "c-z"

    ControlBackslash Key = "c-\\"
    ControlSquareClose Key = "c-]"
    ControlCircumflex Key = "c-^"
    ControlUnderscore Key = "c-_"

    ControlLeft Key = "c-left"
    ControlRight Key = "c-right"
    ControlUp Key = "c-up"
    ControlDown Key = "c-down"

    ShiftLeft Key = "s-left"
    ShiftUp Key = "s-up"
    ShiftDown Key = "s-down"
    ShiftRight Key = "s-right"
    ShiftDelete Key = "s-delete"
    BackTab Key = "s-tab"

    F1 Key = "f1"
    F2 Key = "f2"
    F3 Key = "f3"
    F4 Key = "f4"
    F5 Key = "f5"
    F6 Key = "f6"
    F7 Key = "f7"
    F8 Key = "f8"
    F9 Key = "f9"
    F10 Key = "f10"
    F11 Key = "f11"
    F12 Key = "f12"
    F13 Key = "f13"
    F14 Key = "f14"
    F15 Key = "f15"
    F16 Key = "f16"
    F17 Key = "f17"
    F18 Key = "f18"
    F19 Key = "f19"
    F20 Key = "f20"
    F21 Key = "f21"
    F22 Key = "f22"
    F23 Key = "f23"
    F24 Key = "f24"

    ScrollUp Key = "<scroll-up>"
    ScrollDown Key = "<scroll-down>"

    CPRResponse Key = "<cursor-position-response>"
    Vt100MouseEvent Key = "<vt100-mouse-event>"
    WindowsMouseEvent Key = "<windows-mouse-event>"
    BracketedPaste Key = "<bracketed-paste>"

    Ignore Key = "<ignore>"

    Any Key = "<any>"
)

// Aliases
const (
	BackSpace Key = "backspace"
	ControlSpace Key = "c-space"
	Enter Key = "enter"
	Tab Key = "tab"
)

var keyAliasMap = map[Key]Key {
	BackSpace: ControlH,
	ControlSpace: ControlAt,
	Enter: ControlM,
	Tab: ControlI,
}

const (
	// special Key to allow key processor to flush all pending keypresses
	Flush Key = "<flush>"
)

var CtrlKeyMap = map[string]([]Key) {
	"\x00": []Key{ControlAt},
    "\x01": []Key{ControlA},
    "\x02": []Key{ControlB},
    "\x03": []Key{ControlC},
    "\x04": []Key{ControlD},
    "\x05": []Key{ControlE},
    "\x06": []Key{ControlF},
    "\x07": []Key{ControlG},
    "\x08": []Key{ControlH},
    "\x09": []Key{ControlI},
    "\x0a": []Key{ControlJ},
    "\x0b": []Key{ControlK},
    "\x0c": []Key{ControlL},
    "\x0d": []Key{ControlM},
    "\x0e": []Key{ControlN},
    "\x0f": []Key{ControlO},
    "\x10": []Key{ControlP},
    "\x11": []Key{ControlQ},
    "\x12": []Key{ControlR},
    "\x13": []Key{ControlS},
    "\x14": []Key{ControlT},
    "\x15": []Key{ControlU},
    "\x16": []Key{ControlV},
    "\x17": []Key{ControlW},
    "\x18": []Key{ControlX},
    "\x19": []Key{ControlY},
    "\x1a": []Key{ControlZ},

    "\x1b": []Key{Escape},
    "\x1c": []Key{ControlBackslash},
    "\x1d": []Key{ControlSquareClose},
    "\x1e": []Key{ControlCircumflex},
    "\x1f": []Key{ControlUnderscore},
    "\x7f": []Key{ControlH},
}

var aNSIKeyMap = map[string]([]Key) {
	"\x1b[A": []Key{Up},
	"\x1b[B": []Key{Down},
	"\x1b[C": []Key{Right},
	"\x1b[D": []Key{Left},
	"\x1b[H": []Key{Home},
	"\x1bOH": []Key{Home},
	"\x1b[F": []Key{End},
	"\x1bOF": []Key{End},
	"\x1b[3~": []Key{Delete},
	"\x1b[3;2~": []Key{ShiftDelete},
	"\x1b[3;5~": []Key{ControlDelete},
	"\x1b[1~": []Key{Home},
	"\x1b[4~": []Key{End},
	"\x1b[5~": []Key{PageUp},
	"\x1b[6~": []Key{PageDown},
	"\x1b[7~": []Key{Home},
	"\x1b[8~": []Key{End},
	"\x1b[Z": []Key{BackTab},
	"\x1b[2~": []Key{Insert},

	"\x1bOP": []Key{F1},
	"\x1bOQ": []Key{F2},
	"\x1bOR": []Key{F3},
	"\x1bOS": []Key{F4},
	"\x1b[[A": []Key{F1},
	"\x1b[[B": []Key{F2},
	"\x1b[[C": []Key{F3},
	"\x1b[[D": []Key{F4},
	"\x1b[[E": []Key{F5},
	"\x1b[11~": []Key{F1},
	"\x1b[12~": []Key{F2},
	"\x1b[13~": []Key{F3},
	"\x1b[14~": []Key{F4},
	"\x1b[15~": []Key{F5},
	"\x1b[17~": []Key{F6},
	"\x1b[18~": []Key{F7},
	"\x1b[19~": []Key{F8},
	"\x1b[20~": []Key{F9},
	"\x1b[21~": []Key{F10},
	"\x1b[23~": []Key{F11},
	"\x1b[24~": []Key{F12},
	"\x1b[25~": []Key{F13},
	"\x1b[26~": []Key{F14},
	"\x1b[28~": []Key{F15},
	"\x1b[29~": []Key{F16},
	"\x1b[31~": []Key{F17},
	"\x1b[32~": []Key{F18},
	"\x1b[33~": []Key{F19},
	"\x1b[34~": []Key{F20},

	// Xterm
	"\x1b[1;2P": []Key{F13},
	"\x1b[1;2Q": []Key{F14},
	// "\x1b[1;2R": []Key{F15},  # Conflicts with CPR response.
	"\x1b[1;2S": []Key{F16},
	"\x1b[15;2~": []Key{F17},
	"\x1b[17;2~": []Key{F18},
	"\x1b[18;2~": []Key{F19},
	"\x1b[19;2~": []Key{F20},
	"\x1b[20;2~": []Key{F21},
	"\x1b[21;2~": []Key{F22},
	"\x1b[23;2~": []Key{F23},
	"\x1b[24;2~": []Key{F24},

	"\x1b[1;5A": []Key{ControlUp},
	"\x1b[1;5B": []Key{ControlDown},
	"\x1b[1;5C": []Key{ControlRight},
	"\x1b[1;5D": []Key{ControlLeft},

	"\x1b[1;2A": []Key{ShiftUp},
	"\x1b[1;2B": []Key{ShiftDown},
	"\x1b[1;2C": []Key{ShiftRight},
	"\x1b[1;2D": []Key{ShiftLeft},

	"\x1bOA": []Key{Up},
	"\x1bOB": []Key{Down},
	"\x1bOC": []Key{Right},
	"\x1bOD": []Key{Left},

	"\x1b[5A": []Key{ControlUp},
	"\x1b[5B": []Key{ControlDown},
	"\x1b[5C": []Key{ControlRight},
	"\x1b[5D": []Key{ControlLeft},

	"\x1bOc": []Key{ControlRight},
	"\x1bOd": []Key{ControlLeft},

	// Tmux (Win32 subsystem) sends the following scroll events. Ignored for now.
	"\x1b[62~": []Key{ScrollUp},
	"\x1b[63~": []Key{ScrollDown},

	"\x1b[200~": []Key{BracketedPaste},  // Start of bracketed paste.

	"\x1b[1;3D": []Key{Escape, Left},
	"\x1b[1;3C": []Key{Escape, Right},
	"\x1b[1;3A": []Key{Escape, Up},
	"\x1b[1;3B": []Key{Escape, Down},
	"\x1b[1;9D": []Key{Escape, Left},
	"\x1b[1;9C": []Key{Escape, Right},

	"\x1b[E": []Key{Ignore},
	"\x1b[G": []Key{Ignore},
}

func init() {
	for k, v := range CtrlKeyMap {
		aNSIKeyMap[k] = v
	}
}

// TODO: memoize
func isPrefixOfANSI(s string) bool {
	for k := range aNSIKeyMap {
		if s != k && strings.HasPrefix(k, s) {
			return true
		}
	}
	return false
}
