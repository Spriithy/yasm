package small

type machine struct {
	frame *record
	r     []word
	pc    *word
	sp    *word
	ic    *word
	rx    *word
}

func (m *machine) 