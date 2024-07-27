package vm

import (
	"fmt"
)

type Function struct {
	IsTopLevel      bool
	Source          string
	LineStart       int32
	LineEnd         int32
	UpValues        byte
	Parameters      byte
	VarArg          byte
	MaxStackSize    byte
	Instructions    []*Instruction
	Constants       []*Constant
	Functions       []*Function
	SourceLineCount uint32
	//SourceLinePositions []SourceLinePosition
	Locals   []Local
	Upvalues []Upvalue
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
SourceLineCount: %v
LocalCount: %v
UpvalueCount: %v
`
	data := fmt.Sprintf(tmpl, f.Source, f.LineStart, f.LineEnd, f.UpValues, f.Parameters,
		f.VarArg, f.MaxStackSize, len(f.Instructions), len(f.Constants), f.SourceLineCount,
		len(f.Locals), len(f.Upvalues))

	data += "\nLocals...\n"
	for _, l := range f.Locals {
		data += l.Name + "\n"
	}

	data += "\nUpvalues...\n"
	for _, uv := range f.Upvalues {
		data += uv.Name + "\n"
	}

	return data
}

// ReadFunction takes in Lua 5.1 bytecode, starts at the "start" position (0 being first byte) and reads
// byte by byte until it finds the end of the function declaraction, then returns where it ended, the function
// and hopefully doesn't return an error.
func ReadFunction(topLevel bool, start int64, bc []byte) (int64, *Function, error) {
	current := start // current bytecode index
	var sourceName string
	if topLevel {
		var n uint64
		// Get the source name which is start to ascii nul
		sourceName, n = ReadString(current, bc)
		current = int64(n)
	} else {
		current += 8
	}

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
	funcCount := ReadUInt32(current, bc)
	current += 4

	// read function prototypes
	var funcs []*Function
	var i uint32
	for i = 0; i < funcCount; i++ {
		n, f, err := ReadFunction(false, current, bc)
		if err != nil {
			fmt.Println(err)
		}
		funcs = append(funcs, f)
		current = n
	}

	// read integer number of source line positions
	sourceLineCount := ReadUInt32(current, bc)
	current += 4 // skip to the start of the source line positions

	// read source line positions (IGNORE FOR NOW)
	current += int64(sourceLineCount) * 4 // skip all the source lines

	// read integer number of locals
	localCount := ReadUInt32(current, bc)
	current += 4

	// read locals
	var locals []Local
	for i = 0; i < localCount; i++ {
		local, n := ReadLocal(current, bc)
		current = n
		locals = append(locals, local)
	}

	// read integer number of upvalues
	uvc := ReadUInt32(current, bc)
	current += 4

	// read upvalues
	var upvalues []Upvalue
	for i = 0; i < uvc; i++ {
		newUpvalue, n := ReadUpvalue(current, bc)
		current = n
		upvalues = append(upvalues, newUpvalue)
	}

	newFunc := &Function{
		IsTopLevel:      topLevel,
		Source:          sourceName,
		LineStart:       int32(line),
		LineEnd:         int32(lastLine),
		UpValues:        upvaluesCount,
		Parameters:      paramsCount,
		VarArg:          isVarArg,
		MaxStackSize:    maxStackSize,
		Instructions:    instructions,
		Constants:       constants,
		SourceLineCount: sourceLineCount,
		Locals:          locals,
		Upvalues:        upvalues,
	}

	fmt.Println(newFunc)

	return current, newFunc, nil
}
