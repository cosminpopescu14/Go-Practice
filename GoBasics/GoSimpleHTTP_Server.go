package main

import (
	"fmt"
	"io"
	"net/http"
)

const port = ":8000"

func main() {

	http.HandleFunc("/", foo)
	http.ListenAndServe(port, nil)
	fmt.Println("Listening at: ", port)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<form action = "/" method = "post">
			<input type="text" name="text">
			<input type="submit">
		</form>
		<script src="https://gist.github.com/cosminpopescu14/1c25e7a270840b31426ab70c0ba286e5.js"></script>
	 `)
	io.WriteString(w, "q: "+r.FormValue("q")) //read a query string
	io.WriteString(w, `<br/>`)
	io.WriteString(w, "form value: "+r.FormValue("text"))
	fmt.Println("form value = ", r.Form["text"])
}
