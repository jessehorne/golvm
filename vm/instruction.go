package vm

type Instruction struct {
	Data []byte
}

func NewInstruction(data []byte) *Instruction {
	return &Instruction{
		Data: data,
	}
}

// ReadInstructions takes in a starting bytecode index, the bytecode byte array and the count or
// number of instructions that are expected. It returns a list of Instruction's
func ReadInstructions(start int64, bc []byte, count int64) []*Instruction {
	var instructions []*Instruction

	var i int64
	for i = start; i < start+count*4; i += 4 {
		var inc []byte
		var x int64
		for x = 0; x < 4; x++ {
			inc = append(inc, bc[i+x])
		}

		newInc := NewInstruction(inc)
		instructions = append(instructions, newInc)
	}

	return instructions
}
