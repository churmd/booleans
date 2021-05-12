package ops

func Not(x bool) bool {
	return !x
}

func And(xs... bool) bool {
	for _, x := range xs {
		if !x {
			return false
		}
	}
	return true
}