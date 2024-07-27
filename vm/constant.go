package vm

import (
	"math"
	"strconv"
)

const (
	ConstantTypeNil    = 0x0
	ConstantTypeBool   = 0x01
	ConstantTypeNumber = 0x03
	ConstantTypeString = 0x04
)

type Constant struct {
	Type        byte
	IsNil       bool
	BoolValue   bool
	NumberValue float64
	StringValue string
}

func NewConstant(t byte) *Constant {
	return &Constant{
		Type: t,
	}
}

// ReadConstants reads all the constants and returns them in an array as well as the next bytecode index
// to continue processing the bytecode.
func ReadConstants(current int64, bc []byte, count int64) ([]*Constant, uint64) {
	var constants []*Constant

	var i int64
	for i = 0; i < count; i++ {
		//fmt.Println(current)
		t := bc[current]
		current++

		c := NewConstant(t)

		if t == ConstantTypeNil {
			c.IsNil = true
			current++ // nils don't exist. We can just skip to the next constant
		} else if t == ConstantTypeBool {
			num := ReadUInt32(current, bc)
			if num == 0 {
				c.BoolValue = false
			} else {
				c.BoolValue = true
			}
			current += 4 // skip 4 bytes because booleans are 4 byte integers (0 = false)
		} else if t == ConstantTypeNumber {
			num := ReadUInt64(current, bc)
			floatValue := math.Float64frombits(num)
			c.NumberValue = floatValue
			current += 8 // skip 8 bytes because numbers are 8 byte doubles
		} else if t == ConstantTypeString {
			str, next := ReadString(current, bc)
			current = int64(next)
			c.StringValue = str
		}

		constants = append(constants, c)
	}

	return constants, uint64(current)
}

func (c *Constant) String() string {
	if c.Type == ConstantTypeNil {
		return "nil"
	} else if c.Type == ConstantTypeBool {
		if c.BoolValue {
			return "true"
		} else {
			return "false"
		}
	} else if c.Type == ConstantTypeNumber {
		return strconv.FormatFloat(c.NumberValue, 'f', -1, 64)
	} else if c.Type == ConstantTypeString {
		return c.StringValue
	}

	return "invalid"
}
