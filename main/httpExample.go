package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(res,
		`
      <DOCTYPE html>
      <html>
        <head>
          <title>Hello World</title>
        </head>
        <body>
          Hello world !
        </body>
      </html>
    `,
	)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9099", nil)

	//handle file server
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("dummyDirectory")),
		),
	)
}
