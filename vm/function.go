package vm

import "fmt"

type Function struct {
	Source       string
	LineStart    int32
	LineEnd      int32
	UpValues     byte
	Parameters   byte
	VarArg       byte
	MaxStackSize byte
	Instructions []*Instruction
	Constants    []*Constant
	//FunctionPrototypes  []FunctionPrototype
	//SourceLinePositions []SourceLinePosition
	//Locals              []Local
	//Upvalues            []Upvalue
}

func (f *Function) String() string {
	tmpl := `Source: %s
LineStart: %v
LineEnd: %v
UpValues: %v
Parameters: %v
VarArg: %v
MaxStackSize: %x
InstructionCount: %v
ConstantCount: %v
`
	return fmt.Sprintf(tmpl, f.Source, f.LineStart, f.LineEnd, f.UpValues, f.Parameters,
		f.VarArg, f.MaxStackSize, len(f.Instructions), len(f.Constants))
}

// ReadFunction takes in Lua 5.1 bytecode, starts at the "start" position (0 being first byte) and reads
// byte by byte until it finds the end of the function declaraction, then returns where it ended, the function
// and hopefully doesn't return an error.
func ReadFunction(start int64, bc []byte) (int64, *Function, error) {
	current := start // current bytecode index

	// Get the source name which is start to ascii nul
	sourceName, next := ReadString(current, bc)
	current = int64(next)

	// read int32 line defined
	line := ReadUInt32(current, bc)
	current += 4

	// read int32 last line defined
	lastLine := ReadUInt32(current, bc)
	current += 4

	// read 1 byte upvalues count
	upvaluesCount := bc[current]
	current++

	// read 1 byte parameters count
	paramsCount := bc[current]
	current++

	// read 1 byte is_vararg flag (ignore for now)
	isVarArg := bc[current]
	current++

	// read 1 byte max stack size (number of registers used)
	maxStackSize := bc[current]
	current++

	// read integer number of instructions
	iCount := ReadUInt32(current, bc)
	current += 4

	// read instructions
	//fmt.Println("CURRENT: ", byte(current))
	instructions := ReadInstructions(current, bc, int64(iCount))
	current += int64(len(instructions) * 4) // skip the counter forward past instructions

	// read integer number of constants
	constantCount := ReadUInt32(current, bc)
	current += 4

	// read constants
	constants, next := ReadConstants(current, bc, int64(constantCount))
	current = int64(next)

	// read integer number of function prototypes

	// read function prototypes

	// read integer number of source line positions

	// read source line positions (IGNORE FOR NOW)

	// read integer number of locals

	// read locals

	// read integer number of upvalues

	// read upvalues

	return current, &Function{
		Source:       sourceName,
		LineStart:    int32(line),
		LineEnd:      int32(lastLine),
		UpValues:     upvaluesCount,
		Parameters:   paramsCount,
		VarArg:       isVarArg,
		MaxStackSize: maxStackSize,
		Instructions: instructions,
		Constants:    constants,
	}, nil
}
