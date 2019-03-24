package typeassertion

import (
	"testing"
)

type Inccer interface {
	inc()
}

type myint int64

func (i *myint) inc() {
	*i = *i + 1
}

func incnIntmethod(i *myint, n int) {
	for k := 0; k < n; k++ {
		i.inc()
	}
}

func BenchmarkIntmethod(b *testing.B) {
	i := new(myint)
	incnIntmethod(i, b.N)
}

func incnInterface(any Inccer, n int) {
	for k := 0; k < n; k++ {
		any.inc()
	}
}

func BenchmarkInterface(b *testing.B) {
	i := new(myint)
	incnInterface(i, b.N)
}

func incnSwitch(any Inccer, n int) {
	for k := 0; k < n; k++ {
		switch v := any.(type) {
		case *myint:
			v.inc()
		}
	}
}

func BenchmarkTypeSwitch(b *testing.B) {
	i := new(myint)
	incnSwitch(i, b.N)
}

func incnAssertion(any Inccer, n int) {
	for k := 0; k < n; k++ {
		if newint, ok := any.(*myint); ok {
			newint.inc()
		}
	}
}

func BenchmarkTypeAssertion(b *testing.B) {
	i := new(myint)
	incnAssertion(i, b.N)
}
