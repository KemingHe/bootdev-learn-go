package mystrings

func Reverse(s string) string {
	sLen := len(s)
	rev := []rune(s)
	for i, char := range s {
		rev[sLen-1-i] = char
	}
	return string(rev)
}
