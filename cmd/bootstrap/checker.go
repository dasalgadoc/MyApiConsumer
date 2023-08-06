package bootstrap

func CheckArgumentOnSlice(element string, args []string) bool {
	for _, arg := range args {
		if element == arg {
			return true
		}
	}
	return false
}
