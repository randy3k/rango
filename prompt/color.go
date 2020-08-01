package prompt

import (
	"image/color"
)

type Color int

type ColorDepth int

const (
	ColorDepth1Bit ColorDepth = iota
	ColorDepth4Bits
	ColorDepth8Bits
	ColorDepth24Bits
)

// from https://github.com/tomnomnom/xtermcolor/blob/b78803f00a7e8c5596d5ce64bb717a6a71464792/colors.go

var ColorPalette4Bits = color.Palette{
	// Basic block
	0:  color.RGBA{0x00, 0x00, 0x00, 0xff},
	1:  color.RGBA{0x80, 0x00, 0x00, 0xff},
	2:  color.RGBA{0x00, 0x80, 0x00, 0xff},
	3:  color.RGBA{0x80, 0x80, 0x00, 0xff},
	4:  color.RGBA{0x00, 0x00, 0x80, 0xff},
	5:  color.RGBA{0x80, 0x00, 0x80, 0xff},
	6:  color.RGBA{0x00, 0x80, 0x80, 0xff},
	7:  color.RGBA{0xc0, 0xc0, 0xc0, 0xff},
	8:  color.RGBA{0x80, 0x80, 0x80, 0xff},
	9:  color.RGBA{0xff, 0x00, 0x00, 0xff},
	10: color.RGBA{0x00, 0xff, 0x00, 0xff},
	11: color.RGBA{0xff, 0xff, 0x00, 0xff},
	12: color.RGBA{0x00, 0x00, 0xff, 0xff},
	13: color.RGBA{0xff, 0x00, 0xff, 0xff},
	14: color.RGBA{0x00, 0xff, 0xff, 0xff},
	15: color.RGBA{0xff, 0xff, 0xff, 0xff},
}

var ColorPalette8Bits = color.Palette{
	// Basic block
	0:  color.RGBA{0x00, 0x00, 0x00, 0xff},
	1:  color.RGBA{0x80, 0x00, 0x00, 0xff},
	2:  color.RGBA{0x00, 0x80, 0x00, 0xff},
	3:  color.RGBA{0x80, 0x80, 0x00, 0xff},
	4:  color.RGBA{0x00, 0x00, 0x80, 0xff},
	5:  color.RGBA{0x80, 0x00, 0x80, 0xff},
	6:  color.RGBA{0x00, 0x80, 0x80, 0xff},
	7:  color.RGBA{0xc0, 0xc0, 0xc0, 0xff},
	8:  color.RGBA{0x80, 0x80, 0x80, 0xff},
	9:  color.RGBA{0xff, 0x00, 0x00, 0xff},
	10: color.RGBA{0x00, 0xff, 0x00, 0xff},
	11: color.RGBA{0xff, 0xff, 0x00, 0xff},
	12: color.RGBA{0x00, 0x00, 0xff, 0xff},
	13: color.RGBA{0xff, 0x00, 0xff, 0xff},
	14: color.RGBA{0x00, 0xff, 0xff, 0xff},
	15: color.RGBA{0xff, 0xff, 0xff, 0xff},

	// 6x6x6 cube block
	16:  color.RGBA{0x00, 0x00, 0x00, 0xff},
	17:  color.RGBA{0x00, 0x00, 0x5f, 0xff},
	18:  color.RGBA{0x00, 0x00, 0x87, 0xff},
	19:  color.RGBA{0x00, 0x00, 0xaf, 0xff},
	20:  color.RGBA{0x00, 0x00, 0xd7, 0xff},
	21:  color.RGBA{0x00, 0x00, 0xff, 0xff},
	22:  color.RGBA{0x00, 0x5f, 0x00, 0xff},
	23:  color.RGBA{0x00, 0x5f, 0x5f, 0xff},
	24:  color.RGBA{0x00, 0x5f, 0x87, 0xff},
	25:  color.RGBA{0x00, 0x5f, 0xaf, 0xff},
	26:  color.RGBA{0x00, 0x5f, 0xd7, 0xff},
	27:  color.RGBA{0x00, 0x5f, 0xff, 0xff},
	28:  color.RGBA{0x00, 0x87, 0x00, 0xff},
	29:  color.RGBA{0x00, 0x87, 0x5f, 0xff},
	30:  color.RGBA{0x00, 0x87, 0x87, 0xff},
	31:  color.RGBA{0x00, 0x87, 0xaf, 0xff},
	32:  color.RGBA{0x00, 0x87, 0xd7, 0xff},
	33:  color.RGBA{0x00, 0x87, 0xff, 0xff},
	34:  color.RGBA{0x00, 0xaf, 0x00, 0xff},
	35:  color.RGBA{0x00, 0xaf, 0x5f, 0xff},
	36:  color.RGBA{0x00, 0xaf, 0x87, 0xff},
	37:  color.RGBA{0x00, 0xaf, 0xaf, 0xff},
	38:  color.RGBA{0x00, 0xaf, 0xd7, 0xff},
	39:  color.RGBA{0x00, 0xaf, 0xff, 0xff},
	40:  color.RGBA{0x00, 0xd7, 0x00, 0xff},
	41:  color.RGBA{0x00, 0xd7, 0x5f, 0xff},
	42:  color.RGBA{0x00, 0xd7, 0x87, 0xff},
	43:  color.RGBA{0x00, 0xd7, 0xaf, 0xff},
	44:  color.RGBA{0x00, 0xd7, 0xd7, 0xff},
	45:  color.RGBA{0x00, 0xd7, 0xff, 0xff},
	46:  color.RGBA{0x00, 0xff, 0x00, 0xff},
	47:  color.RGBA{0x00, 0xff, 0x5f, 0xff},
	48:  color.RGBA{0x00, 0xff, 0x87, 0xff},
	49:  color.RGBA{0x00, 0xff, 0xaf, 0xff},
	50:  color.RGBA{0x00, 0xff, 0xd7, 0xff},
	51:  color.RGBA{0x00, 0xff, 0xff, 0xff},
	82:  color.RGBA{0x5f, 0xff, 0x00, 0xff},
	83:  color.RGBA{0x5f, 0xff, 0x5f, 0xff},
	84:  color.RGBA{0x5f, 0xff, 0x87, 0xff},
	85:  color.RGBA{0x5f, 0xff, 0xaf, 0xff},
	86:  color.RGBA{0x5f, 0xff, 0xd7, 0xff},
	87:  color.RGBA{0x5f, 0xff, 0xff, 0xff},
	76:  color.RGBA{0x5f, 0xd7, 0x00, 0xff},
	77:  color.RGBA{0x5f, 0xd7, 0x5f, 0xff},
	78:  color.RGBA{0x5f, 0xd7, 0x87, 0xff},
	79:  color.RGBA{0x5f, 0xd7, 0xaf, 0xff},
	80:  color.RGBA{0x5f, 0xd7, 0xd7, 0xff},
	81:  color.RGBA{0x5f, 0xd7, 0xff, 0xff},
	70:  color.RGBA{0x5f, 0xaf, 0x00, 0xff},
	71:  color.RGBA{0x5f, 0xaf, 0x5f, 0xff},
	72:  color.RGBA{0x5f, 0xaf, 0x87, 0xff},
	73:  color.RGBA{0x5f, 0xaf, 0xaf, 0xff},
	74:  color.RGBA{0x5f, 0xaf, 0xd7, 0xff},
	75:  color.RGBA{0x5f, 0xaf, 0xff, 0xff},
	64:  color.RGBA{0x5f, 0x87, 0x00, 0xff},
	65:  color.RGBA{0x5f, 0x87, 0x5f, 0xff},
	66:  color.RGBA{0x5f, 0x87, 0x87, 0xff},
	67:  color.RGBA{0x5f, 0x87, 0xaf, 0xff},
	68:  color.RGBA{0x5f, 0x87, 0xd7, 0xff},
	69:  color.RGBA{0x5f, 0x87, 0xff, 0xff},
	58:  color.RGBA{0x5f, 0x5f, 0x00, 0xff},
	59:  color.RGBA{0x5f, 0x5f, 0x5f, 0xff},
	60:  color.RGBA{0x5f, 0x5f, 0x87, 0xff},
	61:  color.RGBA{0x5f, 0x5f, 0xaf, 0xff},
	62:  color.RGBA{0x5f, 0x5f, 0xd7, 0xff},
	63:  color.RGBA{0x5f, 0x5f, 0xff, 0xff},
	52:  color.RGBA{0x5f, 0x00, 0x00, 0xff},
	53:  color.RGBA{0x5f, 0x00, 0x5f, 0xff},
	54:  color.RGBA{0x5f, 0x00, 0x87, 0xff},
	55:  color.RGBA{0x5f, 0x00, 0xaf, 0xff},
	56:  color.RGBA{0x5f, 0x00, 0xd7, 0xff},
	57:  color.RGBA{0x5f, 0x00, 0xff, 0xff},
	93:  color.RGBA{0x87, 0x00, 0xff, 0xff},
	92:  color.RGBA{0x87, 0x00, 0xd7, 0xff},
	91:  color.RGBA{0x87, 0x00, 0xaf, 0xff},
	90:  color.RGBA{0x87, 0x00, 0x87, 0xff},
	89:  color.RGBA{0x87, 0x00, 0x5f, 0xff},
	88:  color.RGBA{0x87, 0x00, 0x00, 0xff},
	99:  color.RGBA{0x87, 0x5f, 0xff, 0xff},
	98:  color.RGBA{0x87, 0x5f, 0xd7, 0xff},
	97:  color.RGBA{0x87, 0x5f, 0xaf, 0xff},
	96:  color.RGBA{0x87, 0x5f, 0x87, 0xff},
	95:  color.RGBA{0x87, 0x5f, 0x5f, 0xff},
	94:  color.RGBA{0x87, 0x5f, 0x00, 0xff},
	105: color.RGBA{0x87, 0x87, 0xff, 0xff},
	104: color.RGBA{0x87, 0x87, 0xd7, 0xff},
	103: color.RGBA{0x87, 0x87, 0xaf, 0xff},
	102: color.RGBA{0x87, 0x87, 0x87, 0xff},
	101: color.RGBA{0x87, 0x87, 0x5f, 0xff},
	100: color.RGBA{0x87, 0x87, 0x00, 0xff},
	111: color.RGBA{0x87, 0xaf, 0xff, 0xff},
	110: color.RGBA{0x87, 0xaf, 0xd7, 0xff},
	109: color.RGBA{0x87, 0xaf, 0xaf, 0xff},
	108: color.RGBA{0x87, 0xaf, 0x87, 0xff},
	107: color.RGBA{0x87, 0xaf, 0x5f, 0xff},
	106: color.RGBA{0x87, 0xaf, 0x00, 0xff},
	117: color.RGBA{0x87, 0xd7, 0xff, 0xff},
	116: color.RGBA{0x87, 0xd7, 0xd7, 0xff},
	115: color.RGBA{0x87, 0xd7, 0xaf, 0xff},
	114: color.RGBA{0x87, 0xd7, 0x87, 0xff},
	113: color.RGBA{0x87, 0xd7, 0x5f, 0xff},
	112: color.RGBA{0x87, 0xd7, 0x00, 0xff},
	123: color.RGBA{0x87, 0xff, 0xff, 0xff},
	122: color.RGBA{0x87, 0xff, 0xd7, 0xff},
	121: color.RGBA{0x87, 0xff, 0xaf, 0xff},
	120: color.RGBA{0x87, 0xff, 0x87, 0xff},
	119: color.RGBA{0x87, 0xff, 0x5f, 0xff},
	118: color.RGBA{0x87, 0xff, 0x00, 0xff},
	159: color.RGBA{0xaf, 0xff, 0xff, 0xff},
	158: color.RGBA{0xaf, 0xff, 0xd7, 0xff},
	157: color.RGBA{0xaf, 0xff, 0xaf, 0xff},
	156: color.RGBA{0xaf, 0xff, 0x87, 0xff},
	155: color.RGBA{0xaf, 0xff, 0x5f, 0xff},
	154: color.RGBA{0xaf, 0xff, 0x00, 0xff},
	153: color.RGBA{0xaf, 0xd7, 0xff, 0xff},
	152: color.RGBA{0xaf, 0xd7, 0xd7, 0xff},
	151: color.RGBA{0xaf, 0xd7, 0xaf, 0xff},
	150: color.RGBA{0xaf, 0xd7, 0x87, 0xff},
	149: color.RGBA{0xaf, 0xd7, 0x5f, 0xff},
	148: color.RGBA{0xaf, 0xd7, 0x00, 0xff},
	147: color.RGBA{0xaf, 0xaf, 0xff, 0xff},
	146: color.RGBA{0xaf, 0xaf, 0xd7, 0xff},
	145: color.RGBA{0xaf, 0xaf, 0xaf, 0xff},
	144: color.RGBA{0xaf, 0xaf, 0x87, 0xff},
	143: color.RGBA{0xaf, 0xaf, 0x5f, 0xff},
	142: color.RGBA{0xaf, 0xaf, 0x00, 0xff},
	141: color.RGBA{0xaf, 0x87, 0xff, 0xff},
	140: color.RGBA{0xaf, 0x87, 0xd7, 0xff},
	139: color.RGBA{0xaf, 0x87, 0xaf, 0xff},
	138: color.RGBA{0xaf, 0x87, 0x87, 0xff},
	137: color.RGBA{0xaf, 0x87, 0x5f, 0xff},
	136: color.RGBA{0xaf, 0x87, 0x00, 0xff},
	135: color.RGBA{0xaf, 0x5f, 0xff, 0xff},
	134: color.RGBA{0xaf, 0x5f, 0xd7, 0xff},
	133: color.RGBA{0xaf, 0x5f, 0xaf, 0xff},
	132: color.RGBA{0xaf, 0x5f, 0x87, 0xff},
	131: color.RGBA{0xaf, 0x5f, 0x5f, 0xff},
	130: color.RGBA{0xaf, 0x5f, 0x00, 0xff},
	129: color.RGBA{0xaf, 0x00, 0xff, 0xff},
	128: color.RGBA{0xaf, 0x00, 0xd7, 0xff},
	127: color.RGBA{0xaf, 0x00, 0xaf, 0xff},
	126: color.RGBA{0xaf, 0x00, 0x87, 0xff},
	125: color.RGBA{0xaf, 0x00, 0x5f, 0xff},
	124: color.RGBA{0xaf, 0x00, 0x00, 0xff},
	160: color.RGBA{0xd7, 0x00, 0x00, 0xff},
	161: color.RGBA{0xd7, 0x00, 0x5f, 0xff},
	162: color.RGBA{0xd7, 0x00, 0x87, 0xff},
	163: color.RGBA{0xd7, 0x00, 0xaf, 0xff},
	164: color.RGBA{0xd7, 0x00, 0xd7, 0xff},
	165: color.RGBA{0xd7, 0x00, 0xff, 0xff},
	166: color.RGBA{0xd7, 0x5f, 0x00, 0xff},
	167: color.RGBA{0xd7, 0x5f, 0x5f, 0xff},
	168: color.RGBA{0xd7, 0x5f, 0x87, 0xff},
	169: color.RGBA{0xd7, 0x5f, 0xaf, 0xff},
	170: color.RGBA{0xd7, 0x5f, 0xd7, 0xff},
	171: color.RGBA{0xd7, 0x5f, 0xff, 0xff},
	172: color.RGBA{0xd7, 0x87, 0x00, 0xff},
	173: color.RGBA{0xd7, 0x87, 0x5f, 0xff},
	174: color.RGBA{0xd7, 0x87, 0x87, 0xff},
	175: color.RGBA{0xd7, 0x87, 0xaf, 0xff},
	176: color.RGBA{0xd7, 0x87, 0xd7, 0xff},
	177: color.RGBA{0xd7, 0x87, 0xff, 0xff},
	178: color.RGBA{0xdf, 0xaf, 0x00, 0xff},
	179: color.RGBA{0xdf, 0xaf, 0x5f, 0xff},
	180: color.RGBA{0xdf, 0xaf, 0x87, 0xff},
	181: color.RGBA{0xdf, 0xaf, 0xaf, 0xff},
	182: color.RGBA{0xdf, 0xaf, 0xdf, 0xff},
	183: color.RGBA{0xdf, 0xaf, 0xff, 0xff},
	184: color.RGBA{0xdf, 0xdf, 0x00, 0xff},
	185: color.RGBA{0xdf, 0xdf, 0x5f, 0xff},
	186: color.RGBA{0xdf, 0xdf, 0x87, 0xff},
	187: color.RGBA{0xdf, 0xdf, 0xaf, 0xff},
	188: color.RGBA{0xdf, 0xdf, 0xdf, 0xff},
	189: color.RGBA{0xdf, 0xdf, 0xff, 0xff},
	190: color.RGBA{0xdf, 0xff, 0x00, 0xff},
	191: color.RGBA{0xdf, 0xff, 0x5f, 0xff},
	192: color.RGBA{0xdf, 0xff, 0x87, 0xff},
	193: color.RGBA{0xdf, 0xff, 0xaf, 0xff},
	194: color.RGBA{0xdf, 0xff, 0xdf, 0xff},
	195: color.RGBA{0xdf, 0xff, 0xff, 0xff},
	226: color.RGBA{0xff, 0xff, 0x00, 0xff},
	227: color.RGBA{0xff, 0xff, 0x5f, 0xff},
	228: color.RGBA{0xff, 0xff, 0x87, 0xff},
	229: color.RGBA{0xff, 0xff, 0xaf, 0xff},
	230: color.RGBA{0xff, 0xff, 0xdf, 0xff},
	231: color.RGBA{0xff, 0xff, 0xff, 0xff},
	220: color.RGBA{0xff, 0xdf, 0x00, 0xff},
	221: color.RGBA{0xff, 0xdf, 0x5f, 0xff},
	222: color.RGBA{0xff, 0xdf, 0x87, 0xff},
	223: color.RGBA{0xff, 0xdf, 0xaf, 0xff},
	224: color.RGBA{0xff, 0xdf, 0xdf, 0xff},
	225: color.RGBA{0xff, 0xdf, 0xff, 0xff},
	214: color.RGBA{0xff, 0xaf, 0x00, 0xff},
	215: color.RGBA{0xff, 0xaf, 0x5f, 0xff},
	216: color.RGBA{0xff, 0xaf, 0x87, 0xff},
	217: color.RGBA{0xff, 0xaf, 0xaf, 0xff},
	218: color.RGBA{0xff, 0xaf, 0xdf, 0xff},
	219: color.RGBA{0xff, 0xaf, 0xff, 0xff},
	208: color.RGBA{0xff, 0x87, 0x00, 0xff},
	209: color.RGBA{0xff, 0x87, 0x5f, 0xff},
	210: color.RGBA{0xff, 0x87, 0x87, 0xff},
	211: color.RGBA{0xff, 0x87, 0xaf, 0xff},
	212: color.RGBA{0xff, 0x87, 0xdf, 0xff},
	213: color.RGBA{0xff, 0x87, 0xff, 0xff},
	202: color.RGBA{0xff, 0x5f, 0x00, 0xff},
	203: color.RGBA{0xff, 0x5f, 0x5f, 0xff},
	204: color.RGBA{0xff, 0x5f, 0x87, 0xff},
	205: color.RGBA{0xff, 0x5f, 0xaf, 0xff},
	206: color.RGBA{0xff, 0x5f, 0xdf, 0xff},
	207: color.RGBA{0xff, 0x5f, 0xff, 0xff},
	196: color.RGBA{0xff, 0x00, 0x00, 0xff},
	197: color.RGBA{0xff, 0x00, 0x5f, 0xff},
	198: color.RGBA{0xff, 0x00, 0x87, 0xff},
	199: color.RGBA{0xff, 0x00, 0xaf, 0xff},
	200: color.RGBA{0xff, 0x00, 0xdf, 0xff},
	201: color.RGBA{0xff, 0x00, 0xff, 0xff},

	// Grayscale block
	232: color.RGBA{0x08, 0x08, 0x08, 0xff},
	233: color.RGBA{0x12, 0x12, 0x12, 0xff},
	234: color.RGBA{0x1c, 0x1c, 0x1c, 0xff},
	235: color.RGBA{0x26, 0x26, 0x26, 0xff},
	236: color.RGBA{0x30, 0x30, 0x30, 0xff},
	237: color.RGBA{0x3a, 0x3a, 0x3a, 0xff},
	238: color.RGBA{0x44, 0x44, 0x44, 0xff},
	239: color.RGBA{0x4e, 0x4e, 0x4e, 0xff},
	240: color.RGBA{0x58, 0x58, 0x58, 0xff},
	241: color.RGBA{0x62, 0x62, 0x62, 0xff},
	242: color.RGBA{0x6c, 0x6c, 0x6c, 0xff},
	243: color.RGBA{0x76, 0x76, 0x76, 0xff},
	255: color.RGBA{0xee, 0xee, 0xee, 0xff},
	254: color.RGBA{0xe4, 0xe4, 0xe4, 0xff},
	253: color.RGBA{0xda, 0xda, 0xda, 0xff},
	252: color.RGBA{0xd0, 0xd0, 0xd0, 0xff},
	251: color.RGBA{0xc6, 0xc6, 0xc6, 0xff},
	250: color.RGBA{0xbc, 0xbc, 0xbc, 0xff},
	249: color.RGBA{0xb2, 0xb2, 0xb2, 0xff},
	248: color.RGBA{0xa8, 0xa8, 0xa8, 0xff},
	247: color.RGBA{0x9e, 0x9e, 0x9e, 0xff},
	246: color.RGBA{0x94, 0x94, 0x94, 0xff},
	245: color.RGBA{0x8a, 0x8a, 0x8a, 0xff},
	244: color.RGBA{0x80, 0x80, 0x80, 0xff},
}

var colorCode4BitsCache map[int]int = map[int]int{}

func (c Color) Code4Bits() int {
	if c == 0 {
		return -1
	}
	x := int(c - 1)
	if code, ok := colorCode4BitsCache[x]; ok {
		return code
	}
	r := uint8((x >> 16) & 0xff)
	g := uint8((x >> 8) & 0xff)
	b := uint8(x & 0xff)
	var code int
	code = ColorPalette4Bits.Index(color.RGBA{r, g, b, 0xff})
	colorCode4BitsCache[x] = code
	return code
}

var colorCode8BitsCache map[int]int = map[int]int{}

func (c Color) Code8Bits() int {
	if c == 0 {
		return -1
	}
	x := int(c - 1)
	if code, ok := colorCode8BitsCache[x]; ok {
		return code
	}
	r := uint8((x >> 16) & 0xff)
	g := uint8((x >> 8) & 0xff)
	b := uint8(x & 0xff)
	var code int
	code = ColorPalette8Bits.Index(color.RGBA{r, g, b, 0xff})
	colorCode8BitsCache[x] = code
	return code
}
