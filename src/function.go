package yasm

type Function struct {
	Name   string
	Caller *Function
	start  addr
	pc     int
}
