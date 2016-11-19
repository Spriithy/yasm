package objects

type Null struct{}

func (n Null) Class() string {
	return "null"
}

func (n Null) String() string {
	return "null"
}

func (n Null) Copy() Any {
	return Null{}
}
