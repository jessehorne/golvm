package vm

import (
	"fmt"
	"os"
)

type Bytecode struct {
	Data           []byte
	GlobalHeader   *GlobalHeader
	FunctionBlocks []*Function
	Path           string
}

func NewBytecode(path string) (*Bytecode, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	gh, err := NewGlobalHeader(data)
	if err != nil {
		return nil, err
	}

	fmt.Println(gh)

	var funcs []*Function
	var current int64 = 12
	for {
		var f *Function
		current, f, err = ReadFunction(true, current, data)

		if err != nil {
			fmt.Println(err)
			break
		}
		funcs = append(funcs, f)
		if current >= int64(len(data)) {
			break
		}
		break
	}

	return &Bytecode{
		Data:           data,
		Path:           path,
		GlobalHeader:   gh,
		FunctionBlocks: funcs,
	}, nil
}
