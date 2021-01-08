package couple

// Couple couple the given string
func Couple(str string) []string {
	result := []string{}
	strlen := len(str)
	for i := 0; i < strlen; i += 2 {
		var couple string
		if i+1 >= strlen {
			couple = str[i:i+1] + "*"
		} else {
			couple = str[i : i+2]
		}
		result = append(result, couple)
	}
	return result
}
