package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	// map study
	var countryCapitalMap map[string]string
	countryCapitalMap["1"] = "2"

	for i, i2 := range countryCapitalMap {
		fmt.Println(i)
		fmt.Println(i2)
	}

	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	r.Run(":9999")

}
