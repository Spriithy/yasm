package small

type frame struct {
	r      []word
	locals []word
	caller *frame
}
