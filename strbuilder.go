package main

import (
	"bytes"
	"text/template"
)

type templExecutor func(v interface{}) string

func Make(s string) templExecutor {

	tmpl, err := template.New("new").Parse(s)
	if err != nil {
		panic(err)
	}

	return func(v interface{}) string {

		b := bytes.NewBufferString("")

		err = tmpl.Execute(b, v)
		if err != nil {
			panic(err)
		}
		return b.String()
	}
}
