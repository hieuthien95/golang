package main

import (
	"bytes"
	"fmt"
	"html/template"
)

func main() {
	tmpl, _ := template.ParseFiles("./confirm-mail.html")
	buf := new(bytes.Buffer)
	tmpl.Execute(buf, map[string]string{"email": "hieuthien95@gmail.com"})

	fmt.Println(buf.String())
}
