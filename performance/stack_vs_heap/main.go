package main

func copy(sc string) string {
	ss := sc
	return ss
}

func reference(sr string) *string {
	ss := &sr
	return ss
}

func main() {
	s := "string"
	ss := copy(s)
	sss := reference(s)
	_ = ss
	_ = sss
}
