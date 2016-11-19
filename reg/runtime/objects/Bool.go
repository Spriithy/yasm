package objects

type Bool bool

func (b Bool) Class() string {
	return "bool"
}

func (b Bool) String() string {
	if b {
		return "true"
	}
	return "false"
}

func (b Bool) Copy() Any {
	return Bool(b)
}
