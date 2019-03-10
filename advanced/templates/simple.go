package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

// var builtins = FuncMap{
//     "and":      and,
//     "call":     call,
//     "html":     HTMLEscaper,
//     "index":    index,
//     "js":       JSEscaper,
//     "len":      length,
//     "not":      not,
//     "or":       or,
//     "print":    fmt.Sprint,
//     "printf":   fmt.Sprintf,
//     "println":  fmt.Sprintln,
//     "urlquery": URLQueryEscaper,
// }

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}
type Friend struct {
	Fname string
}

func main() {
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailChecker": emailChecker})
	t, _ = t.Parse(`hello {{.UserName}}!
Emails:
{{range .Emails}}
	{{.|emailChecker}}
{{end}}
Friends:
{{with .Friends}}
{{range $k, $v := .}}
	{{$k}}. {{$v.Fname}}
	{{- if eq $v.Fname "Second Friend" }}!{{end}}
{{end}}
{{end}}
`)

	f1 := Friend{Fname: "First Friend"}
	f2 := Friend{Fname: "Second Friend"}
	p := Person{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}

func emailChecker(arg interface{}) string {
	s, ok := arg.(string)
	if !ok {
		log.Fatal("not a string")
	}
	return fmt.Sprintf("(%s)", strings.Replace(s, "@", " at ", -1))
}
