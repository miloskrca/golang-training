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
	copy(s)
	reference(s)
}
