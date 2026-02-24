package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

// uinput constants
const (
	UINPUT_MAX_NAME_SIZE = 80
	EV_KEY               = 0x01
	EV_SYN               = 0x00

	// Key codes
	KEY_RESERVED       = 0
	KEY_ESC            = 1
	KEY_1              = 2
	KEY_2              = 3
	KEY_3              = 4
	KEY_4              = 5
	KEY_5              = 6
	KEY_6              = 7
	KEY_7              = 8
	KEY_8              = 9
	KEY_9              = 10
	KEY_0              = 11
	KEY_MINUS          = 12
	KEY_EQUAL          = 13
	KEY_BACKSPACE      = 14
	KEY_TAB            = 15
	KEY_Q              = 16
	KEY_W              = 17
	KEY_E              = 18
	KEY_R              = 19
	KEY_T              = 20
	KEY_Y              = 21
	KEY_U              = 22
	KEY_I              = 23
	KEY_O              = 24
	KEY_P              = 25
	KEY_LEFTBRACE      = 26
	KEY_RIGHTBRACE     = 27
	KEY_ENTER          = 28
	KEY_LEFTCTRL       = 29
	KEY_A              = 30
	KEY_S              = 31
	KEY_D              = 32
	KEY_F              = 33
	KEY_G              = 34
	KEY_H              = 35
	KEY_J              = 36
	KEY_K              = 37
	KEY_L              = 38
	KEY_SEMICOLON      = 39
	KEY_APOSTROPHE     = 40
	KEY_GRAVE          = 41
	KEY_LEFTSHIFT      = 42
	KEY_BACKSLASH      = 43
	KEY_Z              = 44
	KEY_X              = 45
	KEY_C              = 46
	KEY_V              = 47
	KEY_B              = 48
	KEY_N              = 49
	KEY_M              = 50
	KEY_COMMA          = 51
	KEY_DOT            = 52
	KEY_SLASH          = 53
	KEY_RIGHTSHIFT     = 54
	KEY_KPSTAR         = 55
	KEY_LEFTALT        = 56
	KEY_SPACE          = 57
	KEY_CAPSLOCK       = 58
	KEY_F1             = 59
	KEY_F2             = 60
	KEY_F3             = 61
	KEY_F4             = 62
	KEY_F5             = 63
	KEY_F6             = 64
	KEY_F7             = 65
	KEY_F8             = 66
	KEY_F9             = 67
	KEY_F10            = 68
	KEY_NUMLOCK        = 69
	KEY_SCROLLLOCK     = 70
	KEY_KP7            = 71
	KEY_KP8            = 72
	KEY_KP9            = 73
	KEY_KPMINUS        = 74
	KEY_KP4            = 75
	KEY_KP5            = 76
	KEY_KP6            = 77
	KEY_KPPLUS         = 78
	KEY_KP1            = 79
	KEY_KP2            = 80
	KEY_KP3            = 81
	KEY_KP0            = 82
	KEY_KPDOT          = 83
	KEY_ZENKAKUHANKAKU = 85
	KEY_F11            = 87
	KEY_F12            = 88
	KEY_RO             = 89
	KEY_KATAKANA       = 94
	KEY_HIRAGANA       = 93
	KEY_HENKAN         = 92
	KEY_KATAKANAHIRAGANA = 91
	KEY_MUHENKAN       = 86
	KEY_KPENTER        = 96
	KEY_RIGHTCTRL      = 97
	KEY_KPSLASH        = 98
	KEY_SYSRQ          = 99
	KEY_RIGHTALT       = 100
	KEY_HOME           = 102
	KEY_UP             = 103
	KEY_PAGEUP         = 104
	KEY_LEFT           = 105
	KEY_RIGHT          = 106
	KEY_END            = 107
	KEY_DOWN           = 108
	KEY_PAGEDOWN       = 109
	KEY_INSERT         = 110
	KEY_DELETE         = 111
	KEY_MAIL           = 112
	KEY_WEBHOME        = 172
	KEY_MUTE           = 113
	KEY_VOLUMEDOWN     = 114
	KEY_VOLUMEUP       = 115
	KEY_POWER          = 116
	KEY_KPEQUAL        = 117
	KEY_PAUSE          = 119
	KEY_KPCOMMA        = 121
	KEY_LEFTMETA       = 125
	KEY_RIGHTMETA      = 126
	KEY_COMPOSE        = 127

	// Modifier masks
	MOD_CTRL  = 1 << 0
	MOD_SHIFT = 1 << 1
	MOD_ALT   = 1 << 2
	MOD_META  = 1 << 3
)

type uinputSetup struct {
	id      [16]byte
	name    [UINPUT_MAX_NAME_SIZE]byte
	ffEffectsMax uint32
}

// Key name to code mapping
var keyCodeMap = map[string]int{
	"esc":           KEY_ESC,
	"escape":        KEY_ESC,
	"1":             KEY_1,
	"2":             KEY_2,
	"3":             KEY_3,
	"4":             KEY_4,
	"5":             KEY_5,
	"6":             KEY_6,
	"7":             KEY_7,
	"8":             KEY_8,
	"9":             KEY_9,
	"0":             KEY_0,
	"-":             KEY_MINUS,
	"=":             KEY_EQUAL,
	"backspace":     KEY_BACKSPACE,
	"tab":           KEY_TAB,
	"q":             KEY_Q,
	"w":             KEY_W,
	"e":             KEY_E,
	"r":             KEY_R,
	"t":             KEY_T,
	"y":             KEY_Y,
	"u":             KEY_U,
	"i":             KEY_I,
	"o":             KEY_O,
	"p":             KEY_P,
	"[":             KEY_LEFTBRACE,
	"]":             KEY_RIGHTBRACE,
	"enter":         KEY_ENTER,
	"ctrl":          KEY_LEFTCTRL,
	"control":       KEY_LEFTCTRL,
	"a":             KEY_A,
	"s":             KEY_S,
	"d":             KEY_D,
	"f":             KEY_F,
	"g":             KEY_G,
	"h":             KEY_H,
	"j":             KEY_J,
	"k":             KEY_K,
	"l":             KEY_L,
	";":             KEY_SEMICOLON,
	"'":             KEY_APOSTROPHE,
	"`":             KEY_GRAVE,
	"shift":         KEY_LEFTSHIFT,
	"\\":            KEY_BACKSLASH,
	"z":             KEY_Z,
	"x":             KEY_X,
	"c":             KEY_C,
	"v":             KEY_V,
	"b":             KEY_B,
	"n":             KEY_N,
	"m":             KEY_M,
	",":             KEY_COMMA,
	".":             KEY_DOT,
	"/":             KEY_SLASH,
	"*":             KEY_KPSTAR,
	"alt":           KEY_LEFTALT,
	"space":         KEY_SPACE,
	"capslock":      KEY_CAPSLOCK,
	"f1":            KEY_F1,
	"f2":            KEY_F2,
	"f3":            KEY_F3,
	"f4":            KEY_F4,
	"f5":            KEY_F5,
	"f6":            KEY_F6,
	"f7":            KEY_F7,
	"f8":            KEY_F8,
	"f9":            KEY_F9,
	"f10":           KEY_F10,
	"f11":           KEY_F11,
	"f12":           KEY_F12,
	"numlock":       KEY_NUMLOCK,
	"scrolllock":    KEY_SCROLLLOCK,
	"home":          KEY_HOME,
	"up":            KEY_UP,
	"pageup":        KEY_PAGEUP,
	"left":          KEY_LEFT,
	"right":         KEY_RIGHT,
	"end":           KEY_END,
	"down":          KEY_DOWN,
	"pagedown":      KEY_PAGEDOWN,
	"insert":        KEY_INSERT,
	"delete":        KEY_DELETE,
	"win":           KEY_LEFTMETA,
	"meta":          KEY_LEFTMETA,
	"super":         KEY_LEFTMETA,
	"compose":       KEY_COMPOSE,
	"pause":         KEY_PAUSE,
	"printscreen":   KEY_SYSRQ,
	"sysrq":         KEY_SYSRQ,
}

// VirtualKeyboard represents a virtual keyboard device
type VirtualKeyboard struct {
	file *os.File
}

// NewVirtualKeyboard creates a new virtual keyboard
func NewVirtualKeyboard() (*VirtualKeyboard, error) {
	// Try /dev/uinput first, then /dev/input/uinput
	fd, err := os.OpenFile("/dev/uinput", os.O_WRONLY, 0644)
	if err != nil {
		fd, err = os.OpenFile("/dev/input/uinput", os.O_WRONLY, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to open uinput: %w (need root or input group)", err)
		}
	}

	vk := &VirtualKeyboard{file: fd}
	if err := vk.setupDevice(); err != nil {
		fd.Close()
		return nil, err
	}

	return vk, nil
}

func (vk *VirtualKeyboard) setupDevice() error {
	// Set up the device
	setup := uinputSetup{}
	copy(setup.name[:], []byte("virtual-keyboard"))

	// Ioctl to set up device
	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		vk.file.Fd(),
		UI_DEV_SETUP,
		uintptr(unsafe.Pointer(&setup)),
	)
	if errno != 0 {
		return fmt.Errorf("ioctl UI_DEV_SETUP failed: %v", errno)
	}

	// Enable key events
	if _, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		vk.file.Fd(),
		UI_SET_EVBIT,
		uintptr(EV_KEY),
	); errno != 0 {
		return fmt.Errorf("ioctl UI_SET_EVBIT failed: %v", errno)
	}

	// Create the device
	if _, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		vk.file.Fd(),
		UI_DEV_CREATE,
		0,
	); errno != 0 {
		return fmt.Errorf("ioctl UI_DEV_CREATE failed: %v", errno)
	}

	return nil
}

// SendKey sends a single key press and release
func (vk *VirtualKeyboard) SendKey(code int) error {
	if err := vk.sendEvent(code, 1); err != nil {
		return err
	}
	return vk.sendEvent(code, 0)
}

// SendKeyWithModifiers sends a key with modifier keys held
func (vk *VirtualKeyboard) SendKeyWithModifiers(modifiers []int, code int) error {
	// Press modifiers
	for _, mod := range modifiers {
		if err := vk.sendEvent(mod, 1); err != nil {
			return err
		}
	}

	// Press and release the main key
	if err := vk.sendEvent(code, 1); err != nil {
		return err
	}
	if err := vk.sendEvent(code, 0); err != nil {
		return err
	}

	// Release modifiers in reverse order
	for i := len(modifiers) - 1; i >= 0; i-- {
		if err := vk.sendEvent(modifiers[i], 0); err != nil {
			return err
		}
	}

	return nil
}

func (vk *VirtualKeyboard) sendEvent(code int, value int) error {
	event := inputEvent{
		time:  syscall.Timeval{},
		type_: EV_KEY,
		code:  uint16(code),
		value: int32(value),
	}

	_, err := vk.file.Write((*[24]byte)(unsafe.Pointer(&event))[:])
	return err
}

// Close closes the virtual keyboard device
func (vk *VirtualKeyboard) Close() error {
	if vk.file != nil {
		return vk.file.Close()
	}
	return nil
}

type inputEvent struct {
	time  syscall.Timeval
	type_ uint16
	code  uint16
	value int32
}

// Ioctl constants
const (
	UI_DEV_SETUP = 0x40045539
	UI_DEV_CREATE = 0x4004553a
	UI_SET_EVBIT = 0x40045564
	UI_SET_KEYBIT = 0x40045565
)

// ParseKeyCombo parses a key combination string like "ctrl+o" or "alt+tab"
func ParseKeyCombo(combo string) (modifiers []int, keyCode int, err error) {
	parts := strings.Split(strings.ToLower(combo), "+")
	
	// Last part is the main key
	mainKey := parts[len(parts)-1]
	
	// Check for dangerous combinations
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "ctrl" || part == "control" {
			modifiers = append(modifiers, KEY_LEFTCTRL)
		} else if part == "alt" {
			modifiers = append(modifiers, KEY_LEFTALT)
		} else if part == "shift" {
			modifiers = append(modifiers, KEY_LEFTSHIFT)
		} else if part == "win" || part == "meta" || part == "super" {
			modifiers = append(modifiers, KEY_LEFTMETA)
		}
	}

	// Check for Ctrl+Alt+Del - BLOCK THIS!
	hasCtrl := false
	hasAlt := false
	hasDel := false
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "ctrl" || part == "control" {
			hasCtrl = true
		} else if part == "alt" {
			hasAlt = true
		} else if part == "del" || part == "delete" {
			hasDel = true
		}
	}
	if hasCtrl && hasAlt && hasDel {
		return nil, 0, fmt.Errorf("SECURITY BLOCKED: Ctrl+Alt+Del combination is not allowed")
	}

	// Only these keys are actual modifiers - don't treat as error if used alone
	modifierKeys := map[string]bool{
		"ctrl": true, "control": true,
		"alt": true,
		"shift": true,
		"win": true, "meta": true, "super": true,
	}

	// Check if main key is a modifier (error only for true modifiers)
	if modifierKeys[mainKey] && len(parts) == 1 {
		return nil, 0, fmt.Errorf("'%s' is a modifier key, use like 'ctrl+enter'", mainKey)
	}

	// Get main key code
	code, exists := keyCodeMap[mainKey]
	if !exists {
		// Try single character
		if len(mainKey) == 1 {
			code = getCharKeyCode(mainKey[0])
			if code == KEY_RESERVED {
				return nil, 0, fmt.Errorf("unknown key: '%s'", mainKey)
			}
		} else {
			return nil, 0, fmt.Errorf("unknown key: '%s'", mainKey)
		}
	}

	return modifiers, code, nil
}

func getCharKeyCode(c byte) int {
	// Check if it's an uppercase letter
	if c >= 'A' && c <= 'Z' {
		return KEY_A + int(c-'A')
	}
	// Check if it's a lowercase letter
	if c >= 'a' && c <= 'z' {
		return KEY_A + int(c-'a')
	}
	// Check if it's a number
	if c >= '0' && c <= '9' {
		return KEY_1 + int(c-'1') + 1 // 1 is KEY_1, so 0 needs special handling
	}
	if c == '0' {
		return KEY_0
	}

	// Special characters
	switch c {
	case ' ': return KEY_SPACE
	case '-': return KEY_MINUS
	case '=': return KEY_EQUAL
	case '[': return KEY_LEFTBRACE
	case ']': return KEY_RIGHTBRACE
	case '\\': return KEY_BACKSLASH
	case ';': return KEY_SEMICOLON
	case '\'': return KEY_APOSTROPHE
	case '`': return KEY_GRAVE
	case ',': return KEY_COMMA
	case '.': return KEY_DOT
	case '/': return KEY_SLASH
	}

	return KEY_RESERVED
}

// TypeString types a string of characters
func (vk *VirtualKeyboard) TypeString(text string) error {
	for _, c := range text {
		code := getCharKeyCode(byte(c))
		if code == KEY_RESERVED {
			// Skip unknown characters
			continue
		}
		
		// Check if shift is needed for this character
		needsShift := needsShift(c)
		
		if needsShift {
			if err := vk.SendKey(KEY_LEFTSHIFT); err != nil {
				return err
			}
		}
		
		if err := vk.SendKey(code); err != nil {
			return err
		}
		
		if needsShift {
			if err := vk.SendKey(KEY_LEFTSHIFT); err != nil {
				return err
			}
		}
	}
	return nil
}

func needsShift(c rune) bool {
	return (c >= 'A' && c <= 'Z') ||
		c == '!' || c == '@' || c == '#' || c == '$' || c == '%' ||
		c == '^' || c == '&' || c == '*' || c == '(' || c == ')' ||
		c == '_' || c == '+' || c == '{' || c == '}' || c == '|' ||
		c == ':' || c == '"' || c == '~' || c == '<' || c == '>' ||
		c == '?' || c == '`'
}

// String returns a string representation of the key code (for debugging)
func KeyCodeToString(code int) string {
	for name, c := range keyCodeMap {
		if c == code {
			return name
		}
	}
	return strconv.Itoa(code)
}
