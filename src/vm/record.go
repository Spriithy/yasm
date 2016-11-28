package small

type record struct {
	r      []word
	locals []word
	caller *record
}
