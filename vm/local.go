package vm

type Local struct {
	Name       string
	StartScope int32
	EndScope   int32
}

func NewLocal() {

}

func ReadLocal(current int64, bc []byte) (Local, int64) {
	name, next := ReadString(current, bc)
	current = int64(next)

	startScope := ReadUInt32(current, bc)
	current += 4

	endScope := ReadUInt32(current, bc)
	current += 4

	return Local{
		Name:       name,
		StartScope: int32(startScope),
		EndScope:   int32(endScope),
	}, current
}
