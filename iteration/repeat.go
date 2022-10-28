package iteration

func Repeat(character string, repeatCount int) string {
	var result string
	for i := 0; i < repeatCount; i++ {
		result += character
	}

	return result
}
