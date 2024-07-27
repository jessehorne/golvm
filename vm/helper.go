package vm

import (
	"encoding/binary"
)

func ReadUInt32(start int64, bc []byte) uint32 {
	var intBytes []byte
	for i := start; i < start+4; i++ {
		intBytes = append(intBytes, bc[i])
	}

	return binary.LittleEndian.Uint32(intBytes)
}

func ReadUInt64(start int64, bc []byte) uint64 {
	var intBytes []byte
	for i := start; i < start+8; i++ {
		intBytes = append(intBytes, bc[i])
	}

	return binary.LittleEndian.Uint64(intBytes)
}

func ReadString(start int64, bc []byte) (string, uint64) {
	// get string size
	stringSize := ReadUInt64(start, bc)
	start += 8
	str := bc[start : uint64(start)+stringSize]

	return string(str), uint64(start) + stringSize
}
