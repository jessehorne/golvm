package vm

type Upvalue struct {
	Name string
}

func ReadUpvalue(current int64, bc []byte) (Upvalue, int64) {
	name, n := ReadString(current, bc)
	return Upvalue{
		Name: name,
	}, int64(n)
}
