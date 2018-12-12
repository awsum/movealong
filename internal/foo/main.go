package foo


func Code() {
	a := 1
	if a > 2 {
		a = 3
	}
	a += 1
}

func NoCoverage() {
	a := 1
	if a > 2 {
		a = 3
	}
	a += 1
}