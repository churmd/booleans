package ops

func Not(x bool) bool {
	return !x
}

func And(xs ...bool) bool {
	for _, x := range xs {
		if !x {
			return false
		}
	}

	return true
}

func Or(xs ...bool) bool {
	for _, x := range xs {
		if x {
			return true
		}
	}

	return false
}

func Xor(xs ...bool) bool {
	acc := false

	for _, x := range xs {
		if acc == x {
			acc = false
		} else {
			acc = true
		}
	}

	return acc
}
