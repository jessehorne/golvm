package main

import "fmt"

func printHelp() {
	fmt.Println(`	golvm - A Lua 5.1 VM

	help				Show help information.
	inspect <filepath>		Show details on Lua 5.1 bytecode.
	run <filepath>			Run a compiled Lua 5.1 bytecode program.`)
}
