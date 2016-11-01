package runtime

import "fmt"

const (
	INITIAL_STACK_SIZE = 1 << 10
	STACK_GROW_SIZE = 1 << 6
	STACK_SHRINK_SIZE = 1 << 5
)

var stackExtensionChunk = make([]*Value, STACK_GROW_SIZE)

type vm struct {
	stack []*Value
	sp    int
}

func VirtualMachine() *vm {
	v := new(vm)
	v.stack = make([]*Value, INITIAL_STACK_SIZE)
	v.sp = -1
	return v
}












func (v *vm) Pop() *Value {
	if len(v.stack) - v.sp > STACK_SHRINK_SIZE {
		v.stack = v.stack[:len(v.stack) - STACK_SHRINK_SIZE]
	}
	item := v.stack[v.sp]
	v.sp--
	return item
}

func (v *vm) Push(item interface{}) {
	if len(v.stack) - v.sp < 1 {
		v.stack = append(v.stack, stackExtensionChunk ...)
	}
	v.sp++
	v.stack[v.sp] = Wrap(item)
}

func (v *vm) peek() *Value {
	return v.stack[v.sp]
}

func (v *vm) empty() bool {
	return v.sp < 0
}

func (v *vm) StackString() string {
	str := "["
	for i := v.sp - 1; i >= 0; i-- {
		str += fmt.Sprintf("%v", v.stack[i].ToString())
		if i != 0 {
			str += ", "
		}
	}
	return str + "]"
}