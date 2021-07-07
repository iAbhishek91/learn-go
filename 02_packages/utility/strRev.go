package utility

func ReverseString(str string) string {
	output := ""
	for _, char := range str {
		output = string(char) + output
	}
	return output
}
