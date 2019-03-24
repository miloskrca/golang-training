package jsonmarshaling

import (
	"encoding/json"
	"fmt"
	"testing"
)

var object = struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}{
	Name:     "Name",
	Lastname: "Lastname",
}

func BenchmarkMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(object)
	}
}

func BenchmarkFMT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf(`
		{
			"name": "%s",
			"lastname": "%s"
		}`, object.Name, object.Lastname)
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = `
		{
			"name": ` + object.Name + `,
			"lastname": ` + object.Lastname + `,
		}`
	}
}
