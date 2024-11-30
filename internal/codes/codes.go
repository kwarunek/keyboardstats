package codes

const PRESSED = 1

const EV_SYN = 0x00
const EV_KEY = 0x01
const EV_REL = 0x02
const EV_ABS = 0x03
const EV_MSC = 0x04
const EV_SW = 0x05
const EV_LED = 0x11
const EV_SND = 0x12
const EV_REP = 0x14
const EV_FF = 0x15
const EV_PWR = 0x16
const EV_FF_STATUS = 0x17
const EV_MAX = 0x1f
const EV_CNT = (EV_MAX + 1)
const KEY_RESERVED = 0
const KEY_ESC = 1
const KEY_1 = 2
const KEY_2 = 3
const KEY_3 = 4
const KEY_4 = 5
const KEY_5 = 6
const KEY_6 = 7
const KEY_7 = 8
const KEY_8 = 9
const KEY_9 = 10
const KEY_0 = 11
const KEY_MINUS = 12
const KEY_EQUAL = 13
const KEY_BACKSPACE = 14
const KEY_TAB = 15
const KEY_Q = 16
const KEY_W = 17
const KEY_E = 18
const KEY_R = 19
const KEY_T = 20
const KEY_Y = 21
const KEY_U = 22
const KEY_I = 23
const KEY_O = 24
const KEY_P = 25
const KEY_LEFTBRACE = 26
const KEY_RIGHTBRACE = 27
const KEY_ENTER = 28
const KEY_LEFTCTRL = 29
const KEY_A = 30
const KEY_S = 31
const KEY_D = 32
const KEY_F = 33
const KEY_G = 34
const KEY_H = 35
const KEY_J = 36
const KEY_K = 37
const KEY_L = 38
const KEY_SEMICOLON = 39
const KEY_APOSTROPHE = 40
const KEY_GRAVE = 41
const KEY_LEFTSHIFT = 42
const KEY_BACKSLASH = 43
const KEY_Z = 44
const KEY_X = 45
const KEY_C = 46
const KEY_V = 47
const KEY_B = 48
const KEY_N = 49
const KEY_M = 50
const KEY_COMMA = 51
const KEY_DOT = 52
const KEY_SLASH = 53
const KEY_RIGHTSHIFT = 54
const KEY_KPASTERISK = 55
const KEY_LEFTALT = 56
const KEY_SPACE = 57
const KEY_CAPSLOCK = 58
const KEY_F1 = 59
const KEY_F2 = 60
const KEY_F3 = 61
const KEY_F4 = 62
const KEY_F5 = 63
const KEY_F6 = 64
const KEY_F7 = 65
const KEY_F8 = 66
const KEY_F9 = 67
const KEY_F10 = 68
const KEY_NUMLOCK = 69
const KEY_SCROLLLOCK = 70
const KEY_KP7 = 71
const KEY_KP8 = 72
const KEY_KP9 = 73
const KEY_KPMINUS = 74
const KEY_KP4 = 75
const KEY_KP5 = 76
const KEY_KP6 = 77
const KEY_KPPLUS = 78
const KEY_KP1 = 79
const KEY_KP2 = 80
const KEY_KP3 = 81
const KEY_KP0 = 82
const KEY_KPDOT = 83
const KEY_ZENKAKUHANKAKU = 85
const KEY_102ND = 86
const KEY_F11 = 87
const KEY_F12 = 88
const KEY_RO = 89
const KEY_KATAKANA = 90
const KEY_HIRAGANA = 91
const KEY_HENKAN = 92
const KEY_KATAKANAHIRAGANA = 93
const KEY_MUHENKAN = 94
const KEY_KPJPCOMMA = 95
const KEY_KPENTER = 96
const KEY_RIGHTCTRL = 97
const KEY_KPSLASH = 98
const KEY_SYSRQ = 99
const KEY_RIGHTALT = 100
const KEY_LINEFEED = 101
const KEY_HOME = 102
const KEY_UP = 103
const KEY_PAGEUP = 104
const KEY_LEFT = 105
const KEY_RIGHT = 106
const KEY_END = 107
const KEY_DOWN = 108
const KEY_PAGEDOWN = 109
const KEY_INSERT = 110
const KEY_DELETE = 111
const KEY_MACRO = 112
const KEY_MUTE = 113
const KEY_VOLUMEDOWN = 114
const KEY_VOLUMEUP = 115
const KEY_POWER = 116
const KEY_KPEQUAL = 117
const KEY_KPPLUSMINUS = 118
const KEY_PAUSE = 119
const KEY_SCALE = 120
const KEY_KPCOMMA = 121
const KEY_HANGEUL = 122
const KEY_HANJA = 123
const KEY_YEN = 124
const KEY_LEFTMETA = 125
const KEY_RIGHTMETA = 126
const KEY_COMPOSE = 127
const KEY_STOP = 128
const KEY_AGAIN = 129
const KEY_PROPS = 130
const KEY_UNDO = 131
const KEY_FRONT = 132
const KEY_COPY = 133
const KEY_OPEN = 134
const KEY_PASTE = 135
const KEY_FIND = 136
const KEY_CUT = 137
const KEY_HELP = 138
const KEY_MENU = 139
const KEY_CALC = 140
const KEY_SETUP = 141
const KEY_SLEEP = 142
const KEY_WAKEUP = 143
const KEY_FILE = 144
const KEY_SENDFILE = 145
const KEY_DELETEFILE = 146
const KEY_XFER = 147
const KEY_PROG1 = 148
const KEY_PROG2 = 149
const KEY_WWW = 150
const KEY_MSDOS = 151
const KEY_COFFEE = 152
const KEY_ROTATE_DISPLAY = 153
const KEY_CYCLEWINDOWS = 154
const KEY_MAIL = 155
const KEY_BOOKMARKS = 156
const KEY_COMPUTER = 157
const KEY_BACK = 158
const KEY_FORWARD = 159
const KEY_CLOSECD = 160
const KEY_EJECTCD = 161
const KEY_EJECTCLOSECD = 162
const KEY_NEXTSONG = 163
const KEY_PLAYPAUSE = 164
const KEY_PREVIOUSSONG = 165
const KEY_STOPCD = 166
const KEY_RECORD = 167
const KEY_REWIND = 168
const KEY_PHONE = 169
const KEY_ISO = 170
const KEY_CONFIG = 171
const KEY_HOMEPAGE = 172
const KEY_REFRESH = 173
const KEY_EXIT = 174
const KEY_MOVE = 175
const KEY_EDIT = 176
const KEY_SCROLLUP = 177
const KEY_SCROLLDOWN = 178
const KEY_KPLEFTPAREN = 179
const KEY_KPRIGHTPAREN = 180
const KEY_NEW = 181
const KEY_REDO = 182
const KEY_F13 = 183
const KEY_F14 = 184
const KEY_F15 = 185
const KEY_F16 = 186
const KEY_F17 = 187
const KEY_F18 = 188
const KEY_F19 = 189
const KEY_F20 = 190
const KEY_F21 = 191
const KEY_F22 = 192
const KEY_F23 = 193
const KEY_F24 = 194
const KEY_PLAYCD = 200
const KEY_PAUSECD = 201
const KEY_PROG3 = 202
const KEY_PROG4 = 203
const KEY_ALL_APPLICATIONS = 204
const KEY_SUSPEND = 205
const KEY_CLOSE = 206
const KEY_PLAY = 207
const KEY_FASTFORWARD = 208
const KEY_BASSBOOST = 209
const KEY_PRINT = 210
const KEY_HP = 211
const KEY_CAMERA = 212
const KEY_SOUND = 213
const KEY_QUESTION = 214
const KEY_EMAIL = 215
const KEY_CHAT = 216
const KEY_SEARCH = 217
const KEY_CONNECT = 218
const KEY_FINANCE = 219
const KEY_SPORT = 220
const KEY_SHOP = 221
const KEY_ALTERASE = 222
const KEY_CANCEL = 223
const KEY_BRIGHTNESSDOWN = 224
const KEY_BRIGHTNESSUP = 225
const KEY_MEDIA = 226
const KEY_KBDILLUMTOGGLE = 228
const KEY_KBDILLUMDOWN = 229
const KEY_KBDILLUMUP = 230
const KEY_SEND = 231
const KEY_REPLY = 232
const KEY_FORWARDMAIL = 233
const KEY_SAVE = 234
const KEY_DOCUMENTS = 235
const KEY_BATTERY = 236
const KEY_BLUETOOTH = 237
const KEY_WLAN = 238
const KEY_UWB = 239
const KEY_UNKNOWN = 240
const KEY_VIDEO_NEXT = 241
const KEY_VIDEO_PREV = 242
const KEY_BRIGHTNESS_CYCLE = 243
const KEY_DISPLAY_OFF = 245
const KEY_WWAN = 246
const KEY_RFKILL = 247
const KEY_MICMUTE = 248
const KEY_OK = 0x160
const KEY_SELECT = 0x161
const KEY_GOTO = 0x162
const KEY_CLEAR = 0x163
const KEY_POWER2 = 0x164
const KEY_OPTION = 0x165
const KEY_INFO = 0x166
const KEY_TIME = 0x167
const KEY_VENDOR = 0x168
const KEY_ARCHIVE = 0x169
const KEY_PROGRAM = 0x16a
const KEY_CHANNEL = 0x16b
const KEY_FAVORITES = 0x16c
const KEY_EPG = 0x16d
const KEY_PVR = 0x16e
const KEY_MHP = 0x16f
const KEY_LANGUAGE = 0x170
const KEY_TITLE = 0x171
const KEY_SUBTITLE = 0x172
const KEY_ANGLE = 0x173
const KEY_FULL_SCREEN = 0x174
const KEY_MODE = 0x175
const KEY_KEYBOARD = 0x176
const KEY_ASPECT_RATIO = 0x177
const KEY_PC = 0x178
const KEY_TV = 0x179
const KEY_TV2 = 0x17a
const KEY_VCR = 0x17b
const KEY_VCR2 = 0x17c
const KEY_SAT = 0x17d
const KEY_SAT2 = 0x17e
const KEY_CD = 0x17f
const KEY_TAPE = 0x180
const KEY_RADIO = 0x181
const KEY_TUNER = 0x182
const KEY_PLAYER = 0x183
const KEY_TEXT = 0x184
const KEY_DVD = 0x185
const KEY_AUX = 0x186
const KEY_MP3 = 0x187
const KEY_AUDIO = 0x188
const KEY_VIDEO = 0x189
const KEY_DIRECTORY = 0x18a
const KEY_LIST = 0x18b
const KEY_MEMO = 0x18c
const KEY_CALENDAR = 0x18d
const KEY_RED = 0x18e
const KEY_GREEN = 0x18f
const KEY_YELLOW = 0x190
const KEY_BLUE = 0x191
const KEY_CHANNELUP = 0x192
const KEY_CHANNELDOWN = 0x193
const KEY_FIRST = 0x194
const KEY_LAST = 0x195
const KEY_AB = 0x196
const KEY_NEXT = 0x197
const KEY_RESTART = 0x198
const KEY_SLOW = 0x199
const KEY_SHUFFLE = 0x19a
const KEY_BREAK = 0x19b
const KEY_PREVIOUS = 0x19c
const KEY_DIGITS = 0x19d
const KEY_TEEN = 0x19e
const KEY_TWEN = 0x19f
const KEY_VIDEOPHONE = 0x1a0
const KEY_GAMES = 0x1a1
const KEY_ZOOMIN = 0x1a2
const KEY_ZOOMOUT = 0x1a3
const KEY_ZOOMRESET = 0x1a4
const KEY_WORDPROCESSOR = 0x1a5
const KEY_EDITOR = 0x1a6
const KEY_SPREADSHEET = 0x1a7
const KEY_GRAPHICSEDITOR = 0x1a8
const KEY_PRESENTATION = 0x1a9
const KEY_DATABASE = 0x1aa
const KEY_NEWS = 0x1ab
const KEY_VOICEMAIL = 0x1ac
const KEY_ADDRESSBOOK = 0x1ad
const KEY_MESSENGER = 0x1ae
const KEY_DISPLAYTOGGLE = 0x1af
const KEY_SPELLCHECK = 0x1b0
const KEY_LOGOFF = 0x1b1
const KEY_DOLLAR = 0x1b2
const KEY_EURO = 0x1b3
const KEY_FRAMEBACK = 0x1b4
const KEY_FRAMEFORWARD = 0x1b5
const KEY_CONTEXT_MENU = 0x1b6
const KEY_MEDIA_REPEAT = 0x1b7
const KEY_10CHANNELSUP = 0x1b8
const KEY_10CHANNELSDOWN = 0x1b9
const KEY_IMAGES = 0x1ba
const KEY_NOTIFICATION_CENTER = 0x1bc
const KEY_PICKUP_PHONE = 0x1bd
const KEY_HANGUP_PHONE = 0x1be
const KEY_DEL_EOL = 0x1c0
const KEY_DEL_EOS = 0x1c1
const KEY_INS_LINE = 0x1c2
const KEY_DEL_LINE = 0x1c3
const KEY_FN = 0x1d0
const KEY_FN_ESC = 0x1d1
const KEY_FN_F1 = 0x1d2
const KEY_FN_F2 = 0x1d3
const KEY_FN_F3 = 0x1d4
const KEY_FN_F4 = 0x1d5
const KEY_FN_F5 = 0x1d6
const KEY_FN_F6 = 0x1d7
const KEY_FN_F7 = 0x1d8
const KEY_FN_F8 = 0x1d9
const KEY_FN_F9 = 0x1da
const KEY_FN_F10 = 0x1db
const KEY_FN_F11 = 0x1dc
const KEY_FN_F12 = 0x1dd
const KEY_FN_1 = 0x1de
const KEY_FN_2 = 0x1df
const KEY_FN_D = 0x1e0
const KEY_FN_E = 0x1e1
const KEY_FN_F = 0x1e2
const KEY_FN_S = 0x1e3
const KEY_FN_B = 0x1e4
const KEY_FN_RIGHT_SHIFT = 0x1e5
const KEY_BRL_DOT1 = 0x1f1
const KEY_BRL_DOT2 = 0x1f2
const KEY_BRL_DOT3 = 0x1f3
const KEY_BRL_DOT4 = 0x1f4
const KEY_BRL_DOT5 = 0x1f5
const KEY_BRL_DOT6 = 0x1f6
const KEY_BRL_DOT7 = 0x1f7
const KEY_BRL_DOT8 = 0x1f8
const KEY_BRL_DOT9 = 0x1f9
const KEY_BRL_DOT10 = 0x1fa
const KEY_NUMERIC_0 = 0x200
const KEY_NUMERIC_1 = 0x201
const KEY_NUMERIC_2 = 0x202
const KEY_NUMERIC_3 = 0x203
const KEY_NUMERIC_4 = 0x204
const KEY_NUMERIC_5 = 0x205
const KEY_NUMERIC_6 = 0x206
const KEY_NUMERIC_7 = 0x207
const KEY_NUMERIC_8 = 0x208
const KEY_NUMERIC_9 = 0x209
const KEY_NUMERIC_STAR = 0x20a
const KEY_NUMERIC_POUND = 0x20b
const KEY_NUMERIC_A = 0x20c
const KEY_NUMERIC_B = 0x20d
const KEY_NUMERIC_C = 0x20e
const KEY_NUMERIC_D = 0x20f
const KEY_CAMERA_FOCUS = 0x210
const KEY_WPS_BUTTON = 0x211
const KEY_TOUCHPAD_TOGGLE = 0x212
const KEY_TOUCHPAD_ON = 0x213
const KEY_TOUCHPAD_OFF = 0x214
const KEY_CAMERA_ZOOMIN = 0x215
const KEY_CAMERA_ZOOMOUT = 0x216
const KEY_CAMERA_UP = 0x217
const KEY_CAMERA_DOWN = 0x218
const KEY_CAMERA_LEFT = 0x219
const KEY_CAMERA_RIGHT = 0x21a
const KEY_ATTENDANT_ON = 0x21b
const KEY_ATTENDANT_OFF = 0x21c
const KEY_ATTENDANT_TOGGLE = 0x21d
const KEY_LIGHTS_TOGGLE = 0x21e
const KEY_ALS_TOGGLE = 0x230
const KEY_ROTATE_LOCK_TOGGLE = 0x231
const KEY_REFRESH_RATE_TOGGLE = 0x232
const KEY_BUTTONCONFIG = 0x240
const KEY_TASKMANAGER = 0x241
const KEY_JOURNAL = 0x242
const KEY_CONTROLPANEL = 0x243
const KEY_APPSELECT = 0x244
const KEY_SCREENSAVER = 0x245
const KEY_VOICECOMMAND = 0x246
const KEY_ASSISTANT = 0x247
const KEY_KBD_LAYOUT_NEXT = 0x248
const KEY_EMOJI_PICKER = 0x249
const KEY_DICTATE = 0x24a
const KEY_CAMERA_ACCESS_ENABLE = 0x24b
const KEY_CAMERA_ACCESS_DISABLE = 0x24c
const KEY_CAMERA_ACCESS_TOGGLE = 0x24d
const KEY_ACCESSIBILITY = 0x24e
const KEY_DO_NOT_DISTURB = 0x24f
const KEY_BRIGHTNESS_MIN = 0x250
const KEY_BRIGHTNESS_MAX = 0x251
const KEY_KBDINPUTASSIST_PREV = 0x260
const KEY_KBDINPUTASSIST_NEXT = 0x261
const KEY_KBDINPUTASSIST_PREVGROUP = 0x262
const KEY_KBDINPUTASSIST_NEXTGROUP = 0x263
const KEY_KBDINPUTASSIST_ACCEPT = 0x264
const KEY_KBDINPUTASSIST_CANCEL = 0x265
const KEY_RIGHT_UP = 0x266
const KEY_RIGHT_DOWN = 0x267
const KEY_LEFT_UP = 0x268
const KEY_LEFT_DOWN = 0x269
const KEY_ROOT_MENU = 0x26a
const KEY_MEDIA_TOP_MENU = 0x26b
const KEY_NUMERIC_11 = 0x26c
const KEY_NUMERIC_12 = 0x26d
const KEY_AUDIO_DESC = 0x26e
const KEY_3D_MODE = 0x26f
const KEY_NEXT_FAVORITE = 0x270
const KEY_STOP_RECORD = 0x271
const KEY_PAUSE_RECORD = 0x272
const KEY_VOD = 0x273
const KEY_UNMUTE = 0x274
const KEY_FASTREVERSE = 0x275
const KEY_SLOWREVERSE = 0x276
const KEY_DATA = 0x277
const KEY_ONSCREEN_KEYBOARD = 0x278
const KEY_PRIVACY_SCREEN_TOGGLE = 0x279
const KEY_SELECTIVE_SCREENSHOT = 0x27a
const KEY_NEXT_ELEMENT = 0x27b
const KEY_PREVIOUS_ELEMENT = 0x27c
const KEY_AUTOPILOT_ENGAGE_TOGGLE = 0x27d
const KEY_MARK_WAYPOINT = 0x27e
const KEY_SOS = 0x27f
const KEY_NAV_CHART = 0x280
const KEY_FISHING_CHART = 0x281
const KEY_SINGLE_RANGE_RADAR = 0x282
const KEY_DUAL_RANGE_RADAR = 0x283
const KEY_RADAR_OVERLAY = 0x284
const KEY_TRADITIONAL_SONAR = 0x285
const KEY_CLEARVU_SONAR = 0x286
const KEY_SIDEVU_SONAR = 0x287
const KEY_NAV_INFO = 0x288
const KEY_BRIGHTNESS_MENU = 0x289
const KEY_MACRO1 = 0x290
const KEY_MACRO2 = 0x291
const KEY_MACRO3 = 0x292
const KEY_MACRO4 = 0x293
const KEY_MACRO5 = 0x294
const KEY_MACRO6 = 0x295
const KEY_MACRO7 = 0x296
const KEY_MACRO8 = 0x297
const KEY_MACRO9 = 0x298
const KEY_MACRO10 = 0x299
const KEY_MACRO11 = 0x29a
const KEY_MACRO12 = 0x29b
const KEY_MACRO13 = 0x29c
const KEY_MACRO14 = 0x29d
const KEY_MACRO15 = 0x29e
const KEY_MACRO16 = 0x29f
const KEY_MACRO17 = 0x2a0
const KEY_MACRO18 = 0x2a1
const KEY_MACRO19 = 0x2a2
const KEY_MACRO20 = 0x2a3
const KEY_MACRO21 = 0x2a4
const KEY_MACRO22 = 0x2a5
const KEY_MACRO23 = 0x2a6
const KEY_MACRO24 = 0x2a7
const KEY_MACRO25 = 0x2a8
const KEY_MACRO26 = 0x2a9
const KEY_MACRO27 = 0x2aa
const KEY_MACRO28 = 0x2ab
const KEY_MACRO29 = 0x2ac
const KEY_MACRO30 = 0x2ad
const KEY_MACRO_RECORD_START = 0x2b0
const KEY_MACRO_RECORD_STOP = 0x2b1
const KEY_MACRO_PRESET_CYCLE = 0x2b2
const KEY_MACRO_PRESET1 = 0x2b3
const KEY_MACRO_PRESET2 = 0x2b4
const KEY_MACRO_PRESET3 = 0x2b5
const KEY_KBD_LCD_MENU1 = 0x2b8
const KEY_KBD_LCD_MENU2 = 0x2b9
const KEY_KBD_LCD_MENU3 = 0x2ba
const KEY_KBD_LCD_MENU4 = 0x2bb
const KEY_KBD_LCD_MENU5 = 0x2bc
const KEY_MAX = 0x2ff
const KEY_CNT = (KEY_MAX + 1)

var KeyCodeToName = map[int]string{
    0:             "KEY_RESERVED",
    1:             "KEY_ESC",
    2:             "KEY_1",
    3:             "KEY_2",
    4:             "KEY_3",
    5:             "KEY_4",
    6:             "KEY_5",
    7:             "KEY_6",
    8:             "KEY_7",
    9:             "KEY_8",
    10:            "KEY_9",
    11:            "KEY_0",
    12:            "KEY_MINUS",
    13:            "KEY_EQUAL",
    14:            "KEY_BACKSPACE",
    15:            "KEY_TAB",
    16:            "KEY_Q",
    17:            "KEY_W",
    18:            "KEY_E",
    19:            "KEY_R",
    20:            "KEY_T",
    21:            "KEY_Y",
    22:            "KEY_U",
    23:            "KEY_I",
    24:            "KEY_O",
    25:            "KEY_P",
    26:            "KEY_LEFTBRACE",
    27:            "KEY_RIGHTBRACE",
    28:            "KEY_ENTER",
    29:            "KEY_LEFTCTRL",
    30:            "KEY_A",
    31:            "KEY_S",
    32:            "KEY_D",
    33:            "KEY_F",
    34:            "KEY_G",
    35:            "KEY_H",
    36:            "KEY_J",
    37:            "KEY_K",
    38:            "KEY_L",
    39:            "KEY_SEMICOLON",
    40:            "KEY_APOSTROPHE",
    41:            "KEY_GRAVE",
    42:            "KEY_LEFTSHIFT",
    43:            "KEY_BACKSLASH",
    44:            "KEY_Z",
    45:            "KEY_X",
    46:            "KEY_C",
    47:            "KEY_V",
    48:            "KEY_B",
    49:            "KEY_N",
    50:            "KEY_M",
    51:            "KEY_COMMA",
    52:            "KEY_DOT",
    53:            "KEY_SLASH",
    54:            "KEY_RIGHTSHIFT",
    55:            "KEY_KPASTERISK",
    56:            "KEY_LEFTALT",
    57:            "KEY_SPACE",
    58:            "KEY_CAPSLOCK",
    59:            "KEY_F1",
    60:            "KEY_F2",
    61:            "KEY_F3",
    62:            "KEY_F4",
    63:            "KEY_F5",
    64:            "KEY_F6",
    65:            "KEY_F7",
    66:            "KEY_F8",
    67:            "KEY_F9",
    68:            "KEY_F10",
    69:            "KEY_NUMLOCK",
    70:            "KEY_SCROLLLOCK",
    71:            "KEY_KP7",
    72:            "KEY_KP8",
    73:            "KEY_KP9",
    74:            "KEY_KPMINUS",
    75:            "KEY_KP4",
    76:            "KEY_KP5",
    77:            "KEY_KP6",
    78:            "KEY_KPPLUS",
    79:            "KEY_KP1",
    80:            "KEY_KP2",
    81:            "KEY_KP3",
    82:            "KEY_KP0",
    83:            "KEY_KPDOT",
    85:            "KEY_ZENKAKUHANKAKU",
    86:            "KEY_102ND",
    87:            "KEY_F11",
    88:            "KEY_F12",
    89:            "KEY_RO",
    90:            "KEY_KATAKANA",
    91:            "KEY_HIRAGANA",
    92:            "KEY_HENKAN",
    93:            "KEY_KATAKANAHIRAGANA",
    94:            "KEY_MUHENKAN",
    95:            "KEY_KPJPCOMMA",
    96:            "KEY_KPENTER",
    97:            "KEY_RIGHTCTRL",
    98:            "KEY_KPSLASH",
    99:            "KEY_SYSRQ",
    100:           "KEY_RIGHTALT",
    101:           "KEY_LINEFEED",
    102:           "KEY_HOME",
    103:           "KEY_UP",
    104:           "KEY_PAGEUP",
    105:           "KEY_LEFT",
    106:           "KEY_RIGHT",
    107:           "KEY_END",
    108:           "KEY_DOWN",
    109:           "KEY_PAGEDOWN",
    110:           "KEY_INSERT",
    111:           "KEY_DELETE",
    112:           "KEY_MACRO",
    113:           "KEY_MUTE",
    114:           "KEY_VOLUMEDOWN",
    115:           "KEY_VOLUMEUP",
    116:           "KEY_POWER",
    117:           "KEY_KPEQUAL",
    118:           "KEY_KPPLUSMINUS",
    119:           "KEY_PAUSE",
    120:           "KEY_SCALE",
    121:           "KEY_KPCOMMA",
    122:           "KEY_HANGEUL",
    123:           "KEY_HANJA",
    124:           "KEY_YEN",
    125:           "KEY_LEFTMETA",
    126:           "KEY_RIGHTMETA",
    127:           "KEY_COMPOSE",
    128:           "KEY_STOP",
    129:           "KEY_AGAIN",
    130:           "KEY_PROPS",
    131:           "KEY_UNDO",
    132:           "KEY_FRONT",
    133:           "KEY_COPY",
    134:           "KEY_OPEN",
    135:           "KEY_PASTE",
    136:           "KEY_FIND",
    137:           "KEY_CUT",
    138:           "KEY_HELP",
    139:           "KEY_MENU",
    140:           "KEY_CALC",
    141:           "KEY_SETUP",
    142:           "KEY_SLEEP",
    143:           "KEY_WAKEUP",
    144:           "KEY_FILE",
    145:           "KEY_SENDFILE",
    146:           "KEY_DELETEFILE",
    147:           "KEY_XFER",
    148:           "KEY_PROG1",
    149:           "KEY_PROG2",
    150:           "KEY_WWW",
    151:           "KEY_MSDOS",
    152:           "KEY_COFFEE",
    153:           "KEY_ROTATE_DISPLAY",
    154:           "KEY_CYCLEWINDOWS",
    155:           "KEY_MAIL",
    156:           "KEY_BOOKMARKS",
    157:           "KEY_COMPUTER",
    158:           "KEY_BACK",
    159:           "KEY_FORWARD",
    160:           "KEY_CLOSECD",
    161:           "KEY_EJECTCD",
    162:           "KEY_EJECTCLOSECD",
    163:           "KEY_NEXTSONG",
    164:           "KEY_PLAYPAUSE",
    165:           "KEY_PREVIOUSSONG",
    166:           "KEY_STOPCD",
    167:           "KEY_RECORD",
    168:           "KEY_REWIND",
    169:           "KEY_PHONE",
    170:           "KEY_ISO",
    171:           "KEY_CONFIG",
    172:           "KEY_HOMEPAGE",
    173:           "KEY_REFRESH",
    174:           "KEY_EXIT",
    175:           "KEY_MOVE",
    176:           "KEY_EDIT",
    177:           "KEY_SCROLLUP",
    178:           "KEY_SCROLLDOWN",
    179:           "KEY_KPLEFTPAREN",
    180:           "KEY_KPRIGHTPAREN",
    181:           "KEY_NEW",
    182:           "KEY_REDO",
    183:           "KEY_F13",
    184:           "KEY_F14",
    185:           "KEY_F15",
    186:           "KEY_F16",
    187:           "KEY_F17",
    188:           "KEY_F18",
    189:           "KEY_F19",
    190:           "KEY_F20",
    191:           "KEY_F21",
    192:           "KEY_F22",
    193:           "KEY_F23",
    194:           "KEY_F24",
    200:           "KEY_PLAYCD",
    201:           "KEY_PAUSECD",
    202:           "KEY_PROG3",
    203:           "KEY_PROG4",
    204:           "KEY_ALL_APPLICATIONS",
    205:           "KEY_SUSPEND",
    206:           "KEY_CLOSE",
    207:           "KEY_PLAY",
    208:           "KEY_FASTFORWARD",
    209:           "KEY_BASSBOOST",
    210:           "KEY_PRINT",
    211:           "KEY_HP",
    212:           "KEY_CAMERA",
    213:           "KEY_SOUND",
    214:           "KEY_QUESTION",
    215:           "KEY_EMAIL",
    216:           "KEY_CHAT",
    217:           "KEY_SEARCH",
    218:           "KEY_CONNECT",
    219:           "KEY_FINANCE",
    220:           "KEY_SPORT",
    221:           "KEY_SHOP",
    222:           "KEY_ALTERASE",
    223:           "KEY_CANCEL",
    224:           "KEY_BRIGHTNESSDOWN",
    225:           "KEY_BRIGHTNESSUP",
    226:           "KEY_MEDIA",
    228:           "KEY_KBDILLUMTOGGLE",
    229:           "KEY_KBDILLUMDOWN",
    230:           "KEY_KBDILLUMUP",
    231:           "KEY_SEND",
    232:           "KEY_REPLY",
    233:           "KEY_FORWARDMAIL",
    234:           "KEY_SAVE",
    235:           "KEY_DOCUMENTS",
    236:           "KEY_BATTERY",
    237:           "KEY_BLUETOOTH",
    238:           "KEY_WLAN",
    239:           "KEY_UWB",
    240:           "KEY_UNKNOWN",
    241:           "KEY_VIDEO_NEXT",
    242:           "KEY_VIDEO_PREV",
    243:           "KEY_BRIGHTNESS_CYCLE",
    245:           "KEY_DISPLAY_OFF",
    246:           "KEY_WWAN",
    247:           "KEY_RFKILL",
    248:           "KEY_MICMUTE",
    0x160:         "KEY_OK",
    0x161:         "KEY_SELECT",
    0x162:         "KEY_GOTO",
    0x163:         "KEY_CLEAR",
    0x164:         "KEY_POWER2",
    0x165:         "KEY_OPTION",
    0x166:         "KEY_INFO",
    0x167:         "KEY_TIME",
    0x168:         "KEY_VENDOR",
    0x169:         "KEY_ARCHIVE",
    0x16a:         "KEY_PROGRAM",
    0x16b:         "KEY_CHANNEL",
    0x16c:         "KEY_FAVORITES",
    0x16d:         "KEY_EPG",
    0x16e:         "KEY_PVR",
    0x16f:         "KEY_MHP",
    0x170:         "KEY_LANGUAGE",
    0x171:         "KEY_TITLE",
    0x172:         "KEY_SUBTITLE",
    0x173:         "KEY_ANGLE",
    0x174:         "KEY_FULL_SCREEN",
    0x175:         "KEY_MODE",
    0x176:         "KEY_KEYBOARD",
    0x177:         "KEY_ASPECT_RATIO",
    0x178:         "KEY_PC",
    0x179:         "KEY_TV",
    0x17a:         "KEY_TV2",
    0x17b:         "KEY_VCR",
    0x17c:         "KEY_VCR2",
    0x17d:         "KEY_SAT",
    0x17e:         "KEY_SAT2",
    0x17f:         "KEY_CD",
    0x180:         "KEY_TAPE",
    0x181:         "KEY_RADIO",
    0x182:         "KEY_TUNER",
    0x183:         "KEY_PLAYER",
    0x184:         "KEY_TEXT",
    0x185:         "KEY_DVD",
    0x186:         "KEY_AUX",
    0x187:         "KEY_MP3",
    0x188:         "KEY_AUDIO",
    0x189:         "KEY_VIDEO",
    0x18a:         "KEY_DIRECTORY",
    0x18b:         "KEY_LIST",
    0x18c:         "KEY_MEMO",
    0x18d:         "KEY_CALENDAR",
    0x18e:         "KEY_RED",
    0x18f:         "KEY_GREEN",
    0x190:         "KEY_YELLOW",
    0x191:         "KEY_BLUE",
    0x192:         "KEY_CHANNELUP",
    0x193:         "KEY_CHANNELDOWN",
    0x194:         "KEY_FIRST",
    0x195:         "KEY_LAST",
    0x196:         "KEY_AB",
    0x197:         "KEY_NEXT",
    0x198:         "KEY_RESTART",
    0x199:         "KEY_SLOW",
    0x19a:         "KEY_SHUFFLE",
    0x19b:         "KEY_BREAK",
    0x19c:         "KEY_PREVIOUS",
    0x19d:         "KEY_DIGITS",
    0x19e:         "KEY_TEEN",
    0x19f:         "KEY_TWEN",
    0x1a0:         "KEY_VIDEOPHONE",
    0x1a1:         "KEY_GAMES",
    0x1a2:         "KEY_ZOOMIN",
    0x1a3:         "KEY_ZOOMOUT",
    0x1a4:         "KEY_ZOOMRESET",
    0x1a5:         "KEY_WORDPROCESSOR",
    0x1a6:         "KEY_EDITOR",
    0x1a7:         "KEY_SPREADSHEET",
    0x1a8:         "KEY_GRAPHICSEDITOR",
    0x1a9:         "KEY_PRESENTATION",
    0x1aa:         "KEY_DATABASE",
    0x1ab:         "KEY_NEWS",
    0x1ac:         "KEY_VOICEMAIL",
    0x1ad:         "KEY_ADDRESSBOOK",
    0x1ae:         "KEY_MESSENGER",
    0x1af:         "KEY_DISPLAYTOGGLE",
    0x1b0:         "KEY_SPELLCHECK",
    0x1b1:         "KEY_LOGOFF",
    0x1b2:         "KEY_DOLLAR",
    0x1b3:         "KEY_EURO",
    0x1b4:         "KEY_FRAMEBACK",
    0x1b5:         "KEY_FRAMEFORWARD",
    0x1b6:         "KEY_CONTEXT_MENU",
    0x1b7:         "KEY_MEDIA_REPEAT",
    0x1b8:         "KEY_10CHANNELSUP",
    0x1b9:         "KEY_10CHANNELSDOWN",
    0x1ba:         "KEY_IMAGES",
    0x1bc:         "KEY_NOTIFICATION_CENTER",
    0x1bd:         "KEY_PICKUP_PHONE",
    0x1be:         "KEY_HANGUP_PHONE",
    0x1c0:         "KEY_DEL_EOL",
    0x1c1:         "KEY_DEL_EOS",
    0x1c2:         "KEY_INS_LINE",
    0x1c3:         "KEY_DEL_LINE",
    0x1d0:         "KEY_FN",
    0x1d1:         "KEY_FN_ESC",
    0x1d2:         "KEY_FN_F1",
    0x1d3:         "KEY_FN_F2",
    0x1d4:         "KEY_FN_F3",
    0x1d5:         "KEY_FN_F4",
    0x1d6:         "KEY_FN_F5",
    0x1d7:         "KEY_FN_F6",
    0x1d8:         "KEY_FN_F7",
    0x1d9:         "KEY_FN_F8",
    0x1da:         "KEY_FN_F9",
    0x1db:         "KEY_FN_F10",
    0x1dc:         "KEY_FN_F11",
    0x1dd:         "KEY_FN_F12",
    0x1de:         "KEY_FN_1",
    0x1df:         "KEY_FN_2",
    0x1e0:         "KEY_FN_D",
    0x1e1:         "KEY_FN_E",
    0x1e2:         "KEY_FN_F",
    0x1e3:         "KEY_FN_S",
    0x1e4:         "KEY_FN_B",
    0x1e5:         "KEY_FN_RIGHT_SHIFT",
    0x1f1:         "KEY_BRL_DOT1",
    0x1f2:         "KEY_BRL_DOT2",
    0x1f3:         "KEY_BRL_DOT3",
    0x1f4:         "KEY_BRL_DOT4",
    0x1f5:         "KEY_BRL_DOT5",
    0x1f6:         "KEY_BRL_DOT6",
    0x1f7:         "KEY_BRL_DOT7",
    0x1f8:         "KEY_BRL_DOT8",
    0x1f9:         "KEY_BRL_DOT9",
    0x1fa:         "KEY_BRL_DOT10",
    0x200:         "KEY_NUMERIC_0",
    0x201:         "KEY_NUMERIC_1",
    0x202:         "KEY_NUMERIC_2",
    0x203:         "KEY_NUMERIC_3",
    0x204:         "KEY_NUMERIC_4",
    0x205:         "KEY_NUMERIC_5",
    0x206:         "KEY_NUMERIC_6",
    0x207:         "KEY_NUMERIC_7",
    0x208:         "KEY_NUMERIC_8",
    0x209:         "KEY_NUMERIC_9",
    0x20a:         "KEY_NUMERIC_STAR",
    0x20b:         "KEY_NUMERIC_POUND",
    0x20c:         "KEY_NUMERIC_A",
    0x20d:         "KEY_NUMERIC_B",
    0x20e:         "KEY_NUMERIC_C",
    0x20f:         "KEY_NUMERIC_D",
    0x210:         "KEY_CAMERA_FOCUS",
    0x211:         "KEY_WPS_BUTTON",
    0x212:         "KEY_TOUCHPAD_TOGGLE",
    0x213:         "KEY_TOUCHPAD_ON",
    0x214:         "KEY_TOUCHPAD_OFF",
    0x215:         "KEY_CAMERA_ZOOMIN",
    0x216:         "KEY_CAMERA_ZOOMOUT",
    0x217:         "KEY_CAMERA_UP",
    0x218:         "KEY_CAMERA_DOWN",
    0x219:         "KEY_CAMERA_LEFT",
    0x21a:         "KEY_CAMERA_RIGHT",
    0x21b:         "KEY_ATTENDANT_ON",
    0x21c:         "KEY_ATTENDANT_OFF",
    0x21d:         "KEY_ATTENDANT_TOGGLE",
    0x21e:         "KEY_LIGHTS_TOGGLE",
    0x230:         "KEY_ALS_TOGGLE",
    0x231:         "KEY_ROTATE_LOCK_TOGGLE",
    0x232:         "KEY_REFRESH_RATE_TOGGLE",
    0x240:         "KEY_BUTTONCONFIG",
    0x241:         "KEY_TASKMANAGER",
    0x242:         "KEY_JOURNAL",
    0x243:         "KEY_CONTROLPANEL",
    0x244:         "KEY_APPSELECT",
    0x245:         "KEY_SCREENSAVER",
    0x246:         "KEY_VOICECOMMAND",
    0x247:         "KEY_ASSISTANT",
    0x248:         "KEY_KBD_LAYOUT_NEXT",
    0x249:         "KEY_EMOJI_PICKER",
    0x24a:         "KEY_DICTATE",
    0x24b:         "KEY_CAMERA_ACCESS_ENABLE",
    0x24c:         "KEY_CAMERA_ACCESS_DISABLE",
    0x24d:         "KEY_CAMERA_ACCESS_TOGGLE",
    0x24e:         "KEY_ACCESSIBILITY",
    0x24f:         "KEY_DO_NOT_DISTURB",
    0x250:         "KEY_BRIGHTNESS_MIN",
    0x251:         "KEY_BRIGHTNESS_MAX",
    0x260:         "KEY_KBDINPUTASSIST_PREV",
    0x261:         "KEY_KBDINPUTASSIST_NEXT",
    0x262:         "KEY_KBDINPUTASSIST_PREVGROUP",
    0x263:         "KEY_KBDINPUTASSIST_NEXTGROUP",
    0x264:         "KEY_KBDINPUTASSIST_ACCEPT",
    0x265:         "KEY_KBDINPUTASSIST_CANCEL",
    0x266:         "KEY_RIGHT_UP",
    0x267:         "KEY_RIGHT_DOWN",
    0x268:         "KEY_LEFT_UP",
    0x269:         "KEY_LEFT_DOWN",
    0x26a:         "KEY_ROOT_MENU",
    0x26b:         "KEY_MEDIA_TOP_MENU",
    0x26c:         "KEY_NUMERIC_11",
    0x26d:         "KEY_NUMERIC_12",
    0x26e:         "KEY_AUDIO_DESC",
    0x26f:         "KEY_3D_MODE",
    0x270:         "KEY_NEXT_FAVORITE",
    0x271:         "KEY_STOP_RECORD",
    0x272:         "KEY_PAUSE_RECORD",
    0x273:         "KEY_VOD",
    0x274:         "KEY_UNMUTE",
    0x275:         "KEY_FASTREVERSE",
    0x276:         "KEY_SLOWREVERSE",
    0x277:         "KEY_DATA",
    0x278:         "KEY_ONSCREEN_KEYBOARD",
    0x279:         "KEY_PRIVACY_SCREEN_TOGGLE",
    0x27a:         "KEY_SELECTIVE_SCREENSHOT",
    0x27b:         "KEY_NEXT_ELEMENT",
    0x27c:         "KEY_PREVIOUS_ELEMENT",
    0x27d:         "KEY_AUTOPILOT_ENGAGE_TOGGLE",
    0x27e:         "KEY_MARK_WAYPOINT",
    0x27f:         "KEY_SOS",
    0x280:         "KEY_NAV_CHART",
    0x281:         "KEY_FISHING_CHART",
    0x282:         "KEY_SINGLE_RANGE_RADAR",
    0x283:         "KEY_DUAL_RANGE_RADAR",
    0x284:         "KEY_RADAR_OVERLAY",
    0x285:         "KEY_TRADITIONAL_SONAR",
    0x286:         "KEY_CLEARVU_SONAR",
    0x287:         "KEY_SIDEVU_SONAR",
    0x288:         "KEY_NAV_INFO",
    0x289:         "KEY_BRIGHTNESS_MENU",
    0x290:         "KEY_MACRO1",
    0x291:         "KEY_MACRO2",
    0x292:         "KEY_MACRO3",
    0x293:         "KEY_MACRO4",
    0x294:         "KEY_MACRO5",
    0x295:         "KEY_MACRO6",
    0x296:         "KEY_MACRO7",
    0x297:         "KEY_MACRO8",
    0x298:         "KEY_MACRO9",
    0x299:         "KEY_MACRO10",
    0x29a:         "KEY_MACRO11",
    0x29b:         "KEY_MACRO12",
    0x29c:         "KEY_MACRO13",
    0x29d:         "KEY_MACRO14",
    0x29e:         "KEY_MACRO15",
    0x29f:         "KEY_MACRO16",
    0x2a0:         "KEY_MACRO17",
    0x2a1:         "KEY_MACRO18",
    0x2a2:         "KEY_MACRO19",
    0x2a3:         "KEY_MACRO20",
    0x2a4:         "KEY_MACRO21",
    0x2a5:         "KEY_MACRO22",
    0x2a6:         "KEY_MACRO23",
    0x2a7:         "KEY_MACRO24",
    0x2a8:         "KEY_MACRO25",
    0x2a9:         "KEY_MACRO26",
    0x2aa:         "KEY_MACRO27",
    0x2ab:         "KEY_MACRO28",
    0x2ac:         "KEY_MACRO29",
    0x2ad:         "KEY_MACRO30",
    0x2b0:         "KEY_MACRO_RECORD_START",
    0x2b1:         "KEY_MACRO_RECORD_STOP",
    0x2b2:         "KEY_MACRO_PRESET_CYCLE",
    0x2b3:         "KEY_MACRO_PRESET1",
    0x2b4:         "KEY_MACRO_PRESET2",
    0x2b5:         "KEY_MACRO_PRESET3",
    0x2b8:         "KEY_KBD_LCD_MENU1",
    0x2b9:         "KEY_KBD_LCD_MENU2",
    0x2ba:         "KEY_KBD_LCD_MENU3",
    0x2bb:         "KEY_KBD_LCD_MENU4",
    0x2bc:         "KEY_KBD_LCD_MENU5",
    0x2ff:         "KEY_MAX",
    (KEY_MAX + 1): "KEY_CNT",
}
