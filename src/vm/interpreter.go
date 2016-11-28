package small

// A Machine is the exported Machine interface
// so fields are presereved during execution
type Machine interface {
	Start()
}

type machine struct {
	code  []Instruction
	globs []word
	frame *frame
	r     []word
}

// Interepreter creates a new machine object that can be then
// started.
func Interepreter(path string, globs []word) Machine {
	m := new(machine)
	m.globs = globs
	m.frame = &frame{nil, m.globs, nil}
	m.r = make([]word, 12)
	m.code = Load(path)
	return m
}

func (m *machine) swi() {
	code := m.r[10]
	println("syscall", code)
}

func (m *machine) Start() {
	c := make(chan int)
	l := word(len(m.code))

	var (
		i  Instruction
		pc *word = &m.r[8]
		sp *word = &m.r[9]
		ic *word = &m.r[10]
		rx *word = &m.r[11]
	)

	go func() {
		for {
			if *pc >= l {
				println("halted")
			}
			i = m.code[*pc]
			switch i.O() {
			case SWI:
				(*ic).set(i.E())
				m.swi()
			case MOV:
				switch i.X() {
				case 0x0:
					m.r[i.RA()] = m.r[i.RB()]
				case 0x1:
					m.r[i.RA()].set(i.E())
				case 0x2:
					m.r[i.RA()].set(m.code[*pc+1].Int32())
					*pc++
				case 0x3:
					m.r[i.RA()] = word(m.code[*pc+1].Int64(&m.code[*pc+2]))
					*pc += 2
				}
			}

			*pc++
			c <- 0
		}
	}()
	<-c
}
