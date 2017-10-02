package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
)

type Server struct{}

func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	left := r.FormValue("left")
	right := r.FormValue("right")
	op := r.FormValue("op")

	leftInt, leftErr := strconv.Atoi(left)
	rightInt, rightErr := strconv.Atoi(right)

	var result string
	var opration string
	if (leftErr == nil) && (rightErr == nil) {
		switch op {
		case "add":
			result = strconv.Itoa(leftInt + rightInt)
			opration = "足し算"
		case "sub":
			result = strconv.Itoa(leftInt - rightInt)
			opration = "引き算"
		case "multi":
			result = strconv.Itoa(leftInt * rightInt)
			opration = "掛け算"
		case "div":
			result = strconv.Itoa(leftInt / rightInt)
			opration = "割り算"
		}
	}

	h := `
  <html><head><title>簡易電卓/title></head><body>
  <form action="/" method="post">
  左項目:<input type="text" name="left"><br>
  右項目:<input type="text" name="right"><br>
  演算子:
  <input type="radio" name="op" value="add" checked> +
  <input type="radio" name="op" value="sub"> -
  <input type="radio" name="op" value="multi"> ×
  <input type="radio" name="op" value="div"> ÷
  <input type="submit" name="送信"><hr>
  <br>
  左項目: ` + html.EscapeString(left) + `<br>
  右項目: ` + html.EscapeString(right) + `<br>
  演算子: ` + html.EscapeString(opration) + `<hr>
  Answer: ` + html.EscapeString(result) + `
  <form>
  </body></html>
  `

	fmt.Fprint(w, h)

}

func main() {
	http.ListenAndServe(":4000", Server{})
}
