package builtins

import "os"

type File struct {
	name String
	file *os.File
}

// TODO
