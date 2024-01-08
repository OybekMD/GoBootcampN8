package encryption

func PasswordDecode(s string) string {
	r := []byte(s)
	for i := 0; i < len(s); i++ {
		r[i] = r[i] + 5
	}
	return string(r)
}
func PasswordEncode(s string) string {
	r := []byte(s)
	for i := 0; i < len(s); i++ {
		r[i] = r[i] - 5
	}
	return string(r)
}