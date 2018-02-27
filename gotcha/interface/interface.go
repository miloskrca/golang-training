package main

type inter interface {
	Func() string
}

type A struct{}

func (a A) Func() string {
	return "a"
}

type B struct{}

func (b *B) Func() string {
	return "b"
}

func callFunc(i inter) string {
	return i.Func()
}

func main() {
	var a = &A{}
	var b = B{}
	callFunc(a)
	callFunc(b)
}
