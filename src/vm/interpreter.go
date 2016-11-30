package small

// A Machine is the exported Machine interface
// so fields are presereved during execution
type Machine interface {
	Start()
}

type machine struct {
	code  []Instruction
	mem   []word
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
	m.mem = make([]word, 0x1000)
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
		// sp *word = &m.r[9]
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
			case MOM:
				m.mem[m.r[i.RA()].int64()] = m.mem[m.r[i.RB()].int64()]
			case LOA:
				m.r[i.RA()] = m.mem[m.r[i.RB()].int64()]
			case STR:
				switch i.X() {
				case 0x0:
					m.mem[i.RA()] = m.r[i.RB()]
				case 0x1:
					m.mem[i.RA()] = word(i.E())
				case 0x2:
					m.mem[i.RA()] = word(m.code[*pc+1].Int32())
					*pc++
				case 0x3:
					m.mem[i.RA()] = word(m.code[*pc+1].Int64(&m.code[*pc+2]))
					*pc += 2
				}
			case ADD:
				switch i.X() {
				case 0x0:
					switch {
					case i.T():
						*rx = m.r[i.RA()] + m.r[i.RB()]
					default:
						(*rx).set(m.r[i.RA()].float64() + m.r[i.RB()].float64())
					}
				case 0x1:
					*rx = m.r[i.RA()] + word(i.E())
				case 0x2:
					switch {
					case i.T():
						*rx = m.r[i.RA()] + word(m.code[*pc+1])
					default:
						*rx = word(m.r[i.RA()].float32() + m.code[*pc+1].Float32())
					}
					*pc++
				case 0x3:
					switch {
					case i.T():
						*rx = word(m.r[i.RA()].int64() + m.code[*pc+1].Int64(&m.code[*pc+2]))
					default:
						*rx = word(m.r[i.RA()].float64() + m.code[*pc+1].Float64(m.code[*pc+2]))
					}
					*pc += 2
				}
			case SUB:
				switch i.X() {
				case 0x0:
					switch {
					case i.T():
						*rx = m.r[i.RA()] - m.r[i.RB()]
					default:
						(*rx).set(m.r[i.RA()].float64() - m.r[i.RB()].float64())
					}
				case 0x1:
					*rx = m.r[i.RA()] - word(i.E())
				case 0x2:
					switch {
					case i.T():
						*rx = m.r[i.RA()] - word(m.code[*pc+1])
					default:
						*rx = word(m.r[i.RA()].float32() - m.code[*pc+1].Float32())
					}
					*pc++
				case 0x3:
					switch {
					case i.T():
						*rx = word(m.r[i.RA()].int64() - m.code[*pc+1].Int64(&m.code[*pc+2]))
					default:
						*rx = word(m.r[i.RA()].float64() - m.code[*pc+1].Float64(m.code[*pc+2]))
					}
					*pc += 2
				}
			case MUL:
				switch i.X() {
				case 0x0:
					switch {
					case i.T():
						*rx = m.r[i.RA()] * m.r[i.RB()]
					default:
						(*rx).set(m.r[i.RA()].float64() * m.r[i.RB()].float64())
					}
				case 0x1:
					*rx = m.r[i.RA()] * word(i.E())
				case 0x2:
					switch {
					case i.T():
						*rx = m.r[i.RA()] * word(m.code[*pc+1])
					default:
						*rx = word(m.r[i.RA()].float32() * m.code[*pc+1].Float32())
					}
					*pc++
				case 0x3:
					switch {
					case i.T():
						*rx = word(m.r[i.RA()].int64() * m.code[*pc+1].Int64(&m.code[*pc+2]))
					default:
						*rx = word(m.r[i.RA()].float64() * m.code[*pc+1].Float64(m.code[*pc+2]))
					}
					*pc += 2
				}
			case DIV:
				switch i.X() {
				case 0x0:
					switch {
					case i.T():
						*rx = m.r[i.RA()] / m.r[i.RB()]
					default:
						(*rx).set(m.r[i.RA()].float64() / m.r[i.RB()].float64())
					}
				case 0x1:
					*rx = m.r[i.RA()] / word(i.E())
				case 0x2:
					switch {
					case i.T():
						*rx = m.r[i.RA()] / word(m.code[*pc+1])
					default:
						*rx = word(m.r[i.RA()].float32() / m.code[*pc+1].Float32())
					}
					*pc++
				case 0x3:
					switch {
					case i.T():
						*rx = word(m.r[i.RA()].int64() / m.code[*pc+1].Int64(&m.code[*pc+2]))
					default:
						*rx = word(m.r[i.RA()].float64() / m.code[*pc+1].Float64(m.code[*pc+2]))
					}
					*pc += 2
				}
			case REM:
				switch i.X() {
				case 0x0:
					if i.T() {
						*rx = m.r[i.RA()] % m.r[i.RB()]
					}
				case 0x1:
					*rx = m.r[i.RA()] % word(i.E())
				case 0x2:
					if i.T() {
						*rx = m.r[i.RA()] % word(m.code[*pc+1])
					}
					*pc++
				case 0x3:
					if i.T() {
						*rx = word(m.r[i.RA()].int64() % m.code[*pc+1].Int64(&m.code[*pc+2]))
					}
					*pc += 2
				}
			case BSL:
			case BSR:
			case INC:
			case DEC:
			case AND:
			case IOR:
			case XOR:
			case NOT:
			case JMP:
			case JZ:
			case JNZ:
			case JEQ:
			case JNE:
			case JLT:
			case JLE:
			case JGT:
			case JGE:
			case SRL:
			case RET:
			}

			*pc++
			c <- 0
		}
	}()
	<-c
}
