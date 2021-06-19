package lazyops

type Predicate func() bool

func Not(p Predicate) Predicate {
	return func() bool {
		return !p()
	}
}

func And(ps ...Predicate) Predicate {
	return func() bool {
		for _, p := range ps {
			if !p() {
				return false
			}
		}

		return true
	}
}

func Or(ps ...Predicate) Predicate {
	return func() bool {
		for _, p := range ps {
			if p() {
				return true
			}
		}

		return false
	}
}

func Xor(ps ...Predicate) Predicate {
	return func() bool {
		acc := false

		for _, p := range ps {
			if acc == p() {
				acc = false
			} else {
				acc = true
			}
		}

		return acc
	}
}
