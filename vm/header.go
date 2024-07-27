package vm

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type GlobalHeader struct {
	ValidSignature  bool
	Version         byte // 0x51 == Lua 5.1 and so on
	Format          byte // 0x0 == official
	Endianness      byte // 1 == little endian
	IntSize         byte
	TSize           byte
	InstructionSize byte
	NumberSize      byte
	Integral        byte
}

func NewGlobalHeader(bc []byte) (*GlobalHeader, error) {
	if len(bc) < 12 {
		return nil, errors.New("invalid header (too small)")
	}

	return &GlobalHeader{
		ValidSignature:  binary.BigEndian.Uint32(bc[0:4]) == 0x1B4C7561,
		Version:         bc[4],  // 0x51 == Lua 5.1
		Format:          bc[5],  // 0 == Official
		Endianness:      bc[6],  // 1 == Little Endian
		IntSize:         bc[7],  // bytes
		TSize:           bc[8],  // bytes
		InstructionSize: bc[9],  // bytes
		NumberSize:      bc[10], // bytes
		Integral:        bc[11], // 1 == integral (integers can be signed OR unsigned)
	}, nil
}

func (h *GlobalHeader) String() string {
	tmpl := `Valid Signature: %t
Version: %x
Format: %s
Endianness: %s
IntSize: %d
TSize: %d
InstructionSize: %d
NumberSize: %d
Integral: %t`
	var format string
	if h.Format == byte(0) {
		format = "Official"
	} else {
		format = "Unofficial"
	}

	var endianness string
	if h.Endianness == byte(1) {
		endianness = "Little Endian"
	} else {
		endianness = "Big Endian"
	}

	return fmt.Sprintf(tmpl, h.ValidSignature, h.Version, format, endianness, h.IntSize, h.TSize, h.InstructionSize,
		h.NumberSize, h.Integral == byte(1))
}
