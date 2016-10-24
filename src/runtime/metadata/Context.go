package metadata

import "github.com/Spriithy/Polaroid/src/builtins"

type Context struct {
	Caller   *Context
	Metadata FunctionMetaData
	Rip      int
	Locals   []builtins.Object
}

func CreateContext(caller *Context, md FunctionMetaData, ip int) *Context {
	c := new(Context)
	c.Metadata = md
	c.Rip = ip
	c.Caller = caller
	c.Locals = make([]builtins.Object, md.Nargs + md.Nlocals)
	return c
}