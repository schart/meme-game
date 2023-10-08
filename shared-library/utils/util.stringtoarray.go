package utils

func ArrayCheck(arr []string) bool {
	for i := 0; i < len(arr); i++ {
		ch := string(arr[i])
		if ch == " " || ch == "" {
			return false
		}

	}
	return true
}
