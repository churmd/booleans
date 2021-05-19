package lazyops

type Predicate func() bool

func And(ps... Predicate) bool {
	for _, p := range ps {
		if !p() {
			return false
		}
	}

	return true
}