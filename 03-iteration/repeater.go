package iteration

// TODO: Write ExampleRepeat to document your function
func Repeat(character string, times int) string {
	repeatCount := times
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}