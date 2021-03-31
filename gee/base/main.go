package main

import (
	"fmt"
	"net/http"
	"playDemo/gee/base/gee"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Head[%q]=%q\n", k, v)
	}
}

// session 2
type Engine struct {
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Head[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	/*http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))*/
	// 2
	/*	e := gee.New()
		e.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
		})
		e.Get("/hello", func(writer http.ResponseWriter, request *http.Request) {
			for k, v := range request.Header {
				fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
			}
		})
		e.Run(":9999")*/
	// 3
	e := gee.New()
	e.Get("/hello/:name", func(c *gee.Context) {
		v := c.Param("name")
		c.String(http.StatusOK, "hello %s, u r at %s:%s\n", c.Query("name"), c.Path, v)
	})
	e.Run(":9999")
}
