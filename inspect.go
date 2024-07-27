package main

import (
	"fmt"
	"github.com/jessehorne/golvm/vm"
)

// inspect prints helpful information about the compiled lua bytecode
func inspect(path string) error {
	_, err := vm.NewBytecode(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}
