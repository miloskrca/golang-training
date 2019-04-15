package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Doc struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	ToOmit    string    `json:"toOmit,omitempty"`
}

func (u Doc) MarshalJSON() ([]byte, error) {
	u.ToOmit = ""
	type Alias Doc
	return json.Marshal(&struct {
		Alias
		UpdatedAt int64 `json:"updatedAt"`
	}{
		Alias:     (Alias)(u),
		UpdatedAt: u.UpdatedAt.Unix(),
	})
}

func TestMarshal(t *testing.T) {
	doc := Doc{ID: 1, Name: "name", ToOmit: "value", UpdatedAt: time.Now()}
	b, _ := json.Marshal(doc)
	fmt.Println(string(b))
}
