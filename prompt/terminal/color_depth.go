package terminal

type ColorDepth int

const (
	ColorDepth1Bit ColorDepth = iota
	ColorDepth4Bits
	ColorDepth8Bits
	ColorDepth24Bits
)
