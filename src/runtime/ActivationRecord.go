package runtime

const DEFAULT_LOCAL_COUNT int = 1 << 3

type ActivationRecord struct {
	caller *ActivationRecord
	locals []Object
	retip  int
}

func Call(caller *ActivationRecord, ip int) *ActivationRecord {
	ar := new(ActivationRecord)
	ar.caller = caller
	ar.retip = ip
	ar.locals = make([]Object, DEFAULT_LOCAL_COUNT)
	return ar
}