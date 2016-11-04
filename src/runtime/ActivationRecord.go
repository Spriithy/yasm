package runtime

type ActivationRecord struct {
	Function
	caller *ActivationRecord
	locals []Object
	retip  int
}

func Call(caller *ActivationRecord, callee Function, from int) *ActivationRecord {
	assert(callee.nlocs >= 0, "Local variable count must be positive or zero!")
	ar := new(ActivationRecord)
	ar.Function = callee
	ar.caller = caller
	ar.retip = from
	ar.name = callee.name
	ar.locals = make([]Object, ar.nlocs)
	return ar
}